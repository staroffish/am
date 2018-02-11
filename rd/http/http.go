package http

import (
	"crypto/tls"
	"fmt"
	"global"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"rd"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/garyburd/redigo/redis"
)

// HttpDownloader - http下载库
type HttpDownloader struct {
	pool  redisPool
	tasks map[string]chan struct{}
	db    int
}

var mapMutex sync.Mutex

// HttpType/HttpsType - 下载类型字符串
var HttpType = "http"
var HttpsType = "https"

type httpTask struct {
	// Id       string `redis:"Id"`       // 任务ID(保存文件的全路径)
	Path     string `redis:"Path"`     // 存储路径
	FileName string `redis:"FileName"` // 文件名
	Url      string `redis:"Url"`      // 下载链接
	// CanTruncate  string  `redis:"CanTruncate"`  // 是否可被分割(是否支持断点续传)
	TaskType     string  `redis:"TaskType"`     // 任务类型(http/https)
	State        string  `redis:"State"`        // 任务状态
	TotalSize    float64 `redis:"TotalSize"`    // 总大小
	FinishedSize float64 `redis:"FinishedSize"` // 完成大小
	CreateTime   string  `redis:"CreateTime"`   // 创建时间
}

func init() {
	rd.SetDownloader(HttpType, &HttpDownloader{db: 0})
	rd.SetDownloader(HttpsType, &HttpDownloader{db: 1})
}

// NewDownloader - 初始化下载器
func (d *HttpDownloader) InitDownloader(cfg *global.Config) error {
	defer global.TraceLog("http.HttpDownloader.NewDownloader")()

	if d.pool.Pool == nil {
		d.pool.initRedis(cfg.RedisAddr, cfg.RedisPasswd, d.db, 3)
	}
	d.tasks = make(map[string]chan struct{})
	var tasks = make([]*httpTask, 0)
	err := d.pool.getAllTask(&tasks)
	if err != nil {
		return fmt.Errorf("HttpDownloader.InitDownloader:%v", err)
	}

	for _, task := range tasks {
		if task.State == rd.Downloading {
			var ctrl = make(chan struct{})
			d.tasks[task.Url] = ctrl
			go d.startDownload(task, ctrl)
		}
	}
	return nil
}

// AddTask - 添加任务
func (d *HttpDownloader) AddTask(t *rd.RdTask) error {
	defer global.TraceLog("http.HttpDownloader.AddTask")()

	task := &httpTask{}
	// task.Id = fmt.Sprintf("%s/%s", t.SavePath, t.Name)
	task.Path = t.SavePath
	task.FileName = t.Name
	task.Url = t.Link
	// task.CanTruncate = CanNotTruncate
	task.TaskType = t.TaskType
	task.State = rd.Downloading
	task.CreateTime = global.FormatTime(time.Now())
	task.TotalSize = 0
	task.FinishedSize = 0

	if err := d.pool.do("HMSET", redis.Args{}.Add(task.Url).AddFlat(task)...); err != nil {
		return fmt.Errorf("HttpDownloader.AddTask:%v", err)
	}

	var ctrl = make(chan struct{})
	mapMutex.Lock()
	d.tasks[task.Url] = ctrl
	mapMutex.Unlock()
	go d.startDownload(task, ctrl)
	<-ctrl
	return nil
}

// PauseTask - 暂停任务
func (d *HttpDownloader) PauseTask(t *rd.RdTask) error {
	defer global.TraceLog("http.HttpDownloader.PauseTask")()

	mapMutex.Lock()
	ctrl, ok := d.tasks[t.Ids]
	if !ok {
		global.Log.Warnf("am:HttpDownloader.PauseTask:task was already paused or deleted:Url=%s", t.Link)
		mapMutex.Unlock()
		return nil
	}
	delete(d.tasks, t.Ids)
	ctrl <- struct{}{}
	mapMutex.Unlock()

	return nil
}

// DeleteTask - 删除任务(只删除任务不删除文件)
func (d *HttpDownloader) DeleteTask(t *rd.RdTask) error {
	defer global.TraceLog("http.HttpDownloader.DeleteTask")()
	mapMutex.Lock()
	ctrl, ok := d.tasks[t.Ids]
	if ok {
		delete(d.tasks, t.Ids)
		ctrl <- struct{}{}
		<-ctrl
	}
	mapMutex.Unlock()
	if err := d.pool.do("DEL", t.Ids); err != nil {
		return fmt.Errorf("HttpDownloader.DeleteTask:%v", err)
	}
	return nil
}

// ResumeTask - 重新开始任务
func (d *HttpDownloader) ResumeTask(t *rd.RdTask) error {
	defer global.TraceLog("http.HttpDownloader.ResumeTask")()

	task := &httpTask{}
	err := d.pool.getOneTask(t.Ids, task)
	if err != nil {
		return fmt.Errorf("HttpDownloader.ResumeTask:%v", err)
	}

	if err := d.pool.do("HSET", t.Ids, "State", rd.Downloading); err != nil {
		return fmt.Errorf("HttpDownloader.AddTask:%v", err)
	}

	ctrl := make(chan struct{})
	mapMutex.Lock()
	d.tasks[t.Ids] = ctrl
	mapMutex.Unlock()

	go d.startDownload(task, ctrl)

	return nil
}

// GetAllTask - 获取所有任务
func (d *HttpDownloader) GetAllTask() ([]rd.RdTask, error) {
	defer global.TraceLog("http.HttpDownloader.GetAllTask")()

	var rdTaskList = make([]rd.RdTask, 0)
	var tasks = make([]*httpTask, 0)

	err := d.pool.getAllTask(&tasks)
	if err != nil {
		return nil, fmt.Errorf("http.GetAllTask:%v", err)
	}

	for _, task := range tasks {
		var rdt rd.RdTask
		// rdt.Ids = task.Id
		rdt.Ids = task.Url
		rdt.Link = task.Url
		rdt.Name = task.FileName
		rdt.SavePath = task.Path

		if task.TotalSize != 0 {
			rdt.Progress = int(task.FinishedSize / task.TotalSize * 100)
		}

		rdt.Size = int64(task.TotalSize)
		rdt.TaskType = task.TaskType
		rdt.State = task.State
		rdt.CreateTime = task.CreateTime
		rdTaskList = append(rdTaskList, rdt)
	}

	return rdTaskList, nil
}

func (d *HttpDownloader) startDownload(t *httpTask, ctrl chan struct{}) {
	defer global.TraceLog("http.startDownload")()

	defer func() {
		if t.State == rd.Downloading {
			if err := d.pool.do("HSET", t.Url, "State", rd.Paused); err != nil {
				global.Log.Errorf("am:http.startDownload:d.pool.Do error:%v", err)
			}
		} else {
			if err := d.pool.do("HSET", t.Url, "State", t.State); err != nil {
				global.Log.Errorf("am:http.startDownload:d.pool.Do error:%v", err)
			}
		}
		ctrl <- struct{}{}
	}()

	req, err := http.NewRequest(http.MethodGet, t.Url, nil)
	if err != nil {
		t.State = rd.Error
		global.Log.Errorf("am:http.startDownload:http.NewRequest error:%v", err)
		return
	}

	req.Header.Set("Range", fmt.Sprintf("bytes=%d-", int64(t.FinishedSize)))

	client := http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // disable verify
	}}

	resp, err := client.Do(req)
	if err != nil {
		t.State = rd.Error
		global.Log.Errorf("am:http.startDownload:http.DefaultClient.Do error:%v", err)
		return
	}

	defer resp.Body.Close()

	// 如果是添加任务
	if len(t.FileName) == 0 {
		//取得文件名
		t.FileName = getFileName(t.Url, resp)
		t.TotalSize, _ = strconv.ParseFloat(resp.Header.Get("Content-Length"), 64)
		if err := d.pool.do("HMSET", t.Url, "FileName", t.FileName, "TotalSize", t.TotalSize); err != nil {
			t.State = rd.Error
			global.Log.Errorf("am:http.startDownload:update task error:%v", err)
			return
		}
		ctrl <- struct{}{}
	}

	var file *os.File
	if resp.StatusCode == http.StatusPartialContent {
		file, err = os.OpenFile(fmt.Sprintf("%s/%s", t.Path, t.FileName),
			os.O_APPEND|os.O_CREATE|os.O_RDWR, 0755)
	} else {
		file, err = os.OpenFile(fmt.Sprintf("%s/%s", t.Path, t.FileName),
			os.O_TRUNC|os.O_CREATE|os.O_RDWR, 0755)
	}
	if err != nil {
		global.Log.Errorf("am:http.startDownload:open file:%v", err)
		t.State = rd.Error
		return
	}
	defer file.Close()

EXIT_LOOP:
	for {
		select {
		case <-ctrl:
			break EXIT_LOOP
		default:
			n, err := io.CopyN(file, resp.Body, 40960)
			if n != 0 {
				t.FinishedSize += float64(n)
				if err := d.pool.do("HSET", t.Url, "FinishedSize", t.FinishedSize); err != nil {
					t.State = rd.Error
					global.Log.Errorf("am:http.startDownload:update size error:%v", err)
					break EXIT_LOOP
				}
			}
			if err == io.EOF {
				t.State = rd.Finished
				mapMutex.Lock()
				delete(d.tasks, t.Url)
				mapMutex.Unlock()
				break EXIT_LOOP
			} else if err != nil {
				if err := d.pool.do("HSET", t.Url, "State", rd.Error); err != nil {
					t.State = rd.Error
					global.Log.Errorf("am:http.startDownload:update state to %s error:%v", rd.Error, err)
					break EXIT_LOOP
				}
				global.Log.Errorf("am:http.startDownload: io.CopyN error:%v", err)
				break EXIT_LOOP
			}
		}
	}
}

func getFileName(url string, resp *http.Response) string {
	var fileName string
	if disp := resp.Header.Get("Content-Disposition"); len(disp) != 0 {
		boundary := "--aaaaaaaa"
		strReader := strings.NewReader(fmt.Sprintf("--%s\r\nContent-Disposition: %s\r\n\r\n--%s--\r\n",
			boundary, disp, boundary))

		reader := multipart.NewReader(strReader, boundary)
		part, err := reader.NextPart()
		if err != nil {
			fmt.Printf("reader.NextPart content-dispostion error:%v\n", err)
			os.Exit(-1)
		}
		fileName = part.FileName()
	} else {
		ind := strings.LastIndex(url, "/")
		if ind == len(url)-1 {
			fileName = time.Now().Format("20060102150405000")
		} else {
			fileName = url[ind+1:]
		}
	}
	return fileName
}

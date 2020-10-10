package deluge

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/staroffish/am/rd"

	"github.com/staroffish/am/global"
)

// TaskType - 下载类型字符串
var taskType = "magnet"

// DelugeDownloader - deluge下载
type DelugeDownloader struct {
	password   string
	requestUrl string
	timeout    int
}

type normalRes struct {
	Id     int         `json:"id"`
	Result interface{} `json:"result"`
	Err    interface{} `json:"error"`
}

type GetHostRes struct {
	Id     int             `json:"id"`
	Result [][]interface{} `json:"result"`
	Err    interface{}     `json:"error"`
}

const (
	loginFmt   = `{"method":"auth.login","params":["%s"],"id":1}`
	logoutFmt  = `{"method":"auth.delete_session","params":[],"id":3}`
	addFmt     = `{"method":"web.add_torrents","params":[[{"path":"%s","options":{"file_priorities":[],"add_paused":false,"sequential_download":false,"pre_allocate_storage":false,"download_location":"%s","move_completed":false,"move_completed_path":"%s","max_connections":-1,"max_download_speed":-1,"max_upload_slots":-1,"max_upload_speed":-1,"prioritize_first_last_pieces":false,"seed_mode":false,"super_seeding":false}}]],"id":31}`
	pauseFmt   = `{"method":"core.pause_torrent","params":[["%s"]],"id":2}`
	deleteFmt  = `{"method":"core.remove_torrent","params":["%s",false],"id":460}`
	resumeFmt  = `{"method":"core.resume_torrent","params":[["%s"]],"id":2}`
	searchFmt  = `{"method":"web.update_ui","params":[["name","total_wanted","state","progress","time_added","save_path"],{}],"id":2}`
	getHostFmt = `{"method":"web.get_hosts","params":[],"id":53}`
	connectFmt = `{"method":"web.connect","params":["%s"],"id":149}`
)

func init() {
	var dnldr DelugeDownloader
	rd.SetDownloader(taskType, &dnldr)
}

// NewDownloader - 初始化下载器
func (d *DelugeDownloader) InitDownloader(cfg *global.Config) error {
	defer global.TraceLog("deluge.DelugeDownloader.NewDownloader")()
	if cfg.BTWebPass == "" ||
		cfg.BTWebUrl == "" ||
		cfg.WebTimeout == 0 {
		return fmt.Errorf("deluge.NewDownloader:Config error")
	}
	d.password = cfg.BTWebPass
	d.requestUrl = cfg.BTWebUrl
	d.timeout = cfg.WebTimeout
	return nil
}

func (d *DelugeDownloader) check() error {
	defer global.TraceLog("deluge.DelugeDownloader.check")()
	if d.password == "" || d.requestUrl == "" {
		return errors.New("BTWebPass or BTWebUrl is empty")
	}
	return nil
}

func (d *DelugeDownloader) connectHost(cookies *http.Cookie) error {
	defer global.TraceLog("deluge.DelugeDownloader.connectHost")()

	// 获取服务器列表
	resp, err := d.sendReq(cookies, getHostFmt)
	if err != nil {
		return fmt.Errorf("am:DelugeDownloader:connectHost:sendReq:%v", err)
	}
	buff := bytes.NewBuffer([]byte{})
	_, err = io.Copy(buff, resp.Body)
	resp.Body.Close()
	if err != nil {
		return fmt.Errorf("am:DelugeDownloader:connectHost:copy http response error:%v", err)
	}
	getHostRes := &GetHostRes{}
	if err := json.Unmarshal(buff.Bytes(), getHostRes); err != nil {
		return fmt.Errorf("am:DelugeDownloader:connectHost:Unmarshal http response error:%v", err)
	}

	if getHostRes.Err != nil {
		return fmt.Errorf("am:DelugeDownloader:connectHost:return error:%v", getHostRes.Err)
	}

	if len(getHostRes.Result) <= 0 || len(getHostRes.Result[0]) <= 0 {
		return fmt.Errorf("am:DelugeDownloader:connectHost:get a invalid response:result=%v", getHostRes.Result)
	}

	hostId, ok := getHostRes.Result[0][0].(string)
	if !ok {
		return fmt.Errorf("am:DelugeDownloader:connectHost:get a invalid response:result=%v", getHostRes.Result)
	}

	// 连接服务器
	connectHostReq := fmt.Sprintf(connectFmt, hostId)
	resp, err = d.sendReq(cookies, connectHostReq)
	if err != nil {
		return fmt.Errorf("am:DelugeDownloader:connectHost:sendReq:%v", err)
	}
	defer resp.Body.Close()

	if checkResp(resp.Body) == false {
		return fmt.Errorf("am:DelugeDownloader.connectHost: connect host error:hostId=%s", hostId)
	}
	return nil
}

func (d *DelugeDownloader) login() (*http.Cookie, error) {
	defer global.TraceLog("deluge.DelugeDownloader.login")()

	jsonStr := fmt.Sprintf(loginFmt, d.password)

	resp, err := d.sendReq(nil, jsonStr)
	if err != nil {
		return nil, fmt.Errorf("am:DelugeDownloader.login: sendReq:%v", err)
	}
	defer resp.Body.Close()

	if checkResp(resp.Body) == false {
		return nil, fmt.Errorf("am:DelugeDownloader.login: auth failure:password=%s", d.password)
	}

	cookies := resp.Cookies()
	if len(cookies) == 0 {
		return nil, nil
	}

	return cookies[0], nil
}

func (d *DelugeDownloader) logout(cookie *http.Cookie) {
	defer global.TraceLog("deluge.DelugeDownloader.logout")()

	resp, err := d.sendReq(cookie, logoutFmt)
	if err != nil {
		global.Log.Errorf("am:DelugeDownloader.login:Do:%v", err)
	}
	resp.Body.Close()
}

func addHeader(r *http.Request, cookie *http.Cookie) {
	defer global.TraceLog("deluge.deluge.addHeader")()
	r.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:57.0) Gecko/20100101 Firefox/57.0")
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-Requested-With", "XMLHttpRequest")
	r.Header.Add("Connection", "keep-alive")
	if cookie != nil {
		r.AddCookie(cookie)
	}
}

func checkResp(body io.ReadCloser) bool {
	defer global.TraceLog("deluge.checkResp")()
	decode := json.NewDecoder(body)
	var res normalRes
	err := decode.Decode(&res)
	if err != nil {
		global.Log.Errorf("am:deluge.checkResp error:%v", err)
		return false
	}
	if res.Err != nil {
		global.Log.Errorf("am:deluge.checkResp errmsg:%v", res.Err)
		return false
	}

	return true
}

// AddTask - 添加任务
func (d *DelugeDownloader) AddTask(t *rd.RdTask) error {
	defer global.TraceLog("deluge.DelugeDownloader.AddTask")()
	cookie, err := d.login()
	if err != nil {
		return err
	}
	defer d.logout(cookie)

	if err := d.connectHost(cookie); err != nil {
		global.Log.Errorf("connect host error:%v", err)
	}

	jsonStr := fmt.Sprintf(addFmt, t.Link, t.SavePath, t.SavePath)

	resp, err := d.sendReq(cookie, jsonStr)
	if err != nil {
		return fmt.Errorf("DelugeDownloader.AddTask:%v", err)
	}
	defer resp.Body.Close()

	if checkResp(resp.Body) == false {
		return fmt.Errorf("DelugeDownloader.AddTask:link=%s:path=%s", t.Link, t.SavePath)
	}

	return nil
}

// PauseTask - 暂停任务
func (d *DelugeDownloader) PauseTask(t *rd.RdTask) error {
	defer global.TraceLog("deluge.DelugeDownloader.PauseTask")()
	cookie, err := d.login()
	if err != nil {
		return err
	}
	defer d.logout(cookie)

	jsonStr := fmt.Sprintf(pauseFmt, t.Ids)

	resp, err := d.sendReq(cookie, jsonStr)
	if err != nil {
		return fmt.Errorf("DelugeDownloader.PauseTask:%v", err)
	}
	defer resp.Body.Close()

	return nil
}

// DeleteTask - 删除任务(只删除任务不删除文件)
func (d *DelugeDownloader) DeleteTask(t *rd.RdTask) error {
	defer global.TraceLog("deluge.DelugeDownloader.DeleteTask")()
	cookie, err := d.login()
	if err != nil {
		return err
	}
	defer d.logout(cookie)

	jsonStr := fmt.Sprintf(deleteFmt, t.Ids)

	resp, err := d.sendReq(cookie, jsonStr)
	if err != nil {
		return fmt.Errorf("DelugeDownloader.DeleteTask:%v", err)
	}
	defer resp.Body.Close()

	if checkResp(resp.Body) == false {
		return fmt.Errorf("DelugeDownloader.DeleteTask:id=%s", t.Ids)
	}

	return nil
}

// ResumeTask - 重新开始任务
func (d *DelugeDownloader) ResumeTask(t *rd.RdTask) error {
	defer global.TraceLog("deluge.DelugeDownloader.ResumeTask")()
	cookie, err := d.login()
	if err != nil {
		return err
	}
	defer d.logout(cookie)

	jsonStr := fmt.Sprintf(resumeFmt, t.Ids)

	resp, err := d.sendReq(cookie, jsonStr)
	if err != nil {
		return fmt.Errorf("DelugeDownloader.ResumeTask:%v", err)
	}
	defer resp.Body.Close()

	return nil
}

// GetAllTask - 获取所有任务
func (d *DelugeDownloader) GetAllTask() ([]rd.RdTask, error) {
	defer global.TraceLog("deluge.DelugeDownloader.GetAllTask")()
	cookie, err := d.login()
	if err != nil {
		return nil, err
	}

	defer d.logout(cookie)
	resp, err := d.sendReq(cookie, searchFmt)
	if err != nil {
		return nil, fmt.Errorf("DelugeDownloader.GetAllTask:%v", err)
	}
	defer resp.Body.Close()

	var tasks = make([]rd.RdTask, 0)

	// 取得body
	buf := bytes.NewBuffer([]byte{})
	_, err = io.Copy(buf, resp.Body)
	if err != nil {
		return nil, fmt.Errorf("DelugeDownloader.GetAllTask:%v", err)
	}

	body := buf.Bytes()

	// 取得任务列表
	sInd := bytes.Index(body, []byte("\"torrents\": "))
	if sInd == -1 {
		return nil, fmt.Errorf("DelugeDownloader.GetAllTask:Error response:%s", body)
	}

	sInd += len("\"torrents\": ")
	eInd := bytes.Index(body[sInd:], []byte(", \"filters\""))
	if eInd == -1 {
		return nil, fmt.Errorf("DelugeDownloader.GetAllTask:Error response:%s", body)
	}

	taskStr := body[sInd : sInd+eInd]
	var taskList = make(map[string]struct {
		State     string  `json:"state"`
		Progress  float64 `json:"progress"`
		SavePath  string  `json:"save_path"`
		TimeAdded float64 `json:"time_added"`
		Size      float64 `json:"total_wanted"`
		Name      string  `json:"name"`
	})
	err = json.Unmarshal(taskStr, &taskList)
	if err != nil {
		return nil, fmt.Errorf("DelugeDownloader.GetAllTask:taskList:%v", err)
	}

	for id, t := range taskList {
		var task rd.RdTask
		task.Ids = id
		task.Name = t.Name
		task.SavePath = t.SavePath
		task.Progress = int(t.Progress)
		task.State = t.State
		task.Size = int64(t.Size)
		task.TaskType = taskType
		task.CreateTime = global.FormatTime(time.Unix(int64(t.TimeAdded), 0))
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (d *DelugeDownloader) sendReq(cookie *http.Cookie, jsonStr string) (*http.Response, error) {
	defer global.TraceLog("deluge.DelugeDownloader.sendReq")()

	r, err := http.NewRequest(http.MethodPost, d.requestUrl, strings.NewReader(jsonStr))
	if err != nil {
		return nil, fmt.Errorf("DelugeDownloader.sendReq.NewRequest:%v", err)
	}

	addHeader(r, cookie)
	timeout := make(chan struct{})
	r.Cancel = timeout
	go func() {
		defer close(timeout)
		time.Sleep(time.Duration(d.timeout) * time.Second)
	}()

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return nil, fmt.Errorf("DelugeDownloader.sendReq.Do:err=%v, jsonStr=%s, url=%s", err, jsonStr, d.requestUrl)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("DelugeDownloader.sendReq:status is not OK")
	}

	return resp, nil
}

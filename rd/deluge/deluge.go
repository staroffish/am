package deluge

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"global"
	"io"
	"net/http"
	"rd"
	"strings"
	"time"
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
	Id     int    `json:"id"`
	Result bool   `json:"result"`
	Err    string `json:"error"`
}

const (
	loginFmt  = `{"method":"auth.login","params":["%s"],"id":1}`
	logoutFmt = `{"method":"auth.delete_session","params":[],"id":3}`
	addFmt    = `{"method":"web.add_torrents","params":[[{"path":"%s","options":{"file_priorities":[],"add_paused":false,"compact_allocation":false,"download_location":"%s","move_completed":false,"move_completed_path":"%s","max_connections":-1,"max_download_speed":-1,"max_upload_slots":-1,"max_upload_speed":-1,"prioritize_first_last_pieces":false}}]],"id":2}`
	pauseFmt  = `{"method":"core.pause_torrent","params":[["%s"]],"id":2}`
	deleteFmt = `{"method":"core.remove_torrent","params":["%s",false],"id":460}`
	resumeFmt = `{"method":"core.resume_torrent","params":[["%s"]],"id":2}`
	searchFmt = `{"method":"web.update_ui","params":[["queue","name","total_wanted","state","progress","num_seeds","total_seeds","num_peers","total_peers","download_payload_rate","upload_payload_rate","eta","ratio","distributed_copies","is_auto_managed","time_added","tracker_host","save_path","total_done","total_uploaded","max_download_speed","max_upload_speed","seeds_peers_ratio"],{}],"id":2}`
)

func init() {
	var dnldr DelugeDownloader
	rd.SetDownloader(taskType, &dnldr)
}

func (d *DelugeDownloader) check() error {
	defer global.TraceLog("deluge.DelugeDownloader.check")()
	if d.password == "" || d.requestUrl == "" {
		return errors.New("BTWebPass or BTWebUrl is empty")
	}
	return nil
}

func (d *DelugeDownloader) login() (*http.Cookie, error) {
	defer global.TraceLog("deluge.DelugeDownloader.login")()
	jsonStr := fmt.Sprintf(loginFmt, d.password)

	resp, err := d.sendReq(nil, jsonStr)
	if err != nil {
		return nil, fmt.Errorf("sendReq:%v", err)
	}
	defer resp.Body.Close()

	if checkResp(resp.Body) == false {
		return nil, fmt.Errorf("DelugeDownloader.login: auth failure:password=%s", d.password)
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
	if res.Err != "" {
		global.Log.Errorf("am:deluge.checkResp errmsg:%v", res.Err)
	}

	if res.Result != true {
		global.Log.Debugf("am:deluge.checkResp res:%v", res)
	}

	return res.Result
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

// AddTask - 添加任务
func (d *DelugeDownloader) AddTask(t *rd.RdTask) error {
	defer global.TraceLog("deluge.DelugeDownloader.AddTask")()
	cookie, err := d.login()
	if err != nil {
		return err
	}
	defer d.logout(cookie)

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

	// 取得任务总数
	body := buf.Bytes()
	sInd := bytes.Index(body, []byte("{\"state\": ["))
	if sInd == -1 {
		return nil, fmt.Errorf("DelugeDownloader.GetAllTask:Error response:%s", body)
	}
	eInd := bytes.Index(body[sInd:], []byte("}"))
	if eInd == -1 {
		return nil, fmt.Errorf("DelugeDownloader.GetAllTask:Error response:%s", body)
	}
	stateStr := body[sInd : sInd+eInd+1]

	var totalInfo struct {
		State []interface{} `json:"state"`
	}
	err = json.Unmarshal(stateStr, &totalInfo)
	if err != nil {
		return nil, fmt.Errorf("DelugeDownloader.GetAllTask:taskCnt:%v", err)
	}

	if len(totalInfo.State) == 0 {
		return nil, fmt.Errorf("DelugeDownloader.GetAllTask:state:0", err)
	}

	states, ok := totalInfo.State[0].([]interface{})
	if !ok {
		return nil, fmt.Errorf("DelugeDownloader.GetAllTask:parse task count error")
	}

	taskCnt, ok := states[1].(float64)
	if !ok {
		return nil, fmt.Errorf("DelugeDownloader.GetAllTask:parse task count error")
	}
	if taskCnt == 0 {
		return tasks, nil
	}

	// 取得任务列表
	sInd = bytes.Index(body, []byte("\"torrents\": "))
	if sInd == -1 {
		return nil, fmt.Errorf("DelugeDownloader.GetAllTask:Error response:%s", body)
	}

	sInd += len("\"torrents\": ")
	eInd = bytes.Index(body[sInd:], []byte(", \"filters\""))
	if eInd == -1 {
		return nil, fmt.Errorf("DelugeDownloader.GetAllTask:Error response:%s", body)
	}

	taskStr := body[sInd : sInd+eInd]
	var taskList = make(map[string]interface{})
	err = json.Unmarshal(taskStr, &taskList)
	if err != nil {
		return nil, fmt.Errorf("DelugeDownloader.GetAllTask:taskList:%v", err)
	}

	for id, t := range taskList {
		var task rd.RdTask
		task.Ids = id
		ctx, ok := t.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("DelugeDownloader.GetAllTask:taskList:no map")
		}
		task.Name, ok = ctx["name"].(string)
		task.SavePath, ok = ctx["save_path"].(string)
		tmpFloat, ok := ctx["progress"].(float64)
		if ok {
			task.Progress = int(tmpFloat)
		}
		task.State, ok = ctx["state"].(string)
		tmpFloat, ok = ctx["total_wanted"].(float64)
		if ok {
			task.Size = int64(tmpFloat)
		}
		task.TaskType = taskType

		tmpTime, ok := ctx["time_added"].(float64)
		if !ok {
			task.CreateTime = global.FormatTime(time.Time{})
		} else {
			task.CreateTime = global.FormatTime(time.Unix(int64(tmpTime), 0))
		}

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
		return nil, fmt.Errorf("DelugeDownloader.sendReq.Do:%v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("DelugeDownloader.sendReq:status is not OK")
	}

	return resp, nil
}

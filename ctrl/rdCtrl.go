package ctrl

import (
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/websocket"

	"github.com/staroffish/am/global"

	"github.com/staroffish/am/rd"
	"github.com/staroffish/am/view"
)

// MainCtrl 主页逻辑控制
type RdCtrl struct {
	view.RdPage
	PushChan chan struct{}
}

// Init - 初始化
func (r *RdCtrl) Init() error {
	r.PushChan = make(chan struct{})
	return r.RdPage.Init()
}

func (r *RdCtrl) Close() {
	rd.PauseAllTask()
}

// Process 处理函数
func (r *RdCtrl) Process(jr *JSONRequest, w http.ResponseWriter) error {
	defer global.TraceLog("RdCtrl.Process")()

	var link, taskType, id string
	var idx int
	if jr.Method != "get_task" {
		param1, ok := jr.Params[0].(string)
		if !ok {
			return fmt.Errorf("RdCtrl.Process:rd.add_task:link type error:%T", jr.Params[0])
		}

		link = param1
		idx = strings.LastIndex(param1, ":")
		if idx != -1 {
			id = param1[:idx]
			taskType = param1[idx+1:]
			link = param1
		}
	}

	switch jr.Method {
	case "add_task":
		idx = strings.Index(link, ":")
		if idx == -1 {
			link = fmt.Sprintf("http://%s", link)
			idx = 4
		}
		taskType = link[:idx]
		savePath, ok := jr.Params[1].(string)
		if !ok {
			return fmt.Errorf("RdCtrl.Process:rd.add_task:savePath type error:%T", jr.Params[1])
		}
		if err := rd.AddTask(&rd.RdTask{Link: link, SavePath: savePath, TaskType: taskType}); err != nil {
			return fmt.Errorf("RdCtrl.Process:rd.add_task:rd.AddTask:%v", err)
		}
	case "start_task":
		if err := rd.ResumeTask(&rd.RdTask{Ids: id, TaskType: taskType}); err != nil {
			return fmt.Errorf("RdCtrl.Process:rd.ResumeTask:%v", err)
		}
	case "pause_task":
		if err := rd.PauseTask(&rd.RdTask{Ids: id, TaskType: taskType}); err != nil {
			return fmt.Errorf("RdCtrl.Process:rd.PauseTask:%v", err)
		}
	case "del_task":
		if err := rd.DeleteTask(&rd.RdTask{Ids: id, TaskType: taskType}); err != nil {
			return fmt.Errorf("RdCtrl.Process:rd.PauseTask:%v", err)
		}
	case "get_task":
		_, err := w.Write(r.RdPage.Head)
		if err != nil {
			return fmt.Errorf("RdCtrl.Process:rd.get_task:%v", err)
		}
	}

	rd.UpdateChan <- struct{}{}
	<-rd.UpdateChan
	r.PushChan <- struct{}{}

	return nil
}

func (r *RdCtrl) PushTasks(w http.ResponseWriter, req *http.Request) {
	defer global.TraceLog("RdCtrl.PushTasks")()
	websocket.Handler(func(ws *websocket.Conn) {
		ticker := time.NewTicker(time.Second * 5)
		for {
			select {
			case <-r.PushChan:
			case <-ticker.C:
			}
			rdTasks := rd.GetCachedTask()
			if err := r.RdPage.PushPageCtx(ws, rdTasks); err != nil {
				_, ok := err.(*net.OpError)
				if !ok {
					global.Log.Errorf("Push rd task error:%v", err)
				}

				return
			}
		}
	}).ServeHTTP(w, req)
}

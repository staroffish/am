package ctrl

import (
	"fmt"
	"global"
	"net/http"
	"view"
	"rd"
	"time"
)

// MainCtrl 主页逻辑控制
type RdCtrl struct {
	view.RdPage
}

// Init - 初始化
func (r *RdCtrl) Init() error {
	return r.RdPage.Init()
}

// Process 处理函数
func (r *RdCtrl) Process(jr *JSONRequest, w http.ResponseWriter) error {
	defer global.TraceLog("RdCtrl.Process")()

	switch jr.Method {
	case "add_task":
		link, ok := jr.Params[0].(string)
		if !ok {
			return fmt.Errorf("RdCtrl.Process:rd.add_task:link type error:%T", jr.Params[0])
		}
		savePath, ok := jr.Params[1].(string)
		if !ok {
			return fmt.Errorf("RdCtrl.Process:rd.add_task:savePath type error:%T", jr.Params[1])
		}
		if err := rd.AddTask(&rd.RdTask{Link:link, SavePath:savePath}, "magnet"); err != nil {
			return fmt.Errorf("RdCtrl.Process:rd.add_task:rd.AddTask:%v", err)
		}
	case "start_task":
		ids, ok := jr.Params[0].(string)
		if !ok {
			return fmt.Errorf("RdCtrl.Process:rd.ResumeTask:ids type error:%T", jr.Params[0])
		}
		if err := rd.ResumeTask(&rd.RdTask{Ids:ids}, "magnet"); err != nil {
			return fmt.Errorf("RdCtrl.Process:rd.ResumeTask:%v", err)
		}
	case "pause_task":
		ids, ok := jr.Params[0].(string)
		if !ok {
			return fmt.Errorf("RdCtrl.Process:rd.PauseTask:ids type error:%T", jr.Params[0])
		}
		if err := rd.PauseTask(&rd.RdTask{Ids:ids}, "magnet"); err != nil {
			return fmt.Errorf("RdCtrl.Process:rd.PauseTask:%v", err)
		}
	case "del_task":
		ids, ok := jr.Params[0].(string)
		if !ok {
			return fmt.Errorf("RdCtrl.Process:rd.PauseTask:ids type error:%T", jr.Params[0])
		}
		if err := rd.DeleteTask(&rd.RdTask{Ids:ids}, "magnet"); err != nil {
			return fmt.Errorf("RdCtrl.Process:rd.PauseTask:%v", err)
		}
	}

	time.Sleep(500 * time.Millisecond)
	tasks, err := rd.GetAllTask()
	if err != nil {
		return fmt.Errorf("RdPage:rd.GetAllTask:%v", err)
	}

	return r.ShowPage(tasks, w)
}

package view

import (
	"db"
	"fmt"
	"global"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

type DeleteAdTask struct{}

func (d *DeleteAdTask) Init() error {
	return nil
}

// ShowPageCtx - 显示页面
func (d *DeleteAdTask) ShowPageCtx(req *JSONRequest, w http.ResponseWriter) error {
	defer global.TraceLog("DeleteAdTask.ShowPageCtx")()

	if len(req.Params) < 1 {
		return fmt.Errorf("DeleteAdTask:ShowPageCtx:Parameter Num is less then 8")
	}

	_id, ok := req.Params[0].(string)
	if !ok {
		return fmt.Errorf("UpdateAdTask:ShowPageCtx:Parameter type error:%v", req.Params[0])
	}


	err := db.DeleteAdTask(bson.ObjectIdHex(_id));
	if err != nil {
		return fmt.Errorf("UpdateAdTask:ShowPageCtx:DeletedTask err:%v", err)
	}

	fmt.Fprint(w, "show_adtask()");
	return nil
}

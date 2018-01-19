package view

import (
	"db"
	"fmt"
	"global"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

type DeleteAnime struct{}

func (d *DeleteAnime) Init() error {
	return nil
}

// ShowPageCtx - 显示页面
func (d *DeleteAnime) ShowPageCtx(req *JSONRequest, w http.ResponseWriter) error {
	defer global.TraceLog("DeleteAnime.ShowPageCtx")()

	if len(req.Params) < 1 {
		return fmt.Errorf("DeleteAnime:ShowPageCtx:Parameter Num is less then 8")
	}

	_id, ok := req.Params[0].(string)
	if !ok {
		return fmt.Errorf("DeleteAnime:ShowPageCtx:Parameter type error:%v", req.Params[0])
	}

	err := db.DeletedAnime(bson.ObjectIdHex(_id));
	if err != nil {
		return fmt.Errorf("DeleteAnime:ShowPageCtx:DeletedTask err:%v", err)
	}

	fmt.Fprint(w, "show_collection('main')");
	return nil
}

package view

import (
	"strconv"
	"db"
	"fmt"
	"global"
	"gopkg.in/mgo.v2/bson"
	// "html/template"
	"net/http"
)

type UpdateAdTask struct{}

func (u *UpdateAdTask) Init() error {
	return nil
}

// ShowPageCtx - 显示页面
func (u *UpdateAdTask) ShowPageCtx(req *JSONRequest, w http.ResponseWriter) error {
	defer global.TraceLog("UpdateAdTask.ShowPageCtx")()

	var adTask db.AdTask
	var ani db.Anime
	if len(req.Params) < 8 {
		return fmt.Errorf("UpdateAdTask:ShowPageCtx:Parameter Num is less then 8")
	}

	_id, ok := req.Params[0].(string)
	if !ok {
		return fmt.Errorf("UpdateAdTask:ShowPageCtx:Parameter type error:%v", req.Params[0])
	}

	if _id == "" {
		_id = bson.NewObjectId().Hex()
	}
	adTask.Id = bson.ObjectIdHex(_id);
	var i = 1
	ani.AnimeNameJp, ok = req.Params[i].(string)
	if !ok {
		return fmt.Errorf("UpdateAdTask:ShowPageCtx:Parameter type error:%v", req.Params[i])
	}
	i++

	adTask.Url, ok = req.Params[i].(string)
	if !ok {
		return fmt.Errorf("UpdateAdTask:ShowPageCtx:Parameter type error:%v", req.Params[i])
	}
	i++

	adTask.UrlParam, ok = req.Params[i].(string)
	if !ok {
		return fmt.Errorf("UpdateAdTask:ShowPageCtx:Parameter type error:%v", req.Params[i])
	}
	i++

	adTask.SchExp, ok = req.Params[i].(string)
	if !ok {
		return fmt.Errorf("UpdateAdTask:ShowPageCtx:Parameter type error:%v", req.Params[i])
	}
	i++

	adTask.MagExp, ok = req.Params[i].(string)
	if !ok {
		return fmt.Errorf("UpdateAdTask:ShowPageCtx:Parameter type error:%v", req.Params[i])
	}
	i++

	chapter, ok := req.Params[i].(string)
	if !ok {
		return fmt.Errorf("UpdateAdTask:ShowPageCtx:Parameter type error:%v", req.Params[i])
	}
	var err error
	adTask.SchChapt, err = strconv.Atoi(chapter)
	if err != nil {
		return fmt.Errorf("UpdateAdTask:ShowPageCtx:SchChapt type error:%v", err)
	}
	i++

	ani.StorDir, ok = req.Params[i].(string)
	if !ok {
		return fmt.Errorf("UpdateAdTask:ShowPageCtx:Parameter type error:%v", req.Params[i])
	}
	i++

	ani.PlayDir, ok = req.Params[i].(string)
	if !ok {
		return fmt.Errorf("UpdateAdTask:ShowPageCtx:Parameter type error:%v", req.Params[i])
	}
	i++

	animeID, ok := req.Params[i].(string)
	if !ok {
		return fmt.Errorf("UpdateAdTask:ShowPageCtx:Parameter type error:%v", req.Params[i])
	}
	if animeID != "" {
		adTask.AnimeID = bson.ObjectIdHex(animeID)
	}
	i++

	var addAniFlg string
	if len(req.Params) > 10 {
		addAniFlg, ok = req.Params[i].(string)
		if !ok {
			return fmt.Errorf("UpdateAdTask:ShowPageCtx:Parameter type error:%v", req.Params[i])
		}
	}

	if addAniFlg == "true" {
		ani.ID = bson.NewObjectId()
		ani.AnimeID = ani.ID.Hex()
		err := db.SaveAnime(&ani)
		if err != nil {
			return fmt.Errorf("UpdateAdTask:ShowPageCtx:SaveAdTask err:%v", err)
		}
		adTask.AnimeID = ani.ID
	}
	
	err = db.SaveAdTask(&adTask)
	if err != nil {
		return fmt.Errorf("UpdateAdTask:ShowPageCtx:SaveAdTask err:%v", err)
	}

	fmt.Fprint(w, "show_adtask()");
	return nil
}

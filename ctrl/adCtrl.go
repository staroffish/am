package ctrl

import (
	"db"
	"fmt"
	"global"
	"html/template"
	"net/http"
	"strconv"
	"time"
	"view"

	"gopkg.in/mgo.v2/bson"
)

// MainCtrl 主页逻辑控制
type AdCtrl struct {
	view.AdTaskPage
	view.EditAdTaskPage
}

// Init - 初始化
func (a *AdCtrl) Init() error {
	if err := a.AdTaskPage.Init(); err != nil {
		return err
	}
	return a.EditAdTaskPage.Init()

}

type taskData struct {
	Id          string
	AnimeID     string
	AnimeNameJp string
	Url         string
	UrlParam    string
	SchExp      string
	SchChapt    int
	MagExp      string
	StorDir     string
	PlayDir     string
	CheckBox    template.HTML
	Disabled    string
	UpdateTime  time.Time
}

func (a *AdCtrl) Close() {}

// Process 处理函数
func (a *AdCtrl) Process(jr *JSONRequest, w http.ResponseWriter) error {
	defer global.TraceLog("RdCtrl.Process")()

	switch jr.Method {
	case "show_adtask":
		return a.showAdTask(w)
	case "edit_adtask":
		var td taskData

		if len(jr.Params) < 1 {
			return fmt.Errorf("RdCtrl.Process:Parameter num is less then 1")
		}

		id, ok := jr.Params[0].(string)
		if !ok {
			return fmt.Errorf("RdCtrl.Process:show edit page:Id type error:%T", jr.Params[0])
		}
		// 添加自动下载
		if id != "" {
			//显示编辑页面

			tasks := db.GetAdTaskByKey(bson.M{"_id": bson.ObjectIdHex(id)})
			if tasks == nil {
				return fmt.Errorf("RdCtrl.Process:tasks not found:id:%s", id)
			}
			td.Id = tasks[0].Id.Hex()
			td.Url = tasks[0].Url
			td.UrlParam = tasks[0].UrlParam
			td.SchExp = tasks[0].SchExp
			td.SchChapt = tasks[0].SchChapt
			td.MagExp = tasks[0].MagExp
			td.AnimeID = tasks[0].AnimeID.Hex()

			anime := db.GetAnimeByKey(bson.M{"_id": tasks[0].AnimeID}, 0)
			if anime == nil {
				return fmt.Errorf("RdCtrl.Process:anime of tasks not found:id:%s", tasks[0].AnimeID.Hex())
			}
			td.AnimeNameJp = anime[0].AnimeNameJp
			td.StorDir = anime[0].StorDir
			td.PlayDir = anime[0].PlayDir
			td.Disabled = "disabled"
		} else {
			td.SchChapt = 0
			td.CheckBox = template.HTML(`<tr><td>是否创建动漫选项</td><td><input type="checkbox" checked="true" /> </td></tr>`)
		}
		return a.EditAdTaskPage.ShowPage(td, w)
	case "update_adTask":
		return a.updateAdTask(jr, w)
	case "delete_adTask":
		if len(jr.Params) < 1 {
			return fmt.Errorf("RdCtrl.Process:ShowPageCtx:Parameter Num is less then 1")
		}

		_id, ok := jr.Params[0].(string)
		if !ok {
			return fmt.Errorf("RdCtrl.Process:ShowPageCtx:Parameter type error:%v", jr.Params[0])
		}

		err := db.DeleteAdTask(bson.ObjectIdHex(_id))
		if err != nil {
			return fmt.Errorf("RdCtrl.Process:ShowPageCtx:DeletedTask err:%v", err)
		}
		return a.showAdTask(w)
	}

	return fmt.Errorf("RdCtrl.Process:Unsupport Method of %s", jr.Method)
}

func (a *AdCtrl) showAdTask(w http.ResponseWriter) error {
	defer global.TraceLog("RdCtrl.showAdTask")()
	tds := []taskData{}

	tasks := db.GetAdTaskByKey(bson.M{})
	if tasks != nil {
		animeList := db.GetAnimeByKey(bson.M{}, 0)
		if animeList != nil {
			var animeMap = make(map[string]*db.Anime)
			for i := 0; i < len(animeList); i++ {
				animeList[i].ImageBin = nil
				animeMap[animeList[i].AnimeID] = &animeList[i]
			}
			for _, task := range tasks {
				var td taskData
				td.Id = task.Id.Hex()
				td.SchChapt = task.SchChapt
				anime, ok := animeMap[task.AnimeID.Hex()]
				if !ok {
					global.Log.Debugf("am:db.AdTaskPage:animeID in adtask but not in anime:%s", task.AnimeID.Hex())
					continue
				}
				td.StorDir = anime.StorDir
				td.AnimeNameJp = anime.AnimeNameJp
				td.UpdateTime = task.UpdateTime
				tds = append(tds, td)
			}
		}
	}
	return a.AdTaskPage.ShowPage(tds, w)
}

func (a *AdCtrl) updateAdTask(req *JSONRequest, w http.ResponseWriter) error {
	defer global.TraceLog("RdCtrl.updateAdTask")()
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
	adTask.Id = bson.ObjectIdHex(_id)
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

	return a.showAdTask(w)
}

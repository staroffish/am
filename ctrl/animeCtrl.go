package ctrl

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/staroffish/am/global"

	"github.com/staroffish/am/db"
	"github.com/staroffish/am/view"

	"gopkg.in/mgo.v2/bson"
)

// AnimeCtrl 动漫页面控制类
type AnimeCtrl struct {
	view.AnimePage
	view.AnimeCollectionPage
	view.EditAnimePage
	view.MainPage
}

type collection struct {
	Anime []db.Anime
	Key   string
}

type showAnime struct {
	Anime   *db.Anime
	Chapter []struct {
		FullPath string
		FileName string
	}
	PrePage template.JS
}

// Init - 初始化
func (a *AnimeCtrl) Init() error {
	if err := a.AnimePage.Init(); err != nil {
		return err
	}
	if err := a.AnimeCollectionPage.Init(); err != nil {
		return err
	}
	return a.EditAnimePage.Init()
}

func (a *AnimeCtrl) Close() {}

// Process 处理函数
func (a *AnimeCtrl) Process(req *JSONRequest, w http.ResponseWriter) error {
	defer global.TraceLog("AnimeCtrl.Process")()

	switch req.Method {
	// 显示已收藏动漫页面
	case "show_collection":
		err := a.showCollection(req, w)
		if err != nil {
			return err
		}
		break
	// 显示动漫页面
	case "show_anime":
		err := a.showAnime(req, w)
		if err != nil {
			return err
		}
		break
	// 显示动漫编辑/添加页面
	case "edit_anime":
		if len(req.Params) < 2 {
			return fmt.Errorf("AnimeCtrl:Process:edit_anime:Parameter Num is less then 2")
		}
		_id, ok := req.Params[0].(string)
		if !ok {
			return fmt.Errorf("AnimeCtrl:Process:edit_anime:Parameter type error:%v", req.Params[0])
		}

		prePage, ok := req.Params[1].(string)
		if !ok {
			return fmt.Errorf("AnimeCtrl:Process:edit_anime:Parameter type error:%T", req.Params[0])
		}

		var sa showAnime
		if _id == "" {
			// 新增动漫页面所以没有参数直接返回页面
			sa.Anime = &db.Anime{}
			year, season := global.GetNowSeason()
			defaultDir := fmt.Sprintf("%s%02d%02d/", global.Cfg.AnimeDefaultDirPre, year, season)
			sa.Anime.StorDir = defaultDir
			sa.Anime.PlayDir = defaultDir
			sa.PrePage = template.JS(prePage)
		} else {
			// 编辑动漫页面
			ani := db.GetAnime(bson.ObjectIdHex(_id))
			if ani == nil {
				// 已经做了错误应答并出了ERROR LOG 所以 直接返回NIL
				http.Error(w, "Page Not Found", http.StatusNotFound)
				return nil
			}
			sa.Anime = ani
			sa.PrePage = template.JS(prePage)
		}
		if err := a.EditAnimePage.ShowPage(sa, w); err != nil {
			return fmt.Errorf("AnimeCtrl:Process:edit_anime.Execute:%v", err)
		}
		break
	case "update_anime":
		// 更新或添加动漫
		err := a.updateAnime(req, w)
		if err != nil {
			return err
		}
		break
	case "del_anime":
		if len(req.Params) < 1 {
			return fmt.Errorf("AnimeCtrl:Process:del_anime:Parameter Num is less then 1")
		}

		_id, ok := req.Params[0].(string)
		if !ok {
			return fmt.Errorf("AnimeCtrl:Process:del_anime:Parameter type error:%v", req.Params[0])
		}

		err := db.DeletedAnime(bson.ObjectIdHex(_id))
		if err != nil {
			return fmt.Errorf("AnimeCtrl:Process:del_anime:DeletedTask err:%v", err)
		}
		req.Params[0] = ""
		err = a.showCollection(req, w)
		if err != nil {
			return err
		}
		break
	default:
		return fmt.Errorf("AnimeCtrl:Process:unsupport method")

	}

	return nil
}

func (a *AnimeCtrl) showCollection(req *JSONRequest, w http.ResponseWriter) error {
	defer global.TraceLog("AnimeCtrl.showCollection")()
	var col collection
	var key = bson.M{}
	// 带查询条件
	if len(req.Params) >= 1 {
		keyword, ok := req.Params[0].(string)
		if !ok {
			return fmt.Errorf("AnimeCtrl:showCollection:Parameter type error:keyword:%T", req.Params[0])
		}
		col.Key = keyword
		if len(keyword) > 0 {
			key = bson.M{"$or": []bson.M{
				bson.M{"animenamecn": bson.RegEx{Pattern: keyword}},
				bson.M{"animenamejp": bson.RegEx{Pattern: keyword}},
				bson.M{"cast": bson.RegEx{Pattern: keyword}},
				bson.M{"serialsduri": bson.RegEx{Pattern: keyword}},
				bson.M{"type": bson.RegEx{Pattern: keyword}},
				bson.M{"status": bson.RegEx{Pattern: keyword}},
				bson.M{"stordir": bson.RegEx{Pattern: keyword}},
			}}
		}
	}

	col.Anime = db.GetAnimeByKey(key, 0)
	if col.Anime == nil {
		// 已经做了错误应答并出了ERROR LOG 所以 直接返回NIL
		retStr := fmt.Sprintf(`<html><body>动漫未找到<input type="button" value="返回" onclick="javascript:show_collection()"/></body></html>`)
		http.Error(w, retStr, http.StatusNotFound)
		return nil
	}
	err := a.AnimeCollectionPage.ShowPage(col, w)
	if err != nil {
		return fmt.Errorf("AnimeCtrl:showCollection:%v", err)
	}
	return nil
}

func (a *AnimeCtrl) showAnime(req *JSONRequest, w http.ResponseWriter) error {
	var shower showAnime
	if len(req.Params) < 2 {
		return fmt.Errorf("AnimeCtrl:Process:show_anime:Parameter Num is zero")
	}

	_id, ok := req.Params[0].(string)
	if !ok {
		return fmt.Errorf("AnimeCtrl:Process:show_anime:Parameter type error:%T", req.Params[0])
	}

	if _id == "" {
		// id 为空来自添加页面 所以不存在id
		req.Params = req.Params[0:0]
		return a.showCollection(req, w)
	}

	ani := db.GetAnime(bson.ObjectIdHex(_id))
	if ani == nil {
		// 已经做了错误应答并出了ERROR LOG 所以 直接返回NIL
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return nil
	}

	shower.Anime = ani

	prePage, ok := req.Params[1].(string)
	if !ok {
		return fmt.Errorf("AnimeCtrl:Process:show_anime:Parameter type error:%T", req.Params[1])
	}

	shower.PrePage = template.JS(prePage)
	// 取得文件夹下面的所有文件
	fileList, err := ioutil.ReadDir(shower.Anime.StorDir)
	if err == nil {
		shower.Chapter = []struct {
			FullPath string
			FileName string
		}{}
		for _, file := range fileList {
			playPath := fmt.Sprintf("%s/%s", shower.Anime.PlayDir, file.Name())
			shower.Chapter = append(shower.Chapter,
				struct {
					FullPath string
					FileName string
				}{playPath, file.Name()})
		}
	} else {
		global.Log.Infof("AnimeCtrl:Process:show_anime:open or read dir error:%s:%v", shower.Anime.StorDir, err)
	}
	err = a.AnimePage.ShowPage(shower, w)
	if err != nil {
		return fmt.Errorf("AnimeCtrl:Process:show_anime:%v", err)
	}
	return nil
}

func (a *AnimeCtrl) updateAnime(req *JSONRequest, w http.ResponseWriter) error {
	defer global.TraceLog("AnimeCtrl.updateAnime")()

	if len(req.Params) < 2 {
		return fmt.Errorf("AnimeCtrl:updateAnime:Parameter Num is less then 2")
	}

	prePage, ok := req.Params[1].(string)
	if !ok {
		return fmt.Errorf("AnimeCtrl:updateAnime:Parameter type error:%T", req.Params[1])
	}

	_id, ok := req.Params[0].(string)
	if !ok {
		return fmt.Errorf("AnimeCtrl:updateAnime:Parameter type error:%v", req.Params[0])
	}

	var ani db.Anime
	for _, param := range req.Params[2:] {
		_, ok := param.(string)
		if !ok {
			return fmt.Errorf("AnimeCtrl:updateAnime:Parameter type error:value:%v:type:%T", param, param)
		}
	}

	var i = 2
	// 添加新项目时ani.AnimeID不存在,生成一个新的object id
	if _id == "" {
		_id = bson.NewObjectId().Hex()
		prePage = ""
	}
	ani.ID = bson.ObjectIdHex(_id)
	ani.AnimeNameCn, ok = req.Params[i].(string)
	i++
	ani.AnimeNameJp, ok = req.Params[i].(string)
	i++
	ani.Cast, ok = req.Params[i].(string)
	i++
	ani.Type, ok = req.Params[i].(string)
	i++
	ani.Status, ok = req.Params[i].(string)
	i++
	ani.SerialsDuri, ok = req.Params[i].(string)
	i++
	ani.StorDir, ok = req.Params[i].(string)
	i++
	ani.PlayDir, ok = req.Params[i].(string)
	i++
	imageUrl, ok := req.Params[i].(string)
	i++

	// 如果图片路径存在则取得图片
	if imageUrl != "" {
		resp, err := http.Get(imageUrl)
		if err != nil {
			global.Log.Errorf("am:AnimeCtrl:updateAnime:Get image error:%v", err)
		} else {
			defer resp.Body.Close()
			ani.ImageBin, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				global.Log.Errorf("am:AnimeCtrl:updateAnime:Read image error:%v", err)
			}
		}
	}

	// 更新数据库
	err := db.SaveAnime(&ani)
	if err != nil {
		return fmt.Errorf("AnimeCtrl:updateAnime:SaveAnime:%v", err)
	}

	var jr JSONRequest

	if prePage == "" {
		jr.Method = "show_collection"
		err = a.showCollection(&jr, w)
	} else {
		jr.Method = "show_anime"
		jr.Params = []interface{}{_id, prePage}
		err = a.showAnime(&jr, w)
	}
	return err
}

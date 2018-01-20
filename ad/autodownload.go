package ad

import (
	"rd"
	_ "rd/deluge"
	"io"
	"bytes"
	"strings"
	"db"
	"fmt"
	"global"
	"net/http"
	"regexp"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Ad - 自动下载
type Ad struct {
	config *global.Config
	adTask []db.AdTask
	aniMap map[bson.ObjectId]db.Anime
}

// New 创建 自动下载任务
func New(cfg *global.Config) *Ad {
	defer global.TraceLog("ad.New")()
	ad := Ad{config: cfg}
	rd.InitDownloader()
	return &ad
}

func (ad *Ad) refreshData() error {
	defer global.TraceLog("Ad.updateData")()

	// 取得所有的自动下载任务
	ad.adTask = db.GetAdTaskByKey(bson.M{})
	if ad.adTask == nil {
		return nil
	}

	// 取得动漫项
	idList := []bson.ObjectId{}
	for _, task := range ad.adTask {
		idList = append(idList, task.AnimeID)
	}

	ad.aniMap = make(map[bson.ObjectId]db.Anime)
	animeList := db.GetAnimeByKey(bson.M{"_id": bson.M{"$in": idList}}, 0)
	if animeList == nil {
		ad.adTask = nil
		return nil
	}
	for _, anime := range animeList{
		ad.aniMap[anime.ID] = anime
	}

	return nil
}

// Run - 开始自动下载
func (ad *Ad) Run() {
	defer global.TraceLog("Ad.Run")()
	for{
		time.Sleep(time.Duration(ad.config.AdInter) * time.Second)
		if err := ad.refreshData(); err != nil {
			global.Log.Errorf("am:ad.updateData error:%v", err)
			return;
		}
		for _, t := range ad.adTask{
			if strings.HasSuffix(t.Url, "/") == false {
				t.Url += "/"
			}
		
			t.SchChapt++
			// 取得单集的页面
			requestUrl := fmt.Sprintf("%s%s", t.Url, t.UrlParam)
			schExp := fmt.Sprintf(t.SchExp, t.SchChapt)
		
			webCtx, err := ad.getCtxFromWeb(requestUrl, schExp)
			if err != nil {
				global.Log.Errorf("am:ad.getCtxFromWeb error:%v", err)
				continue;
			} 
		
			// 从单集页面去的磁链
			requestUrl = fmt.Sprintf("%s%s", t.Url, webCtx[1])
			webCtx, err = ad.getCtxFromWeb(requestUrl, t.MagExp)
			if err != nil {
				global.Log.Errorf("am:ad.getCtxFromWeb error:%v", err)
				continue;
			} 
		
			// 提交任务
			rdTask := rd.RdTask{}
			rdTask.Link = webCtx[1]
			anime, ok := ad.aniMap[t.AnimeID]
			if !ok {
				global.Log.Errorf("am:ad.Run:get anime of task is not exist:%s", t.Id.Hex())
				continue;
			}

			rdTask.SavePath = anime.StorDir

			err = rd.AddTask(&rdTask, "magnet")
			if err != nil {
				global.Log.Errorf("am:ad.Run:AddTask error:%v", err)
				continue;
			}

			err = db.SaveAdTask(&t)
			if err != nil {
				global.Log.Errorf("am:ad.Run:SaveAdTask error:%v", err)
				continue;
			}

			err = db.UpdateAnimeTime(t.AnimeID)
			if err != nil {
				global.Log.Errorf("am:ad.Run:UpdateAnimeTime err:%v", err)
				continue;
			}
			
			global.Log.Infof("am:ad.Run:OK:%q", t)
		}
	}
}

// getLinkFromWeb 通过正则表达式获取网页上的内容
func (ad *Ad) getCtxFromWeb(url,schExp string) ([]string, error) {
	defer global.TraceLog("ad.getCtxFromWeb")()

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http request error:url=%s,status=%d",
			url, resp.StatusCode)
	}

	reg, err := regexp.Compile(schExp)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer

	_, err = io.Copy(&buf, resp.Body)
	if err != nil {
		return nil, err
	}

	findList := reg.FindStringSubmatch(buf.String())
	if findList == nil {
		return nil, fmt.Errorf("match error:url=%s,exp=%s", url, schExp); 
	}

	return findList, nil
}
package ad

import (
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
	adTask []task
	aniMap map[bson.ObjectId]db.Anime
}

type task db.AdTask

// New 创建 自动下载任务
func New(cfg *global.Config) *Ad {
	defer global.TraceLog("ad.New")()
	ad := Ad{config: cfg}
	return &ad
}

func (ad *Ad) refreshData() error {
	defer global.TraceLog("Ad.updateData")()
	ad.adTask = make([]task, 0)
	c := db.DB.C("adtask")

	// 取得所有的自动下载任务
	err := c.Find(bson.M{}).All(&ad.adTask)
	if err != nil {
		return err
	}

	if len(ad.adTask) == 0 {
		return nil
	}

	idList := []bson.ObjectId{}
	for _, task := range ad.adTask {
		idList = append(idList, task.AnimeID)
	}
	ad.aniMap = make(map[bson.ObjectId]db.Anime)
	c = db.DB.C("anime")
	it := c.Find(bson.M{"_id": bson.M{"$in": idList}}).Iter()
	for {
		var anime db.Anime
		if !it.Next(&anime) {
			if err = it.Err(); err != nil {
				return err
			}
			break
		}
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
			
			fmt.Printf("magnet:%s\n", webCtx[1]);
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
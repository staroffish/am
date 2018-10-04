package ad

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/sclevine/agouti"
	"github.com/staroffish/am/db"
	"github.com/staroffish/am/global"

	"github.com/staroffish/am/rd"

	"gopkg.in/mgo.v2/bson"
)

// Ad - 自动下载
type Ad struct {
	manualAction chan struct{}
	config       *global.Config
	adTask       []db.AdTask
	aniMap       map[bson.ObjectId]db.Anime
	cookies      []*http.Cookie
}

// New 创建 自动下载
func New(cfg *global.Config) (*Ad, error) {
	defer global.TraceLog("ad.New")()
	ad := Ad{config: cfg}
	if err := rd.InitDownloader(); err != nil {
		return nil, err
	}

	ad.manualAction = make(chan struct{})

	ad.cookies = []*http.Cookie{}

	return &ad, nil
}

// 手动查看是否有可以下载的动画
func (ad *Ad) ManualCheckDownload() {
	go func() { ad.manualAction <- struct{}{} }()
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
	for _, anime := range animeList {
		ad.aniMap[anime.ID] = anime
	}

	return nil
}

// Run - 开始自动下载
func (ad *Ad) Run() {
	defer global.TraceLog("Ad.Run")()
	ticker := time.Tick(time.Duration(ad.config.AdInter) * time.Second)
	ad.ManualCheckDownload()
	for {
		select {
		case <-ticker:
			ad.cookies = []*http.Cookie{}
			// 防止在设置manualAction时候卡死
			_, _ = <-ad.manualAction
			global.Log.Infof("Start autodownload.")
		case <-ad.manualAction:
			global.Log.Infof("Start manual action.")
		}

		if err := ad.refreshData(); err != nil {
			global.Log.Errorf("am:ad.refreshData error:%v", err)
			continue
		}
		urlList := make(map[string][]*db.AdTask)
		// 将URL相同的任务归类
		for i, _ := range ad.adTask {
			if strings.HasSuffix(ad.adTask[i].Url, "/") == false {
				ad.adTask[i].Url += "/"
			}
			requestUrl := fmt.Sprintf("%s%s", ad.adTask[i].Url, ad.adTask[i].UrlParam)
			urlList[requestUrl] = append(urlList[requestUrl], &(ad.adTask[i]))
		}

		for url, taskList := range urlList {
			// 取得页面
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				global.Log.Errorf("am:ad.Run:http.NewRequest:%v", err)
				continue
			}

			// 如果有cookie就添加到request中
			for _, cookie := range ad.cookies {
				req.AddCookie(cookie)
			}

			// 由于go自己的 user-agent貌似被对方屏蔽了 所以 这里改成firefox的user-agent
			req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:59.0) Gecko/20100101 Firefox/59.0")
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				global.Log.Errorf("am:ad.Run:http.Get:%v", err)
				continue
			}

			if resp.StatusCode != http.StatusOK {
				if resp.StatusCode == http.StatusServiceUnavailable {
					// 可能网站存在浏览器检查,调用webdriver取得cookie
					err = ad.getCheckBrowserPageCookies(url)
					if err != nil {
						global.Log.Errorf("am:ad.getCheckBrowserPageCookies:%v", err)
						continue
					}
					ad.ManualCheckDownload()
					break
				}
				resp.Body.Close()
				global.Log.Errorf("am:ad.Run:http request error:url=%s,status=%d",
					url, resp.StatusCode)
				continue
			}

			var buf bytes.Buffer

			_, err = io.Copy(&buf, resp.Body)
			resp.Body.Close()
			if err != nil {
				global.Log.Errorf("am:ad.Run:io.Copy error:url=%s,status=%d",
					url, resp.StatusCode)
				continue
			}

			httpCtx := buf.String()

			for _, t := range taskList {
				// 最新集数+1
				t.SchChapt++

				schExp := fmt.Sprintf(t.SchExp, t.SchChapt)

				reg, err := regexp.Compile(schExp)
				if err != nil {
					global.Log.Errorf("am:ad.Run:regexp.Compile error:%v", err)
					continue
				}

				// 匹配该集动漫的链接
				findList := reg.FindStringSubmatch(httpCtx)
				if findList == nil && len(findList) < 2 {
					global.Log.Errorf("am:ad.Run:match error:url=%s,exp=%s", url, schExp)
					continue
				}

				// 抓取到的链接可能是多集的如[1-2]这种情况
				if len(findList) > 2 {
					chapter, err := strconv.Atoi(findList[2])
					if err == nil {
						t.SchChapt = chapter
					}
				}

				global.Log.Infof("am:ad.Run:match ok:exp=%s,chapter=%d", schExp, t.SchChapt)

				// 取得单集页面 并从中取得磁连
				requestUrl := fmt.Sprintf("%s%s", t.Url, findList[1])
				webCtx, err := ad.getMagnetFromWeb(requestUrl, t.MagExp)
				if err != nil {
					global.Log.Errorf("am:ad.getCtxFromWeb error:%v", err)
					continue
				}

				// 提交任务
				rdTask := rd.RdTask{}
				rdTask.Link = webCtx[1]
				anime, ok := ad.aniMap[t.AnimeID]
				if !ok {
					global.Log.Errorf("am:ad.Run:get anime of task is not exist:%s", t.Id.Hex())
					continue
				}

				rdTask.SavePath = anime.StorDir
				rdTask.TaskType = "magnet"

				err = rd.AddTask(&rdTask)
				if err != nil {
					global.Log.Errorf("am:ad.Run:AddTask error:%v", err)
					continue
				}

				err = db.SaveAdTask(t)
				if err != nil {
					global.Log.Errorf("am:ad.Run:SaveAdTask error:%v", err)
					continue
				}

				err = db.UpdateAnimeTime(t.AnimeID)
				if err != nil {
					global.Log.Errorf("am:ad.Run:UpdateAnimeTime err:%v", err)
					continue
				}
				global.Log.Infof("am:ad.Run:OK:%q", *t)
			}
		}
	}
}

// getLinkFromWeb 通过正则表达式获取网页上的内容
func (ad *Ad) getMagnetFromWeb(url, schExp string) ([]string, error) {
	defer global.TraceLog("ad.getCtxFromWeb")()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:59.0) Gecko/20100101 Firefox/59.0")

	// 如果有cookie就添加到request中
	for _, cookie := range ad.cookies {
		req.AddCookie(cookie)
	}

	// resp, err := http.Get(url)
	resp, err := http.DefaultClient.Do(req)

	// resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	// 因为如果发生浏览器检查时会在取得动漫列表时就失败,
	// 所以这里不做浏览器检查防护
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
	if findList == nil && len(findList) < 2 {
		return nil, fmt.Errorf("match error:url=%s,exp=%s", url, schExp)
	}

	return findList, nil
}

func (ad *Ad) getCheckBrowserPageCookies(url string) error {
	driver := agouti.ChromeDriver(agouti.ChromeOptions("args", []string{
		"--user-agent=Mozilla/5.0 (X11; Linux x86_64; rv:59.0) Gecko/20100101 Firefox/59.0",
		"--headless"}))

	if err := driver.Start(); err != nil {
		log.Fatalf("start error:%v", err)
		return fmt.Errorf("driver.Start error:%v", err)
	}
	defer driver.Stop()

	capa := agouti.NewCapabilities().Browser("chrome").With("javascriptEnabled").
		With("databaseEnabled").With("locationContextEnabled").With("applicationCacheEnabled").With("browserConnectionEnabled").With("cssSelectorsEnabled")
	page, err := driver.NewPage(agouti.Desired(capa))
	if err != nil {
		return fmt.Errorf("driver.NewPage error:%v", err)
	}

	if err := page.Navigate(url); err != nil {
		log.Fatalf("page.Navigate error:%v", err)
	}

	time.Sleep(15 * time.Second)

	cookies, err := page.GetCookies()
	if err != nil {
		log.Fatalf("GetCookies error:%v", err)
	}

	for _, cookie := range cookies {
		global.Log.Infof("Get check broser page cookie:name=%s, value=%s", cookie.Name, cookie.Value)
		ad.cookies = append(ad.cookies, cookie)
	}

	return nil
}

package ad

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/sclevine/agouti"
	"github.com/staroffish/am/db"
	"github.com/staroffish/am/global"

	"github.com/staroffish/am/rd"

	"github.com/go-redis/redis/v8"
	"gopkg.in/mgo.v2/bson"
)

const DayDuration = 24 * time.Hour

// Ad - 自动下载
type Ad struct {
	manualAction chan struct{}
	config       *global.Config
	adTask       []db.AdTask
	aniMap       map[bson.ObjectId]db.Anime
	cookies      map[string][]*http.Cookie
	rdb          *redis.Client
	adTaskMutex  *sync.Mutex
}

// New 创建 自动下载
func New(cfg *global.Config) (*Ad, error) {
	defer global.TraceLog("ad.New")()
	ad := Ad{config: cfg}
	if err := rd.InitDownloader(); err != nil {
		return nil, err
	}

	rdb, err := connectRedis(ad.config.RedisAddr, ad.config.RedisPasswd)
	if err != nil {
		global.Log.Errorf("ad:Run:connectRedis error:%v", err)
		log.Fatalf("ad:Run:connectRedis error:%v", err)
	}

	ad.manualAction = make(chan struct{}, 2)

	ad.cookies = map[string][]*http.Cookie{}
	ad.adTaskMutex = &sync.Mutex{}
	ad.rdb = rdb

	return &ad, nil
}

// 手动查看是否有可以下载的动画
func (ad *Ad) ManualCheckDownload() {
	go func() { ad.manualAction <- struct{}{} }()
}

func (ad *Ad) refreshData() error {
	defer global.TraceLog("Ad.updateData")()

	ad.adTaskMutex.Lock()
	defer ad.adTaskMutex.Unlock()
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

func (ad *Ad) getAdTasks() []db.AdTask {
	ad.adTaskMutex.Lock()
	defer ad.adTaskMutex.Unlock()

	adtasks := make([]db.AdTask, len(ad.adTask))
	copy(adtasks, ad.adTask)
	return adtasks
}

// Run - 开始自动下载
func (ad *Ad) Run() {
	defer global.TraceLog("Ad.Run")()
	crawlerTicker := time.Tick(time.Duration(ad.config.AdInter) * time.Second)
	MatchingTicker := time.Tick(15 * time.Second)

	if err := ad.refreshData(); err != nil {
		global.Log.Errorf("am:ad.refreshData error:%v", err)
		panic("refresh data error when boot up")
	}

	go func() {
		for {
			select {
			case <-ad.manualAction:
				global.Log.Infof("Start manual action.")
			case <-crawlerTicker:
				ad.cookies = map[string][]*http.Cookie{}
				global.Log.Infof("Start autodownload.")
			}
			adTasks := ad.getAdTasks()
			urlList := make(map[string]struct{})

			global.Log.Infof("ad task count:%d,%d", len(adTasks), len(ad.adTask))
			// 将URL相同的任务归类
			for i, _ := range adTasks {
				if strings.HasSuffix(adTasks[i].Url, "/") == false {
					adTasks[i].Url += "/"
				}
				requestUrl := fmt.Sprintf("%s%s", adTasks[i].Url, adTasks[i].UrlParam)
				urlList[requestUrl] = struct{}{}
			}

			for url, _ := range urlList {

				global.Log.Infof("get url page. url=%s", url)
				// 取得页面
				req, err := http.NewRequest("GET", url, nil)
				if err != nil {
					global.Log.Errorf("am:ad.Run:http.NewRequest:%v", err)
					continue
				}

				// 如果有cookie就添加到request中
				ad.addCookieToRequest(req, url)

				// 由于go自己的 user-agent貌似被对方屏蔽了 所以 这里改成firefox的user-agent
				req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:59.0) Gecko/20100101 Firefox/59.0")
				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					global.Log.Errorf("am:ad.Run:http.Get:%v", err)
					continue
				}

				if resp.StatusCode != http.StatusOK {
					resp.Body.Close()
					if resp.StatusCode == http.StatusServiceUnavailable {
						global.Log.Infof("am:ad.Run:http.Get url %s StatusServiceUnavailable", url)
						// 可能网站存在浏览器检查,调用webdriver取得cookie
						err = ad.getCheckBrowserPageCookies(url)
						if err != nil {
							global.Log.Errorf("am:ad.getCheckBrowserPageCookies:%v", err)
							continue
						}
						ad.ManualCheckDownload()
						break
					}
					global.Log.Errorf("am:ad.Run:http request error:url=%s,status=%d",
						url, resp.StatusCode)
					continue
				}

				global.Log.Infof("am:ad.Run:Get page OK:url=%s", url)

				var buf bytes.Buffer

				_, err = io.Copy(&buf, resp.Body)
				resp.Body.Close()
				if err != nil {
					global.Log.Errorf("am:ad.Run:io.Copy error:url=%s,status=%d",
						url, resp.StatusCode)
					continue
				}

				httpCtx := buf.String()
				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				timestamp := time.Now().Unix()
				pageCacheKey := fmt.Sprintf("pagecache:%s:%d", url, timestamp)
				err = ad.rdb.Set(ctx, pageCacheKey, httpCtx, DayDuration*time.Duration(ad.config.PageCacheDuration)).Err()
				cancel()
				if err != nil {
					global.Log.Errorf("am:set pagecache error:key=%s, err=%v", pageCacheKey, err)
					continue
				}

				ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
				err = ad.rdb.ZAdd(ctx, "pagecache:time", &redis.Z{
					Score:  float64(time.Now().Unix()),
					Member: pageCacheKey,
				}).Err()
				cancel()
				if err != nil {
					global.Log.Errorf("am:add pagecache:time error:member=%s, error=%v", pageCacheKey, err)
					continue
				}
			}
		}
	}()

	go func() {
		for {
			select {
			case <-MatchingTicker:
				ad.cookies = map[string][]*http.Cookie{}
				global.Log.Infof("check page cache.")
			}

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			expiration := time.Now().Add(-1 * DayDuration * time.Duration(ad.config.PageCacheDuration))
			err := ad.rdb.ZRemRangeByScore(ctx, "pagecache:time", "0", fmt.Sprintf("%d", expiration.Unix())).Err()
			cancel()
			if err != nil {
				global.Log.Errorf("am:ZRemRangeByScore pagecache:time error:scope=0-%d, error=%v", expiration.Unix(), err)
			}

			if err := ad.refreshData(); err != nil {
				global.Log.Errorf("am:ad.refreshData error:%v", err)
				continue
			}

			pageCache := map[string][]string{}

			ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
			pageCacheKeys, err := ad.rdb.ZRange(ctx, "pagecache:time", 0, -1).Result()
			cancel()
			if err != nil {
				global.Log.Errorf("am:zrange pagecache:time error:%v", err)
				continue
			}

			for _, key := range pageCacheKeys {
				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				pageCacheCtx, err := ad.rdb.Get(ctx, key).Result()
				cancel()
				if err != nil {
					global.Log.Errorf("am:Get %s error:%v", key, err)
					continue
				}

				splitedKeys := key[strings.Index(key, ":")+1 : strings.LastIndex(key, ":")]
				if len(splitedKeys) < 2 {
					global.Log.Errorf("am:incorrect page cache key %s", key)
					continue
				}
				pageCache[splitedKeys] = append(pageCache[splitedKeys], pageCacheCtx)
			}

			adTasks := ad.getAdTasks()

			for _, t := range adTasks {

				if strings.HasSuffix(t.Url, "/") == false {
					t.Url += "/"
				}
				requestUrl := fmt.Sprintf("%s%s", t.Url, t.UrlParam)

				// 最新集数+1
				t.SchChapt++

				schExp := fmt.Sprintf(t.SchExp, t.SchChapt)

				reg, err := regexp.Compile(schExp)
				if err != nil {
					global.Log.Errorf("am:ad.Run:regexp.Compile error:%v", err)
					continue
				}

				var findList []string = nil
				for _, pageCtx := range pageCache[requestUrl] {
					// 匹配该集动漫的链接
					findList = reg.FindStringSubmatch(pageCtx)
					if findList == nil || len(findList) < 2 {
						continue
					}

					// 抓取到的链接可能是多集的如[1-2]这种情况
					if len(findList) > 2 {
						chapter, err := strconv.Atoi(findList[2])
						if err == nil {
							t.SchChapt = chapter
						}
					}

					break
				}

				if findList == nil || len(findList) < 2 {
					global.Log.Errorf("am:ad.Run:match error:url=%s,exp=%s", requestUrl, schExp)
					continue
				}

				global.Log.Infof("am:ad.Run:match ok:exp=%s,chapter=%d", schExp, t.SchChapt)

				// 取得单集页面 并从中取得磁连
				animeChapterUrl := fmt.Sprintf("%s%s", t.Url, findList[1])
				webCtx, err := ad.getMagnetFromWeb(animeChapterUrl, t.MagExp, requestUrl)
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

				err = db.SaveAdTask(&t)
				if err != nil {
					global.Log.Errorf("am:ad.Run:SaveAdTask error:%v", err)
					continue
				}

				err = db.UpdateAnimeTime(t.AnimeID)
				if err != nil {
					global.Log.Errorf("am:ad.Run:UpdateAnimeTime err:%v", err)
					continue
				}
				global.Log.Infof("am:ad.Run:OK:%q", t)
			}
		}
	}()

	// 在最后出发一次网页抓取
	ad.ManualCheckDownload()
}

// getLinkFromWeb 通过正则表达式获取网页上的内容
func (ad *Ad) getMagnetFromWeb(url, schExp, cookieMapKey string) ([]string, error) {
	defer global.TraceLog("ad.getCtxFromWeb")()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:59.0) Gecko/20100101 Firefox/59.0")

	// 如果有cookie就添加到request中
	ad.addCookieToRequest(req, cookieMapKey)

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
		return fmt.Errorf("page.Navigate error:%v", err)
	}

	time.Sleep(15 * time.Second)

	cookies, err := page.GetCookies()
	if err != nil {
		return fmt.Errorf("GetCookies error:%v", err)
	}

	ad.cookies[url] = []*http.Cookie{}
	for _, cookie := range cookies {
		global.Log.Infof("Get check browser page cookie:url=%s, name=%s, value=%s", url, cookie.Name, cookie.Value)
		ad.cookies[url] = append(ad.cookies[url], cookie)
	}

	return nil
}

func (ad *Ad) addCookieToRequest(req *http.Request, url string) {
	_, hasCookie := ad.cookies[url]
	if hasCookie {
		for _, cookie := range ad.cookies[url] {
			global.Log.Debugf("am:ad.Run:addCookie:url=%s, key=%s,value=%s", url, cookie.Name, cookie.Value)
			req.AddCookie(cookie)
		}
	}
}

func connectRedis(addr, password string) (*redis.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       2, // use default DB
	})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("connect to redis error: %v", err)
	}
	return rdb, nil
}

package ctrl

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/staroffish/am/api/common"
	downloadermanagerv1 "github.com/staroffish/am/api/downloadmanager/v1"
	"github.com/staroffish/am/db"
	clientv3 "go.etcd.io/etcd/client/v3"
	ggrpc "google.golang.org/grpc"

	"github.com/staroffish/am/global"

	"github.com/staroffish/am/view"
	"gopkg.in/mgo.v2/bson"
)

// MainCtrl 主页逻辑控制
type AdCtrl struct {
	view.AdTaskPage
	view.EditAdTaskPage
	downloadermanagerv1.DownloadmanagerClient
}

// Init - 初始化
func (a *AdCtrl) Init() error {
	if err := a.AdTaskPage.Init(); err != nil {
		return err
	}
	return a.EditAdTaskPage.Init()

}

func (a *AdCtrl) InitClient() error {

	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{global.Cfg.EtcdEndpoints},
		DialTimeout: time.Duration(global.Cfg.EtcdDialTimeout) * time.Second,
		DialOptions: []ggrpc.DialOption{ggrpc.WithBlock()},
	})
	if err != nil {
		global.Log.Errorf("ad:Run:connectRedis error:%v", err)
		log.Fatalf("ad:Run:connectRedis error:%v", err)
	}

	r := etcd.New(client)

	connDownloadManager, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///download-manager"),
		grpc.WithDiscovery(r),
	)
	if err != nil {
		global.Log.Errorf("ad:Run:grpc.DialInsecure error:%v", err)
		log.Fatalf("ad:Run:grpc.DialInsecure error:%v", err)
	}

	a.DownloadmanagerClient = downloadermanagerv1.NewDownloadmanagerClient(connDownloadManager)

	return nil
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

			task, err := a.getTask(id)
			if err != nil {
				return fmt.Errorf("RdCtrl.Process:getTask %s error :%v", id, err)
			}

			td.Id = task.Id
			td.Url = task.Url
			td.UrlParam = task.UrlParam
			td.SchExp = task.SchExp
			td.SchChapt = task.SchChapt
			td.MagExp = task.MagExp
			td.AnimeID = task.AnimeID.Hex()

			anime := db.GetAnimeByKey(bson.M{"_id": task.AnimeID}, 0)
			if len(anime) == 0 {
				return fmt.Errorf("RdCtrl.Process:anime of tasks not found:id:%s", task.AnimeID)
			}
			td.AnimeNameJp = task.Name
			td.StorDir = task.StorePath
			td.PlayDir = anime[0].PlayDir
			td.Disabled = "disabled"
		} else {
			td.Url = global.Cfg.DefaultAdUrl
			td.MagExp = global.Cfg.DefaultAdMagExp

			year, season := global.GetNowSeason()

			// year = year - (year / 100 * 100)

			defaultDir := fmt.Sprintf("%s%02d%02d/", global.Cfg.AnimeDefaultDirPre, year, season)
			td.StorDir = defaultDir
			td.PlayDir = defaultDir
			td.SchChapt = 0
			td.CheckBox = template.HTML(`<tr><td>是否创建动漫选项</td><td><input type="checkbox" checked="true" /> </td></tr>`)
		}
		return a.EditAdTaskPage.ShowPage(td, w)
	case "update_adTask":
		return a.updateAdTask(jr, w)
	case "delete_adTask":
		if len(jr.Params) < 1 {
			return fmt.Errorf("RdCtrl.Process:Parameter Num is less then 1")
		}

		_id, ok := jr.Params[0].(string)
		if !ok {
			return fmt.Errorf("RdCtrl.Process:Parameter type error:%v", jr.Params[0])
		}

		isDone, ok := jr.Params[1].(bool)
		if ok && isDone {
			task, err := a.getTask(_id)
			if err != nil {
				return fmt.Errorf("RdCtrl.Process:GetAdTaskByKey error:%v", err)
			}

			err = db.UpdateAnimeDone(task.AnimeID)
			if err != nil {
				return fmt.Errorf("RdCtrl.Process:UpdateAnimeStatus err:%v", err)
			}
		}

		err := a.deleteTask(_id)
		if err != nil {
			return fmt.Errorf("RdCtrl.Process:DeletedTask err:%v", err)
		}
		return a.showAdTask(w)
	}

	return fmt.Errorf("RdCtrl.Process:Unsupport Method of %s", jr.Method)
}

func (a *AdCtrl) showAdTask(w http.ResponseWriter) error {
	defer global.TraceLog("RdCtrl.showAdTask")()
	tds := []taskData{}

	tasks, err := a.listTask()
	if err != nil {
		return fmt.Errorf("AdCtrl.showAdTask: listTask error: %v", err)
	}

	for _, task := range tasks {
		var td taskData
		td.Id = task.Id
		td.SchChapt = task.SchChapt
		td.StorDir = task.StorePath
		td.AnimeNameJp = task.Name
		td.UpdateTime = task.UpdateTime
		td.AnimeID = task.AnimeID.Hex()
		tds = append(tds, td)

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

	adTask.Id = _id
	var i = 1
	ani.AnimeNameJp, ok = req.Params[i].(string)
	if !ok {
		return fmt.Errorf("UpdateAdTask:ShowPageCtx:Parameter type error:%v", req.Params[i])
	}
	adTask.Name = ani.AnimeNameJp
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
	adTask.StorePath = ani.StorDir
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
		ani.Status = "连载中"
		year, season := global.GetNowSeason()

		ani.SerialsDuri = fmt.Sprintf("%04d/%02d~", year, season)

		err := db.SaveAnime(&ani)
		if err != nil {
			return fmt.Errorf("UpdateAdTask:ShowPageCtx:SaveAdTask err:%v", err)
		}
		adTask.AnimeID = ani.ID
	}

	if adTask.Id == "" {
		err = a.saveTask(&adTask)
	} else {
		err = a.updateTask(&adTask)
	}

	if err != nil {
		return fmt.Errorf("UpdateAdTask:ShowPageCtx:saveTask or updateTask err:%v: id = %s", err, adTask.Id)
	}

	return a.showAdTask(w)
}

func (a *AdCtrl) saveTask(adTask *db.AdTask) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	req := &downloadermanagerv1.AddTaskRequest{
		Name:          adTask.Name,
		Regexp:        adTask.SchExp,
		LatestChapter: int32(adTask.SchChapt),
		StorePath:     adTask.StorePath,
		AnimeId:       adTask.AnimeID.Hex(),
	}

	if _, err := a.DownloadmanagerClient.AddTask(ctx, req); err != nil {
		return err
	}
	return nil
}

func (a *AdCtrl) updateTask(adTask *db.AdTask) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id, err := strconv.ParseInt(adTask.Id, 10, 64)
	if err != nil {
		return fmt.Errorf("ad task id is not int32 %v", adTask.Id)
	}

	req := &downloadermanagerv1.UpdateTaskRequest{
		Id:            int32(id),
		Name:          adTask.Name,
		Regexp:        adTask.SchExp,
		LatestChapter: int32(adTask.SchChapt),
		StorePath:     adTask.StorePath,
		AnimeId:       adTask.AnimeID.Hex(),
	}

	if _, err := a.DownloadmanagerClient.UpdateTask(ctx, req); err != nil {
		return err
	}
	return nil
}

func (a *AdCtrl) deleteTask(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	taskId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return fmt.Errorf("ad task id is not int32 %v", id)
	}

	req := &downloadermanagerv1.DeleteTaskRequest{
		Id: int32(taskId),
	}

	if _, err := a.DownloadmanagerClient.DeleteTask(ctx, req); err != nil {
		return err
	}
	return nil
}

func (a *AdCtrl) listTask() ([]*db.AdTask, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := a.DownloadmanagerClient.ListTask(ctx, &common.Empty{})
	if err != nil {
		return nil, err
	}

	adTasks := []*db.AdTask{}

	for _, task := range resp.Tasks {
		updateTime, err := time.Parse("2006-01-02 15:04:05", task.UpdateTime)
		if err != nil {
			return nil, err
		}
		adTask := &db.AdTask{
			Id:         fmt.Sprintf("%d", task.Id),
			AnimeID:    bson.ObjectIdHex(task.AnimeId),
			SchExp:     task.Regexp,
			StorePath:  task.StorePath,
			SchChapt:   int(task.LatestChapter),
			Name:       task.Name,
			UpdateTime: updateTime,
		}

		adTasks = append(adTasks, adTask)
	}
	return adTasks, nil
}

func (a *AdCtrl) getTask(idStr string) (*db.AdTask, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return nil, err
	}

	req := &downloadermanagerv1.GetTaskRequest{
		Id: int32(id),
	}

	resp, err := a.DownloadmanagerClient.GetTask(ctx, req)
	if err != nil {
		return nil, err
	}

	updateTime, err := time.Parse("2006-01-02 15:04:05", resp.Task.UpdateTime)
	if err != nil {
		return nil, err
	}

	adTask := &db.AdTask{
		Id:         fmt.Sprintf("%d", resp.Task.Id),
		AnimeID:    bson.ObjectIdHex(resp.Task.AnimeId),
		SchExp:     resp.Task.Regexp,
		StorePath:  resp.Task.StorePath,
		SchChapt:   int(resp.Task.LatestChapter),
		Name:       resp.Task.Name,
		UpdateTime: updateTime,
	}

	return adTask, nil
}

package ammicro

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/staroffish/am/api/common"
	downloaderv1 "github.com/staroffish/am/api/downloader/v1"
	"github.com/staroffish/am/global"
	"github.com/staroffish/am/rd"
	clientv3 "go.etcd.io/etcd/client/v3"
	ggrpc "google.golang.org/grpc"
)

type AmMicroDownloader struct {
	downloader downloaderv1.DownloaderClient
}

// TaskType - 下载类型字符串
var taskType = "magnet"

func init() {
	fmt.Printf("ammicro.init\n")
	var dnldr AmMicroDownloader
	rd.SetDownloader(taskType, &dnldr)
}

func (a *AmMicroDownloader) InitDownloader(cfg *global.Config) error {
	global.Log.Infof("%s", "AmMicroDownloader.InitDownloader")
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{cfg.EtcdEndpoints},
		DialTimeout: time.Duration(cfg.EtcdDialTimeout) * time.Second,
		DialOptions: []ggrpc.DialOption{ggrpc.WithBlock()},
	})
	if err != nil {
		global.Log.Errorf("AmMicroDownloader.InitDownloader:clientv3.New error:%v", err)
		log.Fatalf("AmMicroDownloader.InitDownloader:clientv3.New error:%v", err)
	}

	r := etcd.New(client)

	connDownloader, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///downloader"),
		grpc.WithDiscovery(r),
		grpc.WithTimeout(10*time.Second),
	)
	if err != nil {
		global.Log.Errorf("AmMicroDownloader.InitDownloader:grpc.DialInsecure error:%v", err)
		log.Fatalf("AmMicroDownloader.InitDownloader:grpc.DialInsecure error:%v", err)
	}

	a.downloader = downloaderv1.NewDownloaderClient(connDownloader)
	return nil
}

func (a *AmMicroDownloader) AddTask(t *rd.RdTask) error {
	defer global.TraceLog("AmMicroDownloader.AmMicroDownloader.AddTask")()

	req := &downloaderv1.AddRequest{
		Link:      t.Link,
		StorePath: t.SavePath,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	_, err := a.downloader.Add(ctx, req)
	cancel()
	if err != nil {
		return fmt.Errorf("AmMicroDownloader.AddTask: add task error: %v", err)
	}

	return nil
}
func (a *AmMicroDownloader) PauseTask(t *rd.RdTask) error {
	defer global.TraceLog("AmMicroDownloader.AmMicroDownloader.PauseTask")()

	req := &downloaderv1.PauseRequest{
		Id: t.Ids,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	_, err := a.downloader.Pause(ctx, req)
	cancel()
	if err != nil {
		return fmt.Errorf("AmMicroDownloader.PauseTask: add task error: %v", err)
	}

	return nil
}
func (a *AmMicroDownloader) ResumeTask(t *rd.RdTask) error {
	defer global.TraceLog("AmMicroDownloader.AmMicroDownloader.ResumeTask")()

	req := &downloaderv1.ResumeRequest{
		Id: t.Ids,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	_, err := a.downloader.Resume(ctx, req)
	cancel()
	if err != nil {
		return fmt.Errorf("AmMicroDownloader.ResumeTask: add task error: %v", err)
	}

	return nil
}
func (a *AmMicroDownloader) DeleteTask(t *rd.RdTask) error {
	defer global.TraceLog("AmMicroDownloader.AmMicroDownloader.DeleteTask")()

	req := &downloaderv1.DeleteRequest{
		Id: t.Ids,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	_, err := a.downloader.Delete(ctx, req)
	cancel()
	if err != nil {
		return fmt.Errorf("AmMicroDownloader.DeleteTask: add task error: %v", err)
	}

	return nil
}
func (a *AmMicroDownloader) GetAllTask() ([]rd.RdTask, error) {
	defer global.TraceLog("deluge.AmMicroDownloader.GetAllTask")()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	resp, err := a.downloader.List(ctx, &common.Empty{})
	cancel()
	if err != nil {
		return nil, fmt.Errorf("AmMicroDownloader.GetAllTask: add task error: %v", err)
	}

	tasks := []rd.RdTask{}
	for _, taskInfo := range resp.TaskInfos {
		task := rd.RdTask{
			Ids:        taskInfo.Id,
			SavePath:   taskInfo.StorePath,
			Name:       taskInfo.Name,
			Progress:   int(taskInfo.Progress * 100),
			State:      taskInfo.Status,
			TaskType:   taskType,
			CreateTime: taskInfo.CreatedTime,
			Size:       taskInfo.Size,
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

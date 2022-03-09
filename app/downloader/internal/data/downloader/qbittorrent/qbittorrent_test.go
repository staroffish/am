package qbittorrent

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/staroffish/am/app/downloader/internal/data/downloader"
	"github.com/staroffish/am/common/util"
)

const (
	QBIT_ADDR     = "http://XXX"
	QBIT_USER     = "XX"
	QBIT_PASSWORD = "XX"
)

func NewTestQbittorrentDownloaderRepo(logger log.Logger) *QbittorrentDownloaderRepo {
	redisClient := redis.NewClient(&redis.Options{})
	repo := &QbittorrentDownloaderRepo{
		log:                log.NewHelper(logger),
		DownloaderBaseRepo: downloader.NewDownloaderBaseRepo("test", redisClient, logger),
		mutex:              &sync.RWMutex{},
		qbittorrentConfig: &qbittorrentConfig{
			url:      QBIT_ADDR,
			userName: QBIT_USER,
			password: QBIT_PASSWORD,
		},
	}

	return repo
}

func TestQbittorrentList(t *testing.T) {
	repo := NewTestQbittorrentDownloaderRepo(util.NewTestLogger())

	taskInfos, err := repo.List(context.Background())
	if err != nil {
		t.Fatalf("repo.List error:%v", err)
	}
	for _, taskInfo := range taskInfos {
		fmt.Printf("taskInfo=%v\n", taskInfo)
	}
}

func TestQbittorentPause(t *testing.T) {
	repo := NewTestQbittorrentDownloaderRepo(util.NewTestLogger())

	taskInfos, err := repo.List(context.Background())
	if err != nil || len(taskInfos) == 0 {
		t.Fatalf("repo.List error or taskinfos is empty:%v", err)
	}
	hash := taskInfos[0].Id
	err = repo.Pause(context.Background(), hash)
	if err != nil {
		t.Fatalf("repo.Pause error:%v", err)
	}

	time.Sleep(time.Second * 1)

	taskInfos, err = repo.List(context.Background())
	if err != nil || len(taskInfos) == 0 {
		t.Fatalf("repo.List error or taskinfos is empty:%v", err)
	}

	for _, taskInfo := range taskInfos {
		if taskInfo.Id == hash {
			if taskInfo.Status != "paused" {
				t.Fatalf("taskInfo %s %s status is not paused", taskInfo.Name, taskInfo.Id)
			}
			break
		}
	}
}

func TestQbittorentDelete(t *testing.T) {
	repo := NewTestQbittorrentDownloaderRepo(util.NewTestLogger())

	hash := "c9ff965fedee0f836f065780916ba610cfea8dcc"
	err := repo.Delete(context.Background(), hash)
	if err != nil {
		t.Fatalf("repo.Delete error:%v", err)
	}

	time.Sleep(time.Second * 1)

	taskInfos, err := repo.List(context.Background())
	if err != nil || len(taskInfos) == 0 {
		t.Fatalf("repo.List error or taskinfos is empty:%v", err)
	}

	for _, taskInfo := range taskInfos {
		if taskInfo.Id == hash {
			t.Fatalf("task %s %s exists yet", taskInfo.Name, taskInfo.Id)
			break
		}
	}
}

func TestQbittorentAdd(t *testing.T) {
	repo := NewTestQbittorrentDownloaderRepo(util.NewTestLogger())

	link := `magnet:?xt=urn:btih:ZH7ZMX7N5YHYG3YGK6AJC25GCDH6VDOM&dn=&tr=http%3A%2F%2F104.238.198.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.238.198.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%2Fannounce&tr=http%3A%2F%2Fopen.nyaatorrents.info%3A6544%2Fannounce&tr=http%3A%2F%2Ft2.popgo.org%3A7456%2Fannonce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Ftracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%3A80%2Fannounce&tr=https%3A%2F%2Ftr.bangumi.moe%3A9696%2Fannounce&tr=http%3A%2F%2Ft.acg.rip%3A6699%2Fannounce&tr=http%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Fopen.acgnxtracker.com%2Fannounce&tr=http%3A%2F%2Fsukebei.tracker.wf%3A8888%2Fannounce&tr=http%3A%2F%2Ftracker.kamigami.org%3A2710%2Fannounce&tr=https%3A%2F%2Ftracker.nanoha.org%2Fannounce`
	storePath := "/store/TV/201907/"
	err := repo.Add(context.Background(), link, storePath)
	if err != nil {
		t.Fatalf("repo.Delete error:%v", err)
	}

	time.Sleep(time.Second * 1)

	taskInfos, err := repo.List(context.Background())
	if err != nil || len(taskInfos) == 0 {
		t.Fatalf("repo.List error or taskinfos is empty:%v", err)
	}

	for _, taskInfo := range taskInfos {
		if taskInfo.Id == "c9ff965fedee0f836f065780916ba610cfea8dcc" {
			return
		}
	}

	t.Fatalf("task add failure")
}

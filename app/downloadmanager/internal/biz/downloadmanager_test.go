package biz

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/go-kratos/kratos/v2/log"
	dtoDownloadManager "github.com/staroffish/am/common/dto/downloadmanager"
	dtoDownloadmanager "github.com/staroffish/am/common/dto/downloadmanager"
	dtoSpider "github.com/staroffish/am/common/dto/spider"
	"github.com/staroffish/am/common/util"
)

type testRepo struct{}

func (r *testRepo) CreateDownloadTask(ctx context.Context, id int32, latestChapter int32, MagnetLink string) error {
	return nil
}

func (r *testRepo) GetAnimeMagnets(ctx context.Context) ([]*dtoSpider.AnimeMagnet, error) {
	return []*dtoSpider.AnimeMagnet{
		{
			Name:       `[Skymoon-Raws] 進擊的巨人 第四季 / Shingeki no Kyojin - The Final Season - 82 [ViuTV][WEB-DL][1080p][AVC AAC][繁體外掛][MP4+ASSx2](正式版本)`,
			MagnetLink: `magnet:[Skymoon-Raws] 進擊的巨人 第四季 / Shingeki no Kyojin - The Final Season - 82 [ViuTV][WEB-DL][1080p][AVC AAC][繁體外掛][MP4+ASSx2](正式版本)`,
		},
		{
			Name:       `【豌豆字幕组】[进击的巨人 / Shingeki_no_Kyojin][83-85][简体][1080P][MP4]`,
			MagnetLink: `magnet:【豌豆字幕组】[进击的巨人 / Shingeki_no_Kyojin][83-85][简体][1080P][MP4]`,
		},
		{
			Name:       `【豌豆字幕組】[進擊的巨人 / Shingeki_no_Kyojin][83-86][繁體][1080P][MP4]`,
			MagnetLink: `magnet:【豌豆字幕組】[進擊的巨人 / Shingeki_no_Kyojin][83-86][繁體][1080P][MP4]`,
		},
		{
			Name:       `【豌豆字幕組】[進擊的巨人 / Shingeki_no_Kyojin][87][繁體][1080P][MP4]`,
			MagnetLink: `magnet:【豌豆字幕組】[進擊的巨人 / Shingeki_no_Kyojin][87][繁體][1080P][MP4]`,
		},
		{
			Name:       `【豌豆字幕组】[进击的巨人 / Shingeki_no_Kyojin][82][简体][1080P][MP4]`,
			MagnetLink: `magnet:【豌豆字幕组】[进击的巨人 / Shingeki_no_Kyojin][82][简体][1080P][MP4]`,
		},
		{
			Name:       `【豌豆字幕組】[進擊的巨人 / Shingeki_no_Kyojin][88][繁體][1080P][MP4]`,
			MagnetLink: `magnet:【豌豆字幕組】[進擊的巨人 / Shingeki_no_Kyojin][88][繁體][1080P][MP4]`,
		},
	}, nil
}

func (r *testRepo) AddTaskInfo(context.Context, *dtoDownloadManager.DownloadTaskInfo) error {
	return nil
}
func (r *testRepo) DeleteTask(ctx context.Context, id int32) error {
	return nil
}

func (r *testRepo) GetDownloadTaskInfos() []*dtoDownloadManager.DownloadTaskInfo {
	taskinfos := []*dtoDownloadManager.DownloadTaskInfo{
		{
			Name:          "進撃の巨人 The Final Season Part2",
			Regexp:        `【豌豆字幕組】\[進擊的巨人 / Shingeki_no_Kyojin\]\[(%02d)(?:.|&amp;)?(\d{2})?(?: ?(完|END|End|Fin|FIN))?(?:v\d)?\](?:\[(完|END|End|Fin|FIN)\])?\[繁體\]\[1080P\]\[MP4\]`,
			LatestChapter: 82,
		},
	}
	return taskinfos
}

func (r *testRepo) UpdateLatestChapter(ctx context.Context, id, latestChapter int32) error {
	return nil
}
func (r *testRepo) UpdateTaskInfo(ctx context.Context, taskInfo *dtoDownloadManager.DownloadTaskInfo) error {
	return nil
}

func (r *testRepo) GetDownloadTaskInfoById(id int32) *dtoDownloadmanager.DownloadTaskInfo {
	return nil
}

func TestScan(t *testing.T) {
	dm := &DownloadManager{
		repo: &testRepo{},
		log:  log.NewHelper(util.NewTestLogger()),
	}

	result := []*DownloadTask{
		{
			TaskName:     `【豌豆字幕組】[進擊的巨人 / Shingeki_no_Kyojin][83-86][繁體][1080P][MP4]`,
			MagnetLink:   `magnet:【豌豆字幕組】[進擊的巨人 / Shingeki_no_Kyojin][83-86][繁體][1080P][MP4]`,
			ChapterStart: 83,
			ChapterEnd:   86,
		},
		{
			TaskName:     `【豌豆字幕組】[進擊的巨人 / Shingeki_no_Kyojin][87][繁體][1080P][MP4]`,
			MagnetLink:   `magnet:【豌豆字幕組】[進擊的巨人 / Shingeki_no_Kyojin][87][繁體][1080P][MP4]`,
			ChapterStart: 87,
			ChapterEnd:   87,
		},
		{
			TaskName:     `【豌豆字幕組】[進擊的巨人 / Shingeki_no_Kyojin][88][繁體][1080P][MP4]`,
			MagnetLink:   `magnet:【豌豆字幕組】[進擊的巨人 / Shingeki_no_Kyojin][88][繁體][1080P][MP4]`,
			ChapterStart: 88,
			ChapterEnd:   88,
		},
	}

	downloadTasks, err := dm.Scan(context.Background())
	if err != nil {
		t.Fatalf("DownloadManager.Scan error %v", err)
	}
	for n, _ := range downloadTasks {
		if !reflect.DeepEqual(downloadTasks[n], result[n]) {
			t.Fatalf("downloadTask = %v, result = %v not equal", downloadTasks[n], result[n])
		}
	}

	for _, downloadTask := range downloadTasks {
		fmt.Printf("downloadTasks=%v\n", downloadTask)
	}

}

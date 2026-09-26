package store

import (
	"context"
	"path/filepath"
	"regexp"
	"testing"

	"github.com/staroffish/am/internal/config"
	"github.com/staroffish/am/internal/model"
)

func newTestDB(t *testing.T) *DB {
	t.Helper()
	db, err := NewDB(config.SQLiteConfig{Path: filepath.Join(t.TempDir(), "test.db")})
	if err != nil {
		t.Fatalf("NewDB: %v", err)
	}
	t.Cleanup(func() { _ = db.Close() })
	return db
}

func TestAnimeCRUDAndList(t *testing.T) {
	ctx := context.Background()
	db := newTestDB(t)

	ani := &model.Anime{
		ID:          "aaaaaaaaaaaaaaaaaaaaaaaa",
		AnimeNameJp: "テストアニメ",
		Cast:        "声优A",
		SerialsDuri: "2026/07~",
		Type:        "TV",
		StorDir:     "/store/TV/202607/test",
		Status:      "连载中",
		ImageBin:    []byte("fake-image-bytes"),
	}
	if err := db.SaveAnime(ctx, ani); err != nil {
		t.Fatalf("SaveAnime: %v", err)
	}

	got, err := db.GetAnime(ctx, ani.ID)
	if err != nil {
		t.Fatalf("GetAnime: %v", err)
	}
	if len(got.ImageBin) == 0 {
		t.Fatalf("GetAnime 应该带上 image")
	}
	if got.CreatedAt.IsZero() {
		t.Fatalf("CreatedAt 应被自动填充")
	}

	// 列表要能拿到所有字段，但不应该查询 image
	list, total, err := db.ListAnime(ctx, "", 0, 10, false)
	if err != nil || total != 1 || len(list) != 1 {
		t.Fatalf("ListAnime len=%d total=%d err=%v", len(list), total, err)
	}
	if len(list[0].ImageBin) != 0 {
		t.Fatalf("ListAnime 不应该查询 image 字段")
	}
	if list[0].AnimeNameJp != "テストアニメ" || list[0].StorDir == "" || list[0].Cast != "声优A" {
		t.Fatalf("ListAnime 字段丢失: %+v", list[0])
	}

	// "cast" 是 SQL 关键字，验证关键字搜索仍然可用
	if _, total, err = db.ListAnime(ctx, "声优A", 0, 10, false); err != nil || total != 1 {
		t.Fatalf("按 cast 搜索失败: total=%d err=%v", total, err)
	}
	if _, total, _ = db.ListAnime(ctx, "不存在", 0, 10, false); total != 0 {
		t.Fatalf("无关关键词应该搜不到, total=%d", total)
	}

	// upsert 不能重置 created_at，也不能产生第二行
	created := got.CreatedAt
	ani.Status = "已完结"
	if err := db.SaveAnime(ctx, ani); err != nil {
		t.Fatalf("SaveAnime upsert: %v", err)
	}
	got2, err := db.GetAnime(ctx, ani.ID)
	if err != nil {
		t.Fatalf("GetAnime after upsert: %v", err)
	}
	if !got2.CreatedAt.Equal(created) {
		t.Fatalf("upsert 后 created_at 变了: %v -> %v", created, got2.CreatedAt)
	}
	var count int64
	db.gorm.Model(&model.Anime{}).Count(&count)
	if count != 1 {
		t.Fatalf("upsert 应该只有一行, got %d", count)
	}

	if err := db.DeleteAnime(ctx, ani.ID); err != nil {
		t.Fatalf("DeleteAnime: %v", err)
	}
	if _, err := db.GetAnime(ctx, ani.ID); err == nil {
		t.Fatalf("删除后不应该还能查到")
	}
}

func TestMarkAnimeDone(t *testing.T) {
	ctx := context.Background()
	db := newTestDB(t)

	ani := &model.Anime{ID: "b", AnimeNameJp: "x", StorDir: "/x", Status: "连载中", SerialsDuri: "2026/07~"}
	if err := db.SaveAnime(ctx, ani); err != nil {
		t.Fatalf("SaveAnime: %v", err)
	}
	if err := db.MarkAnimeDone(ctx, "b"); err != nil {
		t.Fatalf("MarkAnimeDone: %v", err)
	}

	got, err := db.GetAnime(ctx, "b")
	if err != nil {
		t.Fatalf("GetAnime: %v", err)
	}
	if got.Status != "已完结" {
		t.Fatalf("status = %q", got.Status)
	}
	re := regexp.MustCompile(`^2026/07~\d{4}/\d{2}$`)
	if !re.MatchString(got.SerialsDuri) {
		t.Fatalf("serialsduri = %q, 期望 2026/07~YYYY/MM", got.SerialsDuri)
	}
}

func TestTaskAutoIncrementAndAnimeIDIndex(t *testing.T) {
	ctx := context.Background()
	db := newTestDB(t)

	t1 := &model.DownloadTask{Regexp: "a", AnimeID: "a1"}
	t2 := &model.DownloadTask{Regexp: "b", AnimeID: "a2"}
	if err := db.CreateTask(ctx, t1); err != nil {
		t.Fatalf("CreateTask 1: %v", err)
	}
	if err := db.CreateTask(ctx, t2); err != nil {
		t.Fatalf("CreateTask 2: %v", err)
	}
	if t1.ID == 0 || t2.ID != t1.ID+1 {
		t.Fatalf("自增 id 异常: %d, %d", t1.ID, t2.ID)
	}

	// 空 anime_id 不参与唯一约束，可以有多条
	if err := db.CreateTask(ctx, &model.DownloadTask{Regexp: "c"}); err != nil {
		t.Fatalf("空 anime_id 第一条: %v", err)
	}
	if err := db.CreateTask(ctx, &model.DownloadTask{Regexp: "d"}); err != nil {
		t.Fatalf("空 anime_id 第二条应该允许: %v", err)
	}

	// 重复的 anime_id 必须被拒绝
	if err := db.CreateTask(ctx, &model.DownloadTask{Regexp: "e", AnimeID: "a1"}); err == nil {
		t.Fatalf("重复 anime_id 应该被唯一索引拒绝")
	}

	before, err := db.GetTask(ctx, t1.ID)
	if err != nil {
		t.Fatalf("GetTask: %v", err)
	}
	upd := &model.DownloadTask{ID: t1.ID, Regexp: "a-updated", LatestChapter: 3, AnimeID: "a1"}
	if err := db.UpdateTask(ctx, upd); err != nil {
		t.Fatalf("UpdateTask: %v", err)
	}
	after, err := db.GetTask(ctx, t1.ID)
	if err != nil {
		t.Fatalf("GetTask after update: %v", err)
	}
	if after.Regexp != "a-updated" || after.LatestChapter != 3 {
		t.Fatalf("更新未生效: %+v", after)
	}
	if !after.CreatedAt.Equal(before.CreatedAt) {
		t.Fatalf("UpdateTask 覆盖了 created_at")
	}

	if err := db.UpdateTaskChapter(ctx, t1.ID, 9); err != nil {
		t.Fatalf("UpdateTaskChapter: %v", err)
	}
	after, _ = db.GetTask(ctx, t1.ID)
	if after.LatestChapter != 9 {
		t.Fatalf("latest_chapter = %d", after.LatestChapter)
	}

	if err := db.DeleteTask(ctx, t1.ID); err != nil {
		t.Fatalf("DeleteTask: %v", err)
	}
	if _, err := db.GetTask(ctx, t1.ID); err == nil {
		t.Fatalf("删除后不应该还能查到")
	}
}

func TestMagnets(t *testing.T) {
	ctx := context.Background()
	db := newTestDB(t)

	added, err := db.SaveMagnets(ctx, "2026-09-25", []*model.AnimeMagnet{
		{Name: "n1", MagnetLink: "m1"},
		{Name: "n2", MagnetLink: "m2"},
	})
	if err != nil || added != 2 {
		t.Fatalf("SaveMagnets added=%d err=%v", added, err)
	}

	// 重复的 (date,name) 忽略，新名字照常插入
	added, err = db.SaveMagnets(ctx, "2026-09-25", []*model.AnimeMagnet{
		{Name: "n1", MagnetLink: "changed"},
		{Name: "n3", MagnetLink: "m3"},
	})
	if err != nil || added != 1 {
		t.Fatalf("去重失败 added=%d err=%v", added, err)
	}

	rows, err := db.ListMagnetsByDate(ctx, "2026-09-25")
	if err != nil || len(rows) != 3 {
		t.Fatalf("ListMagnetsByDate len=%d err=%v", len(rows), err)
	}
	for _, r := range rows {
		if r.Name == "n1" && r.MagnetLink != "m1" {
			t.Fatalf("已存在的磁力被覆盖了: %+v", r)
		}
	}

	dates, err := db.ListMagnetDates(ctx)
	if err != nil || len(dates) != 1 || dates[0] != "2026-09-25" {
		t.Fatalf("ListMagnetDates = %v err=%v", dates, err)
	}

	if _, err := db.SaveMagnets(ctx, "2026-08-01", []*model.AnimeMagnet{{Name: "old", MagnetLink: "m0"}}); err != nil {
		t.Fatalf("SaveMagnets old: %v", err)
	}
	since, err := db.ListMagnetsSince(ctx, "2026-09-01")
	if err != nil || len(since) != 3 {
		t.Fatalf("ListMagnetsSince len=%d err=%v", len(since), err)
	}
	n, err := db.DeleteMagnetsBefore(ctx, "2026-09-01")
	if err != nil || n != 1 {
		t.Fatalf("DeleteMagnetsBefore n=%d err=%v", n, err)
	}
	if dates, _ = db.ListMagnetDates(ctx); len(dates) != 1 {
		t.Fatalf("清理后 dates = %v", dates)
	}
}

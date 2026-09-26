// Command migrate 把旧服务（仍在使用 MongoDB + Redis 的那套）里的数据搬到新的 SQLite。
//
// 用法（旧服务还在运行时执行，然后换新二进制启动）：
//
//	bin/migrate -old http://192.168.3.5:8222 -db /path/to/local/am.db
//
// 数据来源全部走旧服务的 HTTP API，因此本工具不依赖 mongo/redis 驱动。
package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
	"time"

	"github.com/staroffish/am/internal/config"
	"github.com/staroffish/am/internal/model"
	"github.com/staroffish/am/internal/store"
)

var (
	oldURL      = flag.String("old", "http://127.0.0.1:8222", "旧服务地址（仍是 Mongo/Redis 版本）")
	dbPath      = flag.String("db", "data/am.db", "目标 SQLite 文件路径")
	pageSize    = flag.Int("page-size", 200, "番剧列表分页大小")
	skipImages  = flag.Bool("skip-images", false, "不迁移封面图")
	skipMagnets = flag.Bool("skip-magnets", false, "不迁移磁力缓存")
	dryRun      = flag.Bool("dry-run", false, "只拉取统计，不写入 SQLite")
)

type legacyAnime struct {
	ID          string    `json:"id"`
	AnimeNameCn string    `json:"animenamecn"`
	AnimeNameJp string    `json:"animenamejp"`
	Cast        string    `json:"cast"`
	SerialsDuri string    `json:"serialsduri"`
	Type        string    `json:"type"`
	StorDir     string    `json:"stordir"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type legacyTask struct {
	ID            int32     `json:"id"`
	Regexp        string    `json:"regexp"`
	LatestChapter int32     `json:"latest_chapter"`
	AnimeID       string    `json:"anime_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type legacyMagnet struct {
	Name       string `json:"name"`
	MagnetLink string `json:"magnet_link"`
}

func main() {
	flag.Parse()
	log.SetFlags(log.LstdFlags)

	ctx := context.Background()
	client := &http.Client{Timeout: 120 * time.Second}

	var db *store.DB
	if !*dryRun {
		var err error
		db, err = store.NewDB(config.SQLiteConfig{Path: *dbPath})
		if err != nil {
			log.Fatalf("open sqlite: %v", err)
		}
		defer db.Close()
		log.Printf("target sqlite: %s", *dbPath)
	} else {
		log.Printf("DRY RUN: 不会写入任何数据")
	}

	// ---- 番剧 ----
	animes, err := fetchAnimes(ctx, client)
	if err != nil {
		log.Fatalf("fetch anime: %v", err)
	}
	log.Printf("fetched %d anime", len(animes))

	if !*dryRun {
		imported := 0
		withImage := 0
		for i := range animes {
			ani := toAnime(&animes[i])
			if !*skipImages && animes[i].ID != "" {
				if img, err := fetchImage(ctx, client, animes[i].ID); err != nil {
					log.Printf("image for %s (%s): %v", animes[i].ID, animes[i].AnimeNameJp, err)
				} else if len(img) > 0 {
					ani.ImageBin = img
					withImage++
				}
			}
			if err := db.ImportAnime(ctx, ani); err != nil {
				log.Fatalf("import anime %s: %v", animes[i].ID, err)
			}
			imported++
		}
		log.Printf("imported %d anime (%d with image)", imported, withImage)
	}

	// ---- 下载规则 ----
	tasks, err := fetchTasks(ctx, client)
	if err != nil {
		log.Fatalf("fetch tasks: %v", err)
	}
	log.Printf("fetched %d download tasks", len(tasks))
	if !*dryRun {
		for i := range tasks {
			t := &model.DownloadTask{
				ID:            tasks[i].ID,
				Regexp:        tasks[i].Regexp,
				LatestChapter: tasks[i].LatestChapter,
				AnimeID:       tasks[i].AnimeID,
				CreatedAt:     tasks[i].CreatedAt,
				UpdatedAt:     tasks[i].UpdatedAt,
			}
			if err := db.CreateTask(ctx, t); err != nil {
				log.Fatalf("import task %d: %v", tasks[i].ID, err)
			}
		}
		log.Printf("imported %d download tasks", len(tasks))
	}

	// ---- 磁力缓存 ----
	if *skipMagnets {
		log.Printf("skipped magnets")
		return
	}
	dates, err := fetchMagnetDates(ctx, client)
	if err != nil {
		log.Fatalf("fetch magnet dates: %v", err)
	}
	log.Printf("fetched %d magnet dates", len(dates))

	totalMagnets := 0
	for _, date := range dates {
		magnets, err := fetchMagnets(ctx, client, date)
		if err != nil {
			log.Printf("magnets %s: %v", date, err)
			continue
		}
		totalMagnets += len(magnets)
		if *dryRun {
			continue
		}
		rows := make([]*model.AnimeMagnet, 0, len(magnets))
		for _, m := range magnets {
			rows = append(rows, &model.AnimeMagnet{Name: m.Name, MagnetLink: m.MagnetLink})
		}
		if _, err := db.SaveMagnets(ctx, date, rows); err != nil {
			log.Fatalf("import magnets %s: %v", date, err)
		}
	}
	log.Printf("imported %d magnets", totalMagnets)
	log.Printf("done")
}

func toAnime(a *legacyAnime) *model.Anime {
	return &model.Anime{
		ID:          a.ID,
		AnimeNameCn: a.AnimeNameCn,
		AnimeNameJp: a.AnimeNameJp,
		Cast:        a.Cast,
		SerialsDuri: a.SerialsDuri,
		Type:        a.Type,
		StorDir:     a.StorDir,
		Status:      a.Status,
		CreatedAt:   a.CreatedAt,
		UpdatedAt:   a.UpdatedAt,
	}
}

// fetchAnimes 正序 + 倒序各翻一遍并按 id 去重。
// 旧接口只能按 updatetime 排序，同一时间戳的记录在分页时顺序不稳定，两遍可以互相补漏。
func fetchAnimes(ctx context.Context, client *http.Client) ([]legacyAnime, error) {
	seen := make(map[string]legacyAnime)
	for _, order := range []string{"asc", "desc"} {
		skip := 0
		for {
			page, total, err := fetchAnimePage(ctx, client, order, skip)
			if err != nil {
				return nil, err
			}
			for _, a := range page {
				if a.ID != "" {
					seen[a.ID] = a
				}
			}
			skip += len(page)
			if len(page) == 0 || skip >= total || len(page) < *pageSize {
				break
			}
		}
	}

	list := make([]legacyAnime, 0, len(seen))
	for _, a := range seen {
		list = append(list, a)
	}
	sort.Slice(list, func(i, j int) bool { return list[i].ID < list[j].ID })
	return list, nil
}

func fetchAnimePage(ctx context.Context, client *http.Client, order string, skip int) ([]legacyAnime, int, error) {
	var resp struct {
		Items []legacyAnime `json:"items"`
		Total int           `json:"total"`
	}
	u := fmt.Sprintf("%s/api/v1/anime?keyword=&sort=%s&skip=%d&limit=%d", *oldURL, order, skip, *pageSize)
	if err := getJSON(ctx, client, u, &resp); err != nil {
		return nil, 0, err
	}
	return resp.Items, resp.Total, nil
}

func fetchImage(ctx context.Context, client *http.Client, id string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/api/v1/anime/%s/image", *oldURL, id), nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		return nil, nil
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status %d", resp.StatusCode)
	}
	return io.ReadAll(resp.Body)
}

func fetchTasks(ctx context.Context, client *http.Client) ([]legacyTask, error) {
	var tasks []legacyTask
	if err := getJSON(ctx, client, *oldURL+"/api/v1/tasks", &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func fetchMagnetDates(ctx context.Context, client *http.Client) ([]string, error) {
	dates := []string{}
	if err := getJSON(ctx, client, *oldURL+"/api/v1/magnets/dates", &dates); err != nil {
		return nil, err
	}
	return dates, nil
}

func fetchMagnets(ctx context.Context, client *http.Client, date string) ([]legacyMagnet, error) {
	magnets := []legacyMagnet{}
	if err := getJSON(ctx, client, *oldURL+"/api/v1/magnets?date="+date, &magnets); err != nil {
		return nil, err
	}
	return magnets, nil
}

func getJSON(ctx context.Context, client *http.Client, u string, out interface{}) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("GET %s: status %d", u, resp.StatusCode)
	}
	return json.NewDecoder(resp.Body).Decode(out)
}

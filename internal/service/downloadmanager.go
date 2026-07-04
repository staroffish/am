package service

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"time"

	"github.com/staroffish/am/internal/model"
	"github.com/staroffish/am/internal/store"
)

type DownloadManagerService struct {
	mongo         *store.MongoClient
	redis         *store.RedisClient
	dlSvc         *DownloaderService
	animeSvc      *AnimeService
	log           *log.Logger
	magnetTimeout int
}

func NewDownloadManagerService(mongo *store.MongoClient, redis *store.RedisClient, dlSvc *DownloaderService, animeSvc *AnimeService, logger *log.Logger, magnetTimeout int) *DownloadManagerService {
	return &DownloadManagerService{
		mongo:         mongo,
		redis:         redis,
		dlSvc:         dlSvc,
		animeSvc:      animeSvc,
		log:           logger,
		magnetTimeout: magnetTimeout,
	}
}

func (s *DownloadManagerService) Scan(ctx context.Context) ([]model.MatchedTask, error) {
	animeMagnets, err := s.GetAnimeMagnets(ctx)
	if err != nil {
		return nil, fmt.Errorf("get anime magnets: %w", err)
	}

	tasks, err := s.mongo.ListTasks(ctx)
	if err != nil {
		return nil, fmt.Errorf("list tasks: %w", err)
	}

	var matchedTasks []model.MatchedTask

	for _, task := range tasks {
		latestChapter := task.LatestChapter + 1
		regexpStr := fmt.Sprintf(task.Regexp, latestChapter)

		reg, err := regexp.Compile(regexpStr)
		if err != nil {
			s.log.Printf("compile regexp for task %d: %v", task.ID, err)
			continue
		}

		storePath := ""
		if s.animeSvc != nil {
			if a, err := s.animeSvc.Get(ctx, task.AnimeID); err == nil {
				storePath = a.StorDir
			}
		}

		for _, magnet := range animeMagnets {
			found := reg.FindStringSubmatch(magnet.Name)
			if found == nil || len(found) < 2 {
				continue
			}

			matched := model.MatchedTask{
				TaskID:       task.ID,
				TaskName:     magnet.Name,
				MagnetLink:   magnet.MagnetLink,
				ChapterStart: latestChapter,
				ChapterEnd:   latestChapter,
				StorePath:    storePath,
				AnimeID:      task.AnimeID,
			}

			if len(found) > 2 {
				if chapEnd, err := strconv.ParseInt(found[2], 10, 64); err == nil {
					matched.ChapterEnd = int32(chapEnd)
				}
			}

			latestChapter = matched.ChapterEnd + 1
			s.log.Printf("matched: %s, chapter %d-%d", matched.TaskName, matched.ChapterStart, matched.ChapterEnd)
			matchedTasks = append(matchedTasks, matched)

			regexpStr = fmt.Sprintf(task.Regexp, latestChapter)
			reg, err = regexp.Compile(regexpStr)
			if err != nil {
				s.log.Printf("recompile regexp %s: %v", regexpStr, err)
				break
			}
		}
	}

	return matchedTasks, nil
}

func (s *DownloadManagerService) ScanAndDownload(ctx context.Context) ([]model.MatchedTask, error) {
	matchedTasks, err := s.Scan(ctx)
	if err != nil {
		return nil, err
	}

	var created []model.MatchedTask
	var addErrors []string
	for _, task := range matchedTasks {
		s.log.Printf("adding download: %s → %s, magnet: %s", task.TaskName, task.StorePath, task.MagnetLink)
		if err := s.dlSvc.Add(ctx, task.MagnetLink, task.StorePath); err != nil {
			msg := fmt.Sprintf("add %s ch%d failed: %v", task.TaskName, task.ChapterStart, err)
			s.log.Printf(msg)
			addErrors = append(addErrors, msg)
			continue
		}
		if err := s.mongo.UpdateTaskChapter(ctx, task.TaskID, task.ChapterEnd); err != nil {
			s.log.Printf("update task chapter error: %v", err)
		}
		created = append(created, task)
	}

	if len(addErrors) > 0 && len(created) == 0 {
		return nil, fmt.Errorf("all downloads failed: %s", addErrors[0])
	}

	return created, nil
}

func (s *DownloadManagerService) GetAnimeMagnets(ctx context.Context) ([]model.AnimeMagnet, error) {
	var result []model.AnimeMagnet
	now := time.Now()
	maxDays := s.magnetTimeout
	if maxDays <= 0 {
		maxDays = 30
	}

	for i := 0; i <= maxDays+7; i++ {
		date := now.AddDate(0, 0, -i).Format("2006-01-02")
		key := fmt.Sprintf("anime:link:%s", date)

		vals, err := s.redis.HGetAll(ctx, key)
		if err != nil {
			continue
		}

		if i > maxDays && len(vals) > 0 {
			s.redis.Client().Del(ctx, key)
			s.log.Printf("deleted expired magnet key: %s", key)
			continue
		}

		if len(vals) > 0 {
			s.log.Printf("Redis key %s: %d magnets", key, len(vals))
		}
		for name, magnetLink := range vals {
			result = append(result, model.AnimeMagnet{Name: name, MagnetLink: magnetLink})
		}
	}
	return result, nil
}

func (s *DownloadManagerService) AddTask(ctx context.Context, task *model.DownloadTask) error {
	return s.mongo.CreateTask(ctx, task)
}

func (s *DownloadManagerService) UpdateTask(ctx context.Context, task *model.DownloadTask) error {
	return s.mongo.UpdateTask(ctx, task)
}

func (s *DownloadManagerService) DeleteTask(ctx context.Context, id int32) error {
	return s.mongo.DeleteTask(ctx, id)
}

func (s *DownloadManagerService) ListTasks(ctx context.Context) ([]model.DownloadTask, error) {
	return s.mongo.ListTasks(ctx)
}

func (s *DownloadManagerService) GetTask(ctx context.Context, id int32) (*model.DownloadTask, error) {
	return s.mongo.GetTask(ctx, id)
}

func (s *DownloadManagerService) GetTaskByAnimeID(ctx context.Context, animeID string) (*model.DownloadTask, error) {
	return s.mongo.GetTaskByAnimeID(ctx, animeID)
}

package service

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/staroffish/am/internal/model"
	"github.com/staroffish/am/internal/store"
)

type AnimeService struct {
	db            *store.DB
	log           *log.Logger
	mainPageCount int
}

func NewAnimeService(db *store.DB, logger *log.Logger, mainPageCount int) *AnimeService {
	return &AnimeService{db: db, log: logger, mainPageCount: mainPageCount}
}

func (s *AnimeService) Get(ctx context.Context, id string) (*model.Anime, error) {
	return s.db.GetAnime(ctx, id)
}

func (s *AnimeService) List(ctx context.Context, keyword string, skip, limit int, asc bool) ([]model.Anime, int64, error) {
	return s.db.ListAnime(ctx, keyword, skip, limit, asc)
}

func (s *AnimeService) Save(ctx context.Context, ani *model.Anime) error {
	return s.db.SaveAnime(ctx, ani)
}

func (s *AnimeService) Delete(ctx context.Context, id string) error {
	return s.db.DeleteAnime(ctx, id)
}

func (s *AnimeService) MarkDone(ctx context.Context, id string) error {
	return s.db.MarkAnimeDone(ctx, id)
}

func (s *AnimeService) ListFiles(storDir string) ([]model.AnimeFile, error) {
	files, err := os.ReadDir(storDir)
	if err != nil {
		return nil, fmt.Errorf("read dir %s: %w", storDir, err)
	}

	var result []model.AnimeFile
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		result = append(result, model.AnimeFile{
			Name:     f.Name(),
			FullPath: storDir + "/" + f.Name(),
		})
	}
	return result, nil
}

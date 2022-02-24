package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/staroffish/am/app/spider/internal/biz/spider"
	dto "github.com/staroffish/am/common/dto/spider"
)

type AnimeSpiderRepo interface {
	GetWebContent(ctx context.Context) (string, error)
	SaveAnimeMagnets(ctx context.Context, animeMagets []*dto.AnimeMagnet) error
}

type AmSpider struct {
	repo     AnimeSpiderRepo
	amSpider spider.SpiderInterface
	log      *log.Helper
}

func NewAmSpider(repo AnimeSpiderRepo, spider spider.SpiderInterface, logger log.Logger) *AmSpider {
	return &AmSpider{
		repo:     repo,
		amSpider: spider,
		log:      log.NewHelper(logger),
	}
}

func (a *AmSpider) CrawlLink(ctx context.Context) ([]*dto.AnimeMagnet, error) {
	a.log.WithContext(ctx).Infof("Call AmSpider.CrawlLinkOnce")
	webContent, err := a.repo.GetWebContent(ctx)
	if err != nil {
		a.log.WithContext(ctx).Errorf("WebConentRepo.GetWebContent error: %v", err)
		return nil, err
	}

	a.log.WithContext(ctx).Infof("AmSpider.repo.GetWebContent length: %d", len(webContent))

	animeMagnets, err := a.amSpider.ExtractData(ctx, webContent)
	if err != nil {
		a.log.WithContext(ctx).Errorf("amSpider.ExtractData error: %v", err)
		return nil, err
	}

	return animeMagnets, a.repo.SaveAnimeMagnets(ctx, animeMagnets)
}

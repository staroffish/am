package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/staroffish/am/app/am-spider/internal/biz/amspider"
)

type AnimeSpiderRepo interface {
	GetWebContent(ctx context.Context) (string, error)
	SaveAnimeMagnets(ctx context.Context, animeMagets []*amspider.AnimeMagnet) error
}

type AmSpider struct {
	repo     AnimeSpiderRepo
	amSpider amspider.SpiderInterface
	log      *log.Helper
}

func NewAmSpider(repo AnimeSpiderRepo, spider amspider.SpiderInterface, logger log.Logger) *AmSpider {
	return &AmSpider{
		repo:     repo,
		amSpider: spider,
		log:      log.NewHelper(logger),
	}
}

func (a *AmSpider) CrawlLink(ctx context.Context) ([]*amspider.AnimeMagnet, error) {
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

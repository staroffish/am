package spider

import (
	"context"
	"log"

	"github.com/staroffish/am/internal/model"
)

type Spider interface {
	ExtractData(ctx context.Context, webContent string) ([]*model.AnimeMagnet, error)
}

const (
	DMHY     = "dmhy"
	NYAA     = "nyaa"
	MIOBT    = "miobt"
	BANGUMI  = "bangumi"
	MIKANANI = "mikanani"
)

func New(spiderType string, logger *log.Logger) Spider {
	switch spiderType {
	case MIOBT:
		return &MiobtSpider{}
	case DMHY:
		return &DmhySpider{}
	case NYAA:
		return &NyaaSpider{}
	case BANGUMI:
		return &BangumiSpider{}
	case MIKANANI:
		return &MikananiSpider{}
	}
	return nil
}

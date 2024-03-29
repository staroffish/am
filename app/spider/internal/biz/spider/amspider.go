package spider

import (
	"context"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/staroffish/am/app/spider/internal/conf"
	"github.com/staroffish/am/common/dto/spider"
)

type BaseSpider struct {
	log *log.Helper
}

type SpiderInterface interface {
	ExtractData(ctx context.Context, webContent string) ([]*spider.AnimeMagnet, error)
}

const (
	MIOBT   = "miobt"
	DMHY    = "dmhy"
	NYAA    = "nyaa"
	BANGUMI = "bangumi"
)

func NewSpider(spiderConf *conf.SpiderConfig, logger log.Logger) SpiderInterface {
	config := spiderConf.GetConfig()
	baseSpider := BaseSpider{
		log: log.NewHelper(logger),
	}
	switch config.Type {
	case MIOBT:
		return &MiobtSpider{
			BaseSpider: baseSpider,
		}
	case DMHY:
		return &DmhySpider{
			BaseSpider: baseSpider,
		}
	case NYAA:
		return &NyaaSpider{
			BaseSpider: baseSpider,
		}
	case BANGUMI:
		return &BangumiSpider{
			BaseSpider: baseSpider,
		}
	}
	return nil
}

func TrimAnimeName(name string) string {
	return strings.Trim(name, " \n\u200b \t")
}

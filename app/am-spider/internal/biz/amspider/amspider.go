package amspider

import (
	"context"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/staroffish/am/app/am-spider/internal/conf"
)

type BaseSpider struct {
	log *log.Helper
}

type AnimeMagnet struct {
	Name       string
	MagnetLink string
}

type SpiderInterface interface {
	ExtractData(ctx context.Context, webContent string) ([]*AnimeMagnet, error)
}

const (
	MIOBT = "miobt"
	DMHY  = "dmhy"
	NYAA  = "nyaa"
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
	}
	return nil
}

func TrimAnimeName(name string) string {
	return strings.Trim(name, " \n\u200b \t")
}

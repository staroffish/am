package biz

import (
	"github.com/google/wire"
	amspider "github.com/staroffish/am/app/spider/internal/biz/spider"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewAmSpider, amspider.NewSpider)

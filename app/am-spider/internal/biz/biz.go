package biz

import (
	"github.com/google/wire"
	"github.com/staroffish/am/app/am-spider/internal/biz/amspider"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewAmSpider, amspider.NewSpider)

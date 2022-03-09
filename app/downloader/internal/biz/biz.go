package biz

import (
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var BizProviderSet = wire.NewSet(NewDownloader)

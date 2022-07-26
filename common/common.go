package common

import (
	"github.com/google/wire"
	"github.com/staroffish/am/common/util"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(util.NewEtcdWatcher, util.NewLogger)

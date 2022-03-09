package server

import (
	"github.com/google/wire"
)

// ProviderSet is server providers.
var ServerProviderSet = wire.NewSet(NewHTTPServer, NewGRPCServer)

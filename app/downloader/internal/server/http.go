package server

import (
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	v1 "github.com/staroffish/am/api/downloader/v1"
	"github.com/staroffish/am/app/downloader/internal/conf"
	"github.com/staroffish/am/app/downloader/internal/service"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.DownloaderServerConfig, downloaderService *service.DownloaderService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Grpc.Timeout > 0 {
		opts = append(opts, http.Timeout(time.Duration(c.Http.Timeout)*time.Second))
	}
	srv := http.NewServer(opts...)
	v1.RegisterDownloaderHTTPServer(srv, downloaderService)

	return srv
}

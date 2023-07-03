package server

import (
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	v1 "github.com/staroffish/am/api/downloadmanager/v1"
	"github.com/staroffish/am/app/downloadmanager/internal/conf"
	"github.com/staroffish/am/app/downloadmanager/internal/service"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.DownloadManagerServerConfig, taskManager *service.DownloadmanagerService, logger log.Logger) *http.Server {
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
	if c.Http.Timeout > 0 {
		opts = append(opts, http.Timeout(time.Duration(c.Http.Timeout)*time.Second))
	}
	srv := http.NewServer(opts...)
	log := log.NewHelper(logger)
	log.Infof("srr=%v", srv)
	v1.RegisterDownloadmanagerHTTPServer(srv, taskManager)

	return srv
}

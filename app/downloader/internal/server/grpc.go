package server

import (
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	v1 "github.com/staroffish/am/api/downloader/v1"
	"github.com/staroffish/am/app/downloader/internal/conf"
	"github.com/staroffish/am/app/downloader/internal/service"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.DownloaderServerConfig, downloaderService *service.DownloaderService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout > 0 {
		opts = append(opts, grpc.Timeout(time.Duration(c.Grpc.Timeout)*time.Second))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterDownloaderServer(srv, downloaderService)
	return srv
}

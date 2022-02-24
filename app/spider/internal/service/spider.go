package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/staroffish/am/api/spider/v1"
	"github.com/staroffish/am/app/spider/internal/biz"
)

type AmspiderService struct {
	log      *log.Helper
	amSpider *biz.AmSpider
	pb.UnimplementedSpiderServer
}

func NewAmspiderService(amSpider *biz.AmSpider, logger log.Logger) *AmspiderService {
	return &AmspiderService{
		amSpider: amSpider,
		log:      log.NewHelper(logger),
	}
}

func (s *AmspiderService) Crawl(ctx context.Context, _ *pb.CrawlRequest) (*pb.CrawlResponse, error) {
	s.log.WithContext(ctx).Info("Call AmspiderService.Crawl")
	animeMagnets, err := s.amSpider.CrawlLink(ctx)
	animeMagnetdatas := []*pb.CrawlResponse_AnimeMagnetData{}

	for _, animeMagnet := range animeMagnets {
		animeMagnetdatas = append(animeMagnetdatas, &pb.CrawlResponse_AnimeMagnetData{
			ItemName: animeMagnet.Name,
			Magnet:   animeMagnet.MagnetLink,
		})
	}
	return &pb.CrawlResponse{
		AnimeMagnetdatas: animeMagnetdatas,
	}, err
}

package spider

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/staroffish/am/common/dto/spider"
)

type BangumiSpider struct {
	BaseSpider
}

type bangumiItems struct {
	PageCount int              `json: "page_count"`
	Torrents  []bangumiTorrent `json: "torrents"`
}

type bangumiTorrent struct {
	Title  string `json: "title"`
	Magnet string `json: "magnet"`
}

func (m *BangumiSpider) ExtractData(ctx context.Context, webContent string) ([]*spider.AnimeMagnet, error) {
	m.log.WithContext(ctx).Info("Call BangumiSpider.ExtractData")

	animeMagnets := []*spider.AnimeMagnet{}

	items := &bangumiItems{
		Torrents: []bangumiTorrent{},
	}

	if err := json.Unmarshal([]byte(webContent), items); err != nil {
		return nil, fmt.Errorf("BangumiSpider.ExtractData error. err:%v, web content=%v", err, webContent)
	}

	for _, item := range items.Torrents {
		animeMagnets = append(animeMagnets, &spider.AnimeMagnet{
			Name:       item.Title,
			MagnetLink: item.Magnet,
		})
	}

	return animeMagnets, nil
}

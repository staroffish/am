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
	PageCount int              `json:"page_count"`
	Torrents  []bangumiTorrent `json:"torrents"`
}

type bangumiTorrent struct {
	Title  string `json:"title"`
	Magnet string `json:"magnet"`
}

const trackerServerStr = "&tr=https%3A%2F%2Ftr.bangumi.moe%3A9696%2Fannounce&tr=http%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce&tr=udp%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=http%3A%2F%2F208.67.16.113%3A8000%2Fannounce&tr=udp%3A%2F%2F208.67.16.113%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.ktxp.com%3A6868%2Fannounce&tr=http%3A%2F%2Ftracker.ktxp.com%3A7070%2Fannounce&tr=http%3A%2F%2Ft2.popgo.org%3A7456%2Fannonce&tr=http%3A%2F%2Fbt.sc-ol.com%3A2710%2Fannounce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2F61.154.116.205%3A8000%2Fannounce&tr=http%3A%2F%2Fbt.rghost.net%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.nyaatorrents.info%3A6544%2Fannounce"

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
			MagnetLink: item.Magnet + trackerServerStr,
		})
	}

	return animeMagnets, nil
}

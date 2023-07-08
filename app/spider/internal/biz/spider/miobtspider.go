package spider

import (
	"context"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/staroffish/am/common/dto/spider"
)

// 支持的BT网站列表
// www.miobt.com

type MiobtSpider struct {
	BaseSpider
}

func (m *MiobtSpider) ExtractData(ctx context.Context, webContent string) ([]*spider.AnimeMagnet, error) {
	m.log.WithContext(ctx).Info("Call MiobtSpider.ExtractData")

	animeMagnets := []*spider.AnimeMagnet{}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(webContent))
	if err != nil {
		m.log.WithContext(ctx).Errorf("MiobtSpider.ExtractData new document error:%v", err)
		return nil, err
	}

	doc.Find("tbody#data_list").Each(func(_ int, tbodySelector *goquery.Selection) {
		tbodySelector.Find("a[href][target=_blank]").Each(func(_ int, aSelector *goquery.Selection) {
			href, exsists := aSelector.Attr("href")
			if !exsists {
				return
			}
			href = strings.TrimPrefix(href, "show-")
			magnet := strings.TrimSuffix(href, ".html")
			animeMagnets = append(animeMagnets, &spider.AnimeMagnet{
				Name:       TrimAnimeName(aSelector.Text()),
				MagnetLink: fmt.Sprintf("magnet:?xt=urn:btih:%s&tr=http://open.acgtracker.com:1096/announce", magnet),
			})
		})
	})
	return animeMagnets, nil
}

package spider

import (
	"context"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/staroffish/am/common/dto/spider"
)

type MikananiSpider struct {
	BaseSpider
}

func (m *MikananiSpider) ExtractData(ctx context.Context, webContent string) ([]*spider.AnimeMagnet, error) {
	m.log.WithContext(ctx).Info("Call MikananiSpider.ExtractData")

	animeMagnets := []*spider.AnimeMagnet{}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(webContent))
	if err != nil {
		m.log.WithContext(ctx).Errorf("MikananiSpider.ExtractData new document error:%v", err)
		return nil, err
	}

	doc.Find("tbody").First().Find("tr").Each(func(_ int, trSelector *goquery.Selection) {
		animeMagnet := spider.AnimeMagnet{}
		
		// 获取番组名和磁力链接
		trSelector.Find("td").Eq(2).Find("a.magnet-link-wrap").First().EachWithBreak(func(_ int, aSelector *goquery.Selection) bool {
			animeMagnet.Name = TrimAnimeName(aSelector.Text())
			return false
		})

		trSelector.Find("td").Eq(2).Find("a.js-magnet").EachWithBreak(func(_ int, aSelector *goquery.Selection) bool {
			magnet, exists := aSelector.Attr("data-clipboard-text")
			if !exists {
				return false
			}
			animeMagnet.MagnetLink = magnet
			return false
		})

		if animeMagnet.MagnetLink == "" || animeMagnet.Name == "" {
			return
		}
		animeMagnets = append(animeMagnets, &animeMagnet)
	})
	
	return animeMagnets, nil
}
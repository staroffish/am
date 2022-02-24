package spider

import (
	"context"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type NyaaSpider struct {
	BaseSpider
}

func (m *NyaaSpider) ExtractData(ctx context.Context, webContent string) ([]*AnimeMagnet, error) {
	m.log.WithContext(ctx).Info("Call NyaaSpider.ExtractData")

	animeMagnets := []*AnimeMagnet{}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(webContent))
	if err != nil {
		m.log.WithContext(ctx).Errorf("MiobtSpider.ExtractData new document error:%v", err)
		return nil, err
	}

	doc.Find("tbody").Each(func(_ int, tbodySelector *goquery.Selection) {
		tbodySelector.Find("tr[class]").Each(func(_ int, trSelector *goquery.Selection) {
			animeMagnet := AnimeMagnet{}
			trSelector.Find("td>a[href*=view][class!=comments]").EachWithBreak(func(_ int, aSelector *goquery.Selection) bool {
				animeMagnet.Name = TrimAnimeName(aSelector.Text())
				return false
			})

			trSelector.Find("td>a[href*=magnet]").EachWithBreak(func(_ int, aSelector *goquery.Selection) bool {
				href, exsists := aSelector.Attr("href")
				if !exsists {
					return false
				}

				animeMagnet.MagnetLink = href
				return false
			})

			if animeMagnet.MagnetLink == "" || animeMagnet.Name == "" {
				return
			}
			animeMagnets = append(animeMagnets, &animeMagnet)
		})
	})
	return animeMagnets, nil
}

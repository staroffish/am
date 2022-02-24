package spider

import (
	"context"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/staroffish/am/common/dto/spider"
)

type DmhySpider struct {
	BaseSpider
}

func (m *DmhySpider) ExtractData(ctx context.Context, webContent string) ([]*spider.AnimeMagnet, error) {
	m.log.WithContext(ctx).Info("Call DmhySpider.ExtractData")

	animeMagnets := []*spider.AnimeMagnet{}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(webContent))
	if err != nil {
		m.log.WithContext(ctx).Errorf("MiobtSpider.ExtractData new document error:%v", err)
		return nil, err
	}

	doc.Find("tbody").Each(func(_ int, tbodySelector *goquery.Selection) {
		tbodySelector.Find("tr").Each(func(_ int, trSelector *goquery.Selection) {
			animeMagnet := spider.AnimeMagnet{}
			trSelector.Find("td.title>a[href][target=_blank]").EachWithBreak(func(_ int, aSelector *goquery.Selection) bool {
				animeMagnet.Name = TrimAnimeName(aSelector.Text())
				return false
			})

			trSelector.Find("td[nowrap=nowrap][align=center]>a[title=磁力下載]").EachWithBreak(func(_ int, aSelector *goquery.Selection) bool {
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

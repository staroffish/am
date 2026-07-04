package spider

import (
	"context"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/staroffish/am/internal/model"
)

type DmhySpider struct{}

func (s *DmhySpider) ExtractData(ctx context.Context, webContent string) ([]*model.AnimeMagnet, error) {
	animeMagnets := []*model.AnimeMagnet{}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(webContent))
	if err != nil {
		return nil, err
	}

	doc.Find("tbody").Each(func(_ int, tbodySelector *goquery.Selection) {
		tbodySelector.Find("tr").Each(func(_ int, trSelector *goquery.Selection) {
			animeMagnet := model.AnimeMagnet{}
			trSelector.Find("td.title>a[href][target=_blank]").EachWithBreak(func(_ int, aSelector *goquery.Selection) bool {
				animeMagnet.Name = trimName(aSelector.Text())
				return false
			})
			trSelector.Find("td[nowrap=nowrap][align=center]>a[title=磁力下載]").EachWithBreak(func(_ int, aSelector *goquery.Selection) bool {
				href, exists := aSelector.Attr("href")
				if !exists {
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

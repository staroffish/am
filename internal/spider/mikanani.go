package spider

import (
	"context"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/staroffish/am/internal/model"
)

type MikananiSpider struct{}

func (s *MikananiSpider) ExtractData(ctx context.Context, webContent string) ([]*model.AnimeMagnet, error) {
	animeMagnets := []*model.AnimeMagnet{}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(webContent))
	if err != nil {
		return nil, err
	}

	doc.Find("tbody").First().Find("tr").Each(func(_ int, trSelector *goquery.Selection) {
		animeMagnet := model.AnimeMagnet{}

		trSelector.Find("td").Eq(2).Find("a.magnet-link-wrap").First().EachWithBreak(func(_ int, aSelector *goquery.Selection) bool {
			animeMagnet.Name = trimName(aSelector.Text())
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

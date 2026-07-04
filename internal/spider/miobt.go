package spider

import (
	"context"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/staroffish/am/internal/model"
)

type MiobtSpider struct{}

func (s *MiobtSpider) ExtractData(ctx context.Context, webContent string) ([]*model.AnimeMagnet, error) {
	animeMagnets := []*model.AnimeMagnet{}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(webContent))
	if err != nil {
		return nil, err
	}

	doc.Find("tbody#data_list").Each(func(_ int, tbodySelector *goquery.Selection) {
		tbodySelector.Find("a[href][target=_blank]").Each(func(_ int, aSelector *goquery.Selection) {
			href, exists := aSelector.Attr("href")
			if !exists {
				return
			}
			href = strings.TrimPrefix(href, "show-")
			magnet := strings.TrimSuffix(href, ".html")
			animeMagnets = append(animeMagnets, &model.AnimeMagnet{
				Name:       trimName(aSelector.Text()),
				MagnetLink: fmt.Sprintf("magnet:?xt=urn:btih:%s&tr=http://open.acgtracker.com:1096/announce", magnet),
			})
		})
	})
	return animeMagnets, nil
}

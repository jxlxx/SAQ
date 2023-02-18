package saq

import (
	"fmt"
	"github.com/gocolly/colly"
	"net/url"
)

type Api struct {
	colly *colly.Collector
	lang  string
}

func New() *Api {

	return &Api{
		colly: colly.NewCollector(),
		lang:  "en",
	}
}

func (s *Api) createQueryEndpoint(q string) string {
	safeQuery := url.QueryEscape(q)

	return fmt.Sprintf("https://saq.com/%s/catalogsearch/result/?q=%s", s.lang, safeQuery)
}

var product_card_name = "div.product-item-content-container > div.product.details.product-item-details > div.product.content-wrapper > strong.product.name.product-item-name > a"

func (s *Api) Query(q string) {
	queryEndpoint := s.createQueryEndpoint(q)

	s.colly.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting", r.URL.String())
	})

	s.colly.OnHTML(`div.product-item-info`, func(e *colly.HTMLElement) {
		name := e.ChildText(product_card_name)

		fmt.Println("found product: ", name)

	})

	s.colly.Visit(queryEndpoint)

}

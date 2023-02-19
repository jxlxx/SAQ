package saq

import (
	"fmt"
	"github.com/gocolly/colly"
	"net/url"
	"strings"
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

var searchCardRoot = "div.product-item-content-container > div.product.details.product-item-details > div.product.content-wrapper"

var searchCardPaths = map[string]string{
	"name":                searchCardRoot + " > strong.product.name.product-item-name > a",
	"type_volume_country": searchCardRoot + " > strong.product.product-item-identity-format",
	"price":               searchCardRoot + " > div.price-box.price-final_price",
	"saq_code":            searchCardRoot + " > div.saq-code",
	"rating_summary":      searchCardRoot + " > div.product-reviews-summary.short > div.rating-summary",
	"rating_actions":      searchCardRoot + " > div.product-reviews-summary.short > div.rating-actions",
	"marketing":           searchCardRoot + " > div.wrapper-marketing-brand > img",
	// "_":                   searchCardRoot + " > ",
}

var nextPageResults = `#maincontent > div.columns > div.column.main > div.search.results > div.toolbar.toolbar-products.bottom > div.pages > ul > li.item.pages-item-next > a`

func trimSpace(s string) string {
	w := strings.Fields(s)
	return strings.Join(w, " ")
}

func (s *Api) Query(q string) {
	queryEndpoint := s.createQueryEndpoint(q)

	s.colly.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting", r.URL.String())
	})

	s.colly.OnHTML(`div.product-item-info`, func(e *colly.HTMLElement) {

		name := trimSpace(e.ChildText(searchCardPaths["name"]))
		tvc := strings.Split(trimSpace(e.ChildText(searchCardPaths["type_volume_country"])), " | ")
		price := trimSpace(e.ChildText(searchCardPaths["price"]))
		saq_code := trimSpace(e.ChildText(searchCardPaths["saq_code"]))
		rating_summary := trimSpace(e.ChildText(searchCardPaths["rating_summary"]))
		rating_actions := trimSpace(e.ChildText(searchCardPaths["rating_actions"]))

		fmt.Println(name, saq_code, tvc[0], tvc[1], tvc[2], price, rating_summary, rating_actions)

	})

	s.colly.OnHTML(nextPageResults, func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	s.colly.Visit(queryEndpoint)

}

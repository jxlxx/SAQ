package saq

import (
	"fmt"
	"github.com/gocolly/colly"
	"net/url"
	"strings"
)

type Api struct {
	colly   *colly.Collector
	lang    Language
	results *SearchResults
}

func New(lang Language) *Api {

	searchResults := SearchResults{
		list:  make(chan ProductInfo, 100),
		query: "",
	}

	return &Api{
		colly:   colly.NewCollector(),
		lang:    lang,
		results: &searchResults,
	}
}

func (s *Api) createQueryEndpoint(q string) string {
	safeQuery := url.QueryEscape(q)

	return fmt.Sprintf("https://saq.com/%s/catalogsearch/result/?q=%s", s.lang.String(), safeQuery)
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

	s.results = &SearchResults{
		query: q,
		list:  make(chan ProductInfo, 100),
	}

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

		newProduct := ProductInfo{
			name:              name,
			product_link:      s.createProductLink(saq_code),
			saq_code:          saq_code,
			catagory:          tvc[0],
			volume:            tvc[1],
			country_of_origin: tvc[2],
			price:             price,
			rating_summary:    rating_summary,
			rating_actions:    rating_actions,
			bottled_in_quebec: false, // TODO
		}

		s.results.list <- newProduct

	})

	s.colly.OnHTML(nextPageResults, func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	s.colly.Visit(queryEndpoint)

}

func (s *Api) createProductLink(id string) string {
	return fmt.Sprintf("https://saq.com/%s/%s", s.lang.String(), id)
}

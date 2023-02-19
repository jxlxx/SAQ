package saq

import (
	"fmt"
	"github.com/gocolly/colly"
	"net/url"
	"strings"
)

type Scraper struct {
	colly *colly.Collector
	lang  Language
	List  chan ProductInfo
}

func New(lang Language) *Scraper {

	return &Scraper{
		colly: colly.NewCollector(),
		lang:  lang,
		List:  make(chan ProductInfo, 100),
	}
}

func (s *Scraper) createQueryEndpoint(q string) string {
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

func parseSAQCode(s string) string {
	return strings.Split(s, " ")[2]
}

func (s *Scraper) Query(q string) {

	queryEndpoint := s.createQueryEndpoint(q)

	s.colly.OnError(func(r *colly.Response, e error) {
		fmt.Println("error encountered: ", e)
		return
	})

	s.colly.OnHTML(`div.product-item-info`, func(e *colly.HTMLElement) {

		name := trimSpace(e.ChildText(searchCardPaths["name"]))
		tvc := strings.Split(trimSpace(e.ChildText(searchCardPaths["type_volume_country"])), " | ")
		price := trimSpace(e.ChildText(searchCardPaths["price"]))
		saq_code := parseSAQCode(e.ChildText(searchCardPaths["saq_code"]))
		rating_summary := trimSpace(e.ChildText(searchCardPaths["rating_summary"]))
		rating_actions := trimSpace(e.ChildText(searchCardPaths["rating_actions"]))

		if len(name) == 0 || len(tvc) != 3 {
			return
		}

		newProduct := ProductInfo{
			Name:            name,
			ProductLink:     s.createProductLink(saq_code),
			SaqCode:         saq_code,
			Catagory:        tvc[0],
			Volume:          tvc[1],
			CountryOfOrigin: tvc[2],
			Price:           price,
			RatingSummary:   rating_summary,
			RatingActions:   rating_actions,
			BottledInQuebec: false, // TODO
		}

		s.List <- newProduct

	})

	s.colly.OnHTML(nextPageResults, func(e *colly.HTMLElement) {
		err := e.Request.Visit(e.Attr("href"))
		if err != nil {
			fmt.Println("error visiting next page: ", e.Attr("href"))
			return
		}
	})

	s.colly.Visit(queryEndpoint)

	defer close(s.List) // This means that the channel cannot be written to again

}

func (s *Scraper) createProductLink(id string) string {
	return fmt.Sprintf("https://saq.com/%s/%s", s.lang.String(), id)
}

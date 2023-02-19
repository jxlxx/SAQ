package main

import (
	"saq_scraper/saq"
)

func main() {

	api := saq.New()

	api.Query("a")

}

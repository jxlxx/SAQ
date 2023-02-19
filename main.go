package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"saq_scraper/saq"
)

func main() {

	query := "Chartreuse"
	language := saq.English

	file_name := "saq_products.csv"
	file, err := os.Create(file_name)

	if err != nil {
		fmt.Println("error cannot create file: ", file_name, err)
		return
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	scraper := saq.New(language)

	go scraper.Query(query)

	for {
		select {
		case product, ok := <-scraper.List:
			if !ok {
				return
			}
			fmt.Println("Found: ", product.Name)
			err = writer.Write(product.ToStringArray())
			if err != nil {
				fmt.Println("error writing to file: ", err)
				return
			}
		}
	}
}

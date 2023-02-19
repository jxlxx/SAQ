package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"
	"saq_scraper/saq"
)

func main() {

	query := "ugly"
	language := saq.English

	file_name := "saq_products.csv"
	file, err := os.Create(file_name)

	if err != nil {
		fmt.Println("Cannot create file: ", file_name, err)
		return
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	scraper := saq.New(language)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go scraper.Query(query, ctx, cancel)

	for {
		select {
		case product := <-scraper.List:
			fmt.Println("Found", product.Name)
			err = writer.Write(product.ToStringArray())
			if err != nil {
				fmt.Println("error writing to file: ", err)
				return
			}
		case <-ctx.Done():
			return
		}
	}
}

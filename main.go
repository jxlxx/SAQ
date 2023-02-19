package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"
	"saq_scraper/saq"
)

func main() {

	file_name := "saq_products.csv"
	file, err := os.Create(file_name)

	if err != nil {
		fmt.Println("Cannot create file: ", file_name, err)
		return
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	api := saq.New(saq.English)
	query := "a"

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go api.Query(query, ctx, cancel)

	for {
		select {
		case product := <-api.List:
			fmt.Println("product", product.Name)
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

package main

import (
	"fmt"
	"os"
	"saq_scraper/saq"
)

func main() {

	api := saq.New(saq.English)

	file_name := "saq_products.csv"

	file, err := os.Create(file_name)

	if err != nil {
		fmt.Println("Cannot create file: ", file_name, err)
		return
	}

	defer file.Close()

	api.Query("a")

}

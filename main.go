package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
)

func ExampleScrape() {
	fmt.Println("Runnig scrape")
	doc, err := goquery.NewDocument("http://metalsucks.net")
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".sidebar-reviews article .content-block").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		band := s.Find("a").Text()
		title := s.Find("i").Text()
		fmt.Printf("Review %d: %s - %s\n", i, band, title)
	})
}

func main() {
	ExampleScrape()
}
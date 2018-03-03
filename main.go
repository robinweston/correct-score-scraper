package main

import (
	"fmt"
	"log"
	"regexp"
	"github.com/PuerkitoBio/goquery"
)

func reSubMatchMap(r *regexp.Regexp, str string) (map[string]string) {
    match := r.FindStringSubmatch(str)
    subMatchMap := make(map[string]string)
    for i, name := range r.SubexpNames() {
        if i != 0 {
            subMatchMap[name] = match[i]
        }
    }
	return subMatchMap
}

func extractTeams(s *goquery.Selection) string {
	summary, _ := s.Attr("summary")
	r := regexp.MustCompile(`shows (?P<Teams>.+) - Correct Score`)
	return reSubMatchMap(r, summary)["Teams"]
}

func ScrapeScores() {
	url := "http://sports.williamhill.com/bet/en-gb/betting/g/344/Correct+Score.html"
	fmt.Println("Scraping", url)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("table[summary*='Correct Score']").Each(func(i int, s *goquery.Selection) {
		teams := extractTeams(s)
		fmt.Printf("%s\n", teams)
	})
}

func main() {
	ScrapeScores()
}

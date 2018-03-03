package main

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strings"
	"strconv"
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

type score struct {
	Result string
	Odds float64
}

type match struct {
	Teams string
	LikeliestScore score
}

func parseOdds(oddsText string) float64 {
	r := regexp.MustCompile(`(?P<Num>\d+)/(?P<Dem>\d+)`)
	matches := reSubMatchMap(r, oddsText)
	num, _ := strconv.ParseFloat(matches["Num"], 64)
	dem, _ := strconv.ParseFloat(matches["Dem"], 64)
	return num / dem
}

func extractScores(parentTable *goquery.Selection) []score {
	scores := []score{}
	parentTable.Find("div.eventprice").Each(func(i int, s *goquery.Selection) {
		
		result := strings.TrimSpace(s.Closest("td").Next().Text())
		oddsText := strings.TrimSpace(s.Text())
		scores = append(scores, score{Result: result , Odds: parseOdds(oddsText)})
	})
	return scores
}

func findLikeliestScore(scores []score) score {
	sort.Slice(scores, func(i, j int) bool {
		return scores[i].Odds < scores[j].Odds
	})
	return scores[0]
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
		scores := extractScores(s)
		likeliestScore := findLikeliestScore(scores)
		fmt.Printf("%s %v\n", teams, likeliestScore)
	})
}

func main() {
	ScrapeScores()
}

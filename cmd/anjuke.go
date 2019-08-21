package cmd

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/cnbattle/anjuke/city"
	"github.com/cnbattle/anjuke/config"
	"log"
	"net/http"
)

// GrabAll
func GrabAll() {
	// Request the HTML page.
	res, err := http.Get("https://www.anjuke.com/sy-city.html")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	// body > div.content > div > div.letter_city > ul
	doc.Find("div.city_list > a").Each(func(i int, s *goquery.Selection) {
		name := s.Text()
		url, _ := s.Attr("href")
		city.Grab(name, url)
	})
}

// Grab
func Grab() {
	for _, cityInfo := range config.V.Cites {
		city.Grab(cityInfo.Name, cityInfo.Url)
	}
}

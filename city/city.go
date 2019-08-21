package city

import (
	"fmt"
	"github.com/cnbattle/anjuke/config"
	"github.com/cnbattle/anjuke/database"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Grab 抓取某一城市数据
func Grab(cityName, url string) {
	// Request the HTML page
	log.Println("城市:", cityName)
	log.Println("地址:", url)
	res, err := http.Get(url + "/community/?from=navigation")
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

	// 处理区县的链接
	communityCount := grabCommunitiesUrls(cityName, doc)
	log.Println("区县数量:", communityCount)
}

// grabCommunities 获取区县链接
func grabCommunitiesUrls(cityName string, doc *goquery.Document) (communityCount int) {
	doc.Find("body > div.w1180 > div.div-border.items-list > div:nth-child(1) > span.elems-l > a").Each(func(i int, s *goquery.Selection) {
		if i != 0 {
			name := s.Text()
			url, _ := s.Attr("href")
			log.Println("区县:", name)
			log.Println("地址:", url)
			handleCommunity(cityName, name, url)
			communityCount++
		}
	})
	return
}

// handleCommunity 开始处理一个区县的列表页面
func handleCommunity(cityName string, community, url string) (communityCount int) {
START:
	time.Sleep(time.Duration(config.V.Sleep) * time.Second)
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
		time.Sleep(30 * time.Second)
		goto START
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Printf("handle Community status code error: %d %s \n", res.StatusCode, url)
		time.Sleep(30 * time.Second)
		goto START
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Println(err)
		time.Sleep(30 * time.Second)
		goto START
	}
	grabItem(cityName, community, doc)

	if nextPage, exists := doc.Find("a.aNxt").Attr("href"); exists {
		handleCommunity(cityName, community, nextPage)
	}
	return
}

// grabItem 抓取处理小区信息
func grabItem(cityName string, community string, doc *goquery.Document) {
	doc.Find("#list-content > div.li-itemmod").Each(func(i int, s *goquery.Selection) {
		name := s.Find("h3 > a").Text()
		price := s.Find("p > strong").Text()
		url, _ := s.Find("h3 > a").Attr("href")
		address := s.Find("div.li-info  > address").Text()
		date := s.Find("div.li-info  > p.date").Text()
		priceTxt := s.Find("div.li-side > p.price-txt").Text()
		cover, _ := s.Find("a > img").Attr("src")
		var data database.Data
		data.Name = strings.Trim(name, " ")
		data.Cover = strings.Trim(cover, " ")
		data.CityName = strings.Trim(cityName, " ")
		data.Community = strings.Trim(community, " ")
		data.Address = strings.Trim(address, " ")
		data.Url = strings.Trim(url, " ")
		data.Price = strings.Trim(price, " ")
		data.PriceTxt = strings.Trim(priceTxt, " ")
		data.Date = strings.Trim(date, " ")
		database.Local.Create(&data)
		fmt.Print(".")
	})
	fmt.Println()
}

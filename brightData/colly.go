package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	ScrapeWithColly()
}

func ScrapeWithColly() {
	// Define the proxy server with username and password
	proxyUsername := "username"
	proxyPassword := "your_password"
	proxyHost := "Server_info"
	proxyPort := "port"

	proxyStr := fmt.Sprintf("http://%s:%s@%s:%s", url.QueryEscape(proxyUsername), url.QueryEscape(proxyPassword), proxyHost, proxyPort)

	// Parse the proxy URL
	proxyURL, err := url.Parse(proxyStr)
	if err != nil {
		fmt.Println("Error parsing proxy URL:", err)
		return
	}

	//Create an http.Transport that uses the proxy
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true, // Disable SSL certificate verification
		},
	}

	c := colly.NewCollector()

	c.WithTransport(transport)
	// Find and visit all links
	c.OnHTML("body", func(e *colly.HTMLElement) {
		doc := e.DOM
		doc.Find(".section-blog article").Each(func(i int, s *goquery.Selection) {
			article := map[string]string{}
			// For each item found, get the title
			article["title"] = s.Find("div .post-card__title a").Text()
			//fmt.Printf("Title: %s\n", title)

			// For each item found, get the excerpt
			article["excerpt"] = strings.Trim(s.Find("div .post-card__excerpt").Text(), "\n")

			// For each item found, get the category
			article["category"] = s.Find("div .post-card__tag").Text()

			jsonData, err := json.Marshal(article)
			if err != nil {
				log.Fatal(err)
			}
			// Print article details as Json Object
			fmt.Printf("Article %d: %v\n", i, string(jsonData))
		})
	})

	if err := c.Visit("https://itsfoss.com"); err != nil {
		log.Fatal(err)
	}
}

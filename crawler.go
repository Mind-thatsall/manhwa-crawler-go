package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
	"github.com/rs/xid"
)

type manhwa struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

func getInfos() []manhwa {
	var manhwas = []manhwa{}

	// Instantiate default collector
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/110.0"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("div[class='uta']", func(e *colly.HTMLElement) {
		id := xid.New()

		manhwas = append(manhwas, manhwa{ID: id.String(), Name: e.ChildText("h4"), Picture: e.ChildAttr("img", "src")})
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnError(func(r *colly.Response, err error) {

		log.Println("ERROR : ", r.Request.URL, err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited:", r.Request.URL)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished !")
	})

	c.Visit("https://elarcpage.com/")

	return manhwas
}

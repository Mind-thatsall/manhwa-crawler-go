package models

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly"
	"github.com/rs/xid"
)

func getManhwas() []Manhwa {
	var manhwas = []Manhwa{}

	// Instantiate default collector
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/110.0"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("div[class='uta']", func(e *colly.HTMLElement) {
		id := xid.New()

		slug := strings.Join(strings.Split(strings.ToLower(e.ChildText("h4")), " "), "-")

		manhwas = append(manhwas, Manhwa{ID: id.String(), Title: e.ChildText("h4"), Picture: e.ChildAttr("img", "src"), Slug: slug})
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

func getManhwaData(s string) (ManhwaData, []Chapter) {
	var manhwa = ManhwaData{}
	var chapters = []Chapter{}

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/110.0"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("div[class='main-info']", func(e *colly.HTMLElement) {
		id := xid.New()
		chapters = getAllChapters(s)

		manhwa = ManhwaData{ID: id.String(), Title: e.ChildText("div[id='titledesktop']"), Chapters: chapters, Description: e.ChildText("div[itemprop='description']"), Slug: s}
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

	c.Visit("https://elarcpage.com/series/" + s)

	return manhwa, chapters
}

func getAllChapters(s string) []Chapter {
	var chapters = []Chapter{}

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/110.0"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("div[class='main-info']", func(e *colly.HTMLElement) {
		id := xid.New()
		slug := s + "-chapter" + e.ChildAttr("li", "data-num")
		chapter := Chapter{ID: id.String(), Number: e.ChildAttr("li", "data-num"), Date: e.ChildText("span[class='chapterdate']"), Slug: slug}
		fmt.Println(chapter)

		chapters = append(chapters, chapter)
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

	c.Visit("https://elarcpage.com/series/" + s)

	return chapters
}

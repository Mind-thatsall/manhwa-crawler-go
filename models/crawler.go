package models

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly"
	"github.com/rs/xid"
)

func GetManhwas() []Manhwa {
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

func GetManhwaData(s string) ManhwaData {
	var manhwa = ManhwaData{}

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/110.0"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("div[class='main-info']", func(e *colly.HTMLElement) {
		id := xid.New()

		chapters := GetAllChapters(s)

		manhwa = ManhwaData{ID: id.String(), Title: e.ChildText("div[id='titledesktop']"), Description: e.ChildText("div[itemprop='description']"), Slug: s, Chapters: chapters}
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

	return manhwa
}

func GetAllChapters(s string) []Chapter {
	var chapters = []Chapter{}

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/110.0"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("div[id='chapterlist']", func(e *colly.HTMLElement) {
		var allDates = []string{}

		e.ForEach("span[class='chapterdate']", func(i int, h *colly.HTMLElement) {
			allDates = append(allDates, h.Text)
		})

		for i := 0; i < len(e.ChildAttrs("li", "data-num")); i++ {
			id := xid.New()
			chapter := Chapter{ID: id.String(), Number: e.ChildAttrs("li", "data-num")[i], Date: allDates[i], Slug: s + "-chapter" + e.ChildAttrs("li", "data-num")[i]}
			chapters = append(chapters, chapter)
		}

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

func GetChapter(s string) ChapterData {
	var chapter ChapterData

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/110.0"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("div#readerarea", func(e *colly.HTMLElement) {
		chapter = ChapterData{
			ID:       xid.New().String(),
			Pictures: e.ChildAttrs("img", "src"),
		}

		for i := 0; i < len(strings.Split(e.Text, ">")); i++ {
			fmt.Println(strings.Split(e.Text, ">")[i])
		}
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

	c.Visit("https://elarcpage.com/" + s)

	return chapter
}

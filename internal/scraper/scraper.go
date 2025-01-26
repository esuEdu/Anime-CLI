package scraper

import (
	"github.com/gocolly/colly"
	"fmt"
)

func Scrape() {
	c := colly.NewCollector()
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Println(link)
	})
	c.Visit("https://www.crunchyroll.com/")
}
package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func main() {
	fName := "data.csv"
	f, err := os.Create(fName)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
		return
	}

	defer f.Close()

	// c := colly.NewCollector(
	// 	colly.AllowedDomains("en.wikipedia.org"),
	// )
	// c.OnHTML(".mw-parser-output", func(e *colly.HTMLElement) {
	// 	links := e.ChildAttrs("a", "href")
	// 	fmt.Println(links)
	// })

	// c.Visit("https://en.wikipedia.org/wiki/Web_scraping")

	c := colly.NewCollector()
	c.OnHTML("table#customers", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			name := el.ChildText("td:nth-child(1)")
			country := el.ChildText("td:nth-child(2)")
			// fmt.Println(name, country)
			f.WriteString(name + "," + country + " \n")
		})

		fmt.Println("Scraping finished, check file", fName)
	})
	c.Visit("https://www.w3schools.com/html/html_tables.asp")
	writer := csv.NewWriter(f)
	defer writer.Flush()
}

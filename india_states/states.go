package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func main() {
	fname := "states.csv"
	file, err := os.Create(fname)
	if err != nil {
		log.Fatalln("could not create file,", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	c := colly.NewCollector()
	c.OnHTML("table.wikitable:nth-child(77)", func(h *colly.HTMLElement) {
		h.ForEach("tr", func(_ int, row *colly.HTMLElement) {
			rowData := []string{}
			row.ForEach("td", func(_ int, cell *colly.HTMLElement) {
				rowData = append(rowData, cell.Text)
			})
			writer.Write(rowData)
		})
	})
	c.Visit("https://en.wikipedia.org/wiki/States_and_union_territories_of_India")
}

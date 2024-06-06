package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

type TableRow struct {
	Columns []string
}

func cleanText(text string) string {
	return strings.Join(strings.Fields(text), " ")
}

func main() {
	c := colly.NewCollector()

	var rows []TableRow
	var headers []string

	tableCount := 0

	c.OnHTML(".wikitable", func(h *colly.HTMLElement) {
		if tableCount == 3 {
			h.ForEach("tr", func(_ int, el *colly.HTMLElement) {
				var row TableRow
				el.ForEach("th, td", func(_ int, col *colly.HTMLElement) {
					clean := cleanText(col.Text)
					row.Columns = append(row.Columns, clean)
				})
				rows = append(rows, row)
			})
		}
		tableCount++
	})

	c.Visit("https://en.wikipedia.org/wiki/2023%E2%80%9324_Premier_League")

	file, err := os.Create("premier_league.csv")
	if err != nil {
		log.Fatalf("Could not create CSV file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, row := range rows {
		if len(headers) == 0 {
			headers = row.Columns
			if err := writer.Write(headers); err != nil {
				log.Fatalf("Could not write headers to CSV: %v", err)
			}
		} else {
			if err := writer.Write(row.Columns); err != nil {
				log.Fatalf("Could not write row to CSV: %v", err)
			}
		}
	}

	fmt.Println("Utworzono plik premier_league.csv")
}
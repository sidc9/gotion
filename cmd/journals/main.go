package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/sidc9/gotion"
)

func main() {
	apiKey, err := loadAPIKey()
	if err != nil {
		log.Fatal(err)
	}

	gotion.Init(apiKey, gotion.DefaultURL)
	c := gotion.GetClient()

	// TODO
	// add an abstraction layer
	// - input: database name, sort, filter, page-size
	// - output: list of pages, with the ability to get next batch

	db, err := c.SearchDatabaseByTitle("Daily Tracking")
	if err != nil {
		log.Fatalf("database not found: %v", err)
	}

	sort := gotion.NewPropertySort("Date", gotion.SortAscending)
	query := gotion.NewQuery().
		WithSorts([]*gotion.Sort{sort}).
		WithPageSize(5).
		WithLimit(23)

	pgIter := db.NewIterator(query)
	for pgIter.HasNext() {
		p, err := pgIter.GetNext()
		if err != nil {
			log.Fatalf("failed to GetNext: %v", err)
		}

		if highlight, ok := p.Properties.GetRichText("Summary/Highlight"); ok {
			if highlight != "" {
				fmt.Printf("\n%s: %s", p.Title(), highlight)
			} else {
				fmt.Printf(".")
			}
		}
	}
}

func loadAPIKey() (string, error) {
	b, err := ioutil.ReadFile(".env")
	if err != nil {
		return "", err
	}

	return strings.TrimSuffix(string(b), "\n"), nil
}

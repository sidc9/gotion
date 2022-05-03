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

	sorts := gotion.NewPropertySort("Date", gotion.SortAscending)
	query := gotion.NewDBQuery().WithSorts([]*gotion.Sort{sorts})
	query.PageSize = 50
	// pgList, err := c.QueryDatabase(db.ID, query)
	pgIter, err := c.QueryDatabase(db.ID, query)
	// pgIter, err := db.Query(query)
	if err != nil {
		log.Printf("query failed: %v\n", err)
		if gotionErr, ok := err.(*gotion.ErrResponse); ok {
			log.Println(gotionErr.Code, gotionErr.Message)
		}
	}

	for pgIter.HasNext() {
		p := pgIter.GetNext()
		if highlight, ok := p.Properties["Summary/Highlight"]; ok {
			rt := highlight.RichText
			if len(rt) > 0 {
				txt := rt[0].PlainText
				if txt != "" {
					fmt.Printf("%s: %s\n", p.Title(), txt)
				}
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

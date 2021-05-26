// +build examples

package gotion_test

import (
	"fmt"
	"reflect"

	"github.com/sidc9/gotion"
)

func ExampleDatabase_Query() {
	apiKey, err := loadAPIKey()
	if err != nil {
		fmt.Println(err)
	}

	c := gotion.NewClient(apiKey, gotion.DefaultURL)

	db, err := c.GetDatabase("934c6132-4ea7-485e-9b0d-cf1a083e0f3f")
	if err != nil {
		fmt.Println(err)
	}

	s := gotion.NewPropertySort("age", gotion.SortAscending)
	q := gotion.NewDBQuery().WithSorts([]*gotion.Sort{s})
	pages, err := db.Query(q)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(len(pages.Results))

	// Output: 2
}

func ExampleClient_GetPage() {
	apiKey, err := loadAPIKey()
	if err != nil {
		fmt.Println(err)
	}

	c := gotion.NewClient(apiKey, gotion.DefaultURL)

	dbs, err := c.ListDatabases()
	if err != nil {
		fmt.Println(err)
	}

	pages, err := c.QueryDatabase(dbs.Results[0].ID, nil)
	if err != nil {
		fmt.Println(err)
	}

	page, err := c.GetPage(pages.Results[0].ID)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(reflect.DeepEqual(page, pages.Results[0]))

	// Output: true
}

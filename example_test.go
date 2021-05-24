package gotion

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

func ExampleClient_GetPage() {
	apiKey, err := loadAPIKey()
	if err != nil {
		fmt.Println(err)
	}

	c := NewClient(apiKey, DefaultURL)

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

func loadAPIKey() (string, error) {
	b, err := ioutil.ReadFile(".env")
	if err != nil {
		return "", err
	}

	return strings.TrimSuffix(string(b), "\n"), nil
}

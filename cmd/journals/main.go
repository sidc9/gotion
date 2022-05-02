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
	dbs, err := c.ListDatabases()
	if err != nil {
		log.Fatal(err)
	}

	if dbs.Response.HasMore {

	}

	for _, db := range dbs.Results {
		fmt.Println(db.Title[0].PlainText)
	}
}

func loadAPIKey() (string, error) {
	b, err := ioutil.ReadFile(".env")
	if err != nil {
		return "", err
	}

	return strings.TrimSuffix(string(b), "\n"), nil
}

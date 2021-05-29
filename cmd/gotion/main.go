package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/sidc9/gotion"
)

func main() {
	if err := run(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func run() error {
	apiKey, err := loadAPIKey()
	if err != nil {
		return err
	}

	// c := gotion.NewClient(apiKey, gotion.DefaultURL)

	// resp, err := c.ListDatabases()
	// if err != nil {
	//     return err
	// }

	// for _, r := range resp.Results {
	//     fmt.Println(r.Title[0].PlainText, r.ID)
	// }

	gotion.Init(apiKey, gotion.DefaultURL)
	c := gotion.GetClient()

	pg, err := c.GetPage("a0e3feca-85c9-440f-91cc-8c367d6aa9f4")
	if err != nil {
		return err
	}

	content, err := pg.Content()
	if err != nil {
		return err
	}

	children, err := content.Results[1].GetChildren()
	if err != nil {
		return err
	}

	ch, _ := json.Marshal(children)
	fmt.Println(string(ch))

	return nil
}

func loadAPIKey() (string, error) {
	b, err := ioutil.ReadFile(".env")
	if err != nil {
		return "", err
	}

	return strings.TrimSuffix(string(b), "\n"), nil
}

type ObjectType interface {
	TypeName() string
}

var objectTypes = map[string]struct{}{
	ObjectTypeList:     struct{}{},
	ObjectTypeDatabase: struct{}{},
}

const (
	ObjectTypeDatabase = "database"
	ObjectTypeList     = "list"
)

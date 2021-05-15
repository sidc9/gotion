package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/kr/pretty"
)

var (
	API_VERSION_KEY = "Notion-Version"
	API_VERSION_VAL = "2021-05-13"
	baseURL         = "https://api.notion.com/v1/"
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

	c := NewClient(apiKey)

	resp, err := c.ListDatabases()
	if err != nil {
		return err
	}
	pretty.Println(resp)

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

package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	API_KEY         = ""
	API_VERSION_KEY = "Notion-Version"
	API_VERSION_VAL = "2021-05-13"
	baseURL         = "https://api.notion.com/v1/"
)

func main() {
	if err := loadAPIKey(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	ListDatabases()
}

func loadAPIKey() error {
	b, err := ioutil.ReadFile(".env")
	if err != nil {
		return err
	}

	API_KEY = strings.TrimSuffix(string(b), "\n")
	return nil
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

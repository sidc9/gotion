package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/kr/pretty"
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

func ListDatabases() {
	type ListResponse struct {
		Response
		Results []Database `json:"results"`
	}

	var resp ListResponse

	// if err := makeRequest(http.MethodGet, "databases", nil, &resp); err != nil {
	if err := readFile(&resp); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// fmt.Println(resp)
	pretty.Println(resp)
}

func makeRequest(method, path string, body io.Reader, response interface{}) error {
	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", baseURL, path), body)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Add("Authorization", "Bearer "+API_KEY)
	req.Header.Add(API_VERSION_KEY, API_VERSION_VAL)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("http status: %s", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return fmt.Errorf("failed to json unmarshal the response: %w", err)
	}

	return nil
}

type Response struct {
	HasMore    bool   `json:"has_more"`
	NextCursor string `json:"next_cursor"`
	Object     string `json:"object"`
	// Results    interface{} `json:"results"`
}

type ObjectType interface {
	TypeName() string
}

type Database struct {
	CreatedTime    string                 `json:"created_time"`
	ID             string                 `json:"id"`
	LastEditedTime string                 `json:"last_edited_time"`
	Object         string                 `json:"object"`
	Properties     map[string]interface{} `json:"properties"`
}

func (*Database) TypeName() string {
	return ObjectTypeDatabase
}

func (d *Database) String() string {
	return ""
}

var objectTypes = map[string]struct{}{
	ObjectTypeList:     struct{}{},
	ObjectTypeDatabase: struct{}{},
}

const (
	ObjectTypeDatabase = "database"
	ObjectTypeList     = "list"
)

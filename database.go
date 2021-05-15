package main

import (
	"fmt"
	"net/http"
)

type Database struct {
	CreatedTime    string     `json:"created_time"`
	ID             string     `json:"id"`
	LastEditedTime string     `json:"last_edited_time"`
	Object         string     `json:"object"`
	Properties     Properties `json:"properties"`
}

func (*Database) TypeName() string {
	return ObjectTypeDatabase
}

func (d *Database) String() string {
	return ""
}

type DatabaseList struct {
	Response
	Results []*Database `json:"results"`
}

func (c *Client) ListDatabases() (*DatabaseList, error) {
	var dbList DatabaseList

	if err := c.makeRequest(http.MethodGet, "databases", nil, &dbList); err != nil {
		// if err := readFile(&resp, "list_db.txt"); err != nil {
		return nil, err
	}

	return &dbList, nil
}

func (c *Client) GetDatabase(id string) (*Database, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	var db Database
	// if err := c.makeRequest(http.MethodGet, fmt.Sprintf("databases/%s", id), nil, &db); err != nil {
	if err := readFile(&db, "get_db.txt"); err != nil {
		return nil, err
	}

	return &db, nil
}

type Page struct {
	CreatedTime    string                 `json:"created_time"`
	ID             string                 `json:"id"`
	LastEditedTime string                 `json:"last_edited_time"`
	Object         string                 `json:"object"`
	Properties     map[string]interface{} `json:"properties"`
	Archived       bool                   `json:"archived"`
	Parent         map[string]interface{} `json:"parent"`
}

type PageList struct {
	Response
	Results []*Page `json:"results"`
}

// TODO
func (c *Client) QueryDatabase(id string, query *DBQuery) (*PageList, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	var pgList PageList
	if err := c.makeRequest(http.MethodPost, fmt.Sprintf("databases/%s/query", id), query, &pgList); err != nil {
		// if err := readFile(&pgList, "query_db.txt"); err != nil {
		return nil, err
	}

	return &pgList, nil
}

type DBQuery struct {
	Filter *Filter `json:"filter,omitempty"`
	Sorts  *Sort   `json:"sorts,omitempty"`
}

func NewDBQuery() *DBQuery {
	return &DBQuery{}
}

func (q *DBQuery) Do() error {
	return nil
}

func (q *DBQuery) WithFilter(filter *Filter) *DBQuery {
	q.Filter = filter
	return q
}

func (q *DBQuery) WithSort(sort *Sort) *DBQuery {
	q.Sorts = sort
	return q
}

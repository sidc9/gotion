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

type DatabaseList struct {
	Response
	Results []*Database `json:"results"`
}

func (c *Client) ListDatabases() (*DatabaseList, error) {
	var dbList DatabaseList

	if err := c.doRequest(http.MethodGet, "databases", nil, &dbList); err != nil {
		return nil, err
	}

	return &dbList, nil
}

func (c *Client) GetDatabase(id string) (*Database, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	var db Database
	if err := c.doRequest(http.MethodGet, fmt.Sprintf("databases/%s", id), nil, &db); err != nil {
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

func (c *Client) QueryDatabase(id string, query *DBQuery) (*PageList, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	var (
		pgList PageList
		err    error
	)

	if query == nil {
		err = c.doRequest(http.MethodPost, fmt.Sprintf("databases/%s/query", id), nil, &pgList)
	} else {
		err = c.doRequest(http.MethodPost, fmt.Sprintf("databases/%s/query", id), query, &pgList)
	}

	if err != nil {
		return nil, err
	}

	return &pgList, nil
}

type DBQuery struct {
	Filter *NumberFilter `json:"filter,omitempty"`
	Sorts  []*Sort       `json:"sorts,omitempty"`
}

func NewDBQuery() *DBQuery {
	return &DBQuery{}
}

func (q *DBQuery) Do() error {
	return nil
}

func (q *DBQuery) WithFilter(filter *NumberFilter) *DBQuery {
	q.Filter = filter
	return q
}

func (q *DBQuery) WithSorts(sorts []*Sort) *DBQuery {
	q.Sorts = sorts
	return q
}

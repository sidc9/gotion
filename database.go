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

type ListResponse struct {
	Response
	Results []*Database `json:"results"`
}

func (c *Client) ListDatabases() (*ListResponse, error) {
	var resp ListResponse

	if err := c.makeRequest(http.MethodGet, "databases", nil, &resp); err != nil {
		// if err := readFile(&resp, "list_db.txt"); err != nil {
		return nil, err
	}

	return &resp, nil
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

type DBQuery struct {
	Filter *Filter `json:"filter"`
	Sorts  *Sort   `json:"sorts"`
}

func (c *Client) NewDBQuery() *DBQuery {
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

type Sort struct{}

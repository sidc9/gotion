package main

import (
	"encoding/json"
	"fmt"
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

func ListDatabases() (*ListResponse, error) {
	var resp ListResponse

	// if err := makeRequest(http.MethodGet, "databases", nil, &resp); err != nil {
	if err := readFile(&resp, "list_db.txt"); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) GetDatabase(id string) (*Database, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	var db Database
	// if err := makeRequest(http.MethodGet, fmt.Sprintf("databases/%s", id), nil, &db); err != nil {
	if err := readFile(&db, "get_db.txt"); err != nil {
		return nil, err
	}

	return &db, nil
}

type Property struct {
	ID    string      `json:"id"`
	Type  string      `json:"type"`
	Value interface{} `json:-`
}

type NumberProperty struct {
	Property
	Format string `json:"format"`
}

type Properties map[string]*Property

func (p *Property) UnmarshalJSON(b []byte) error {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}

	p.Type = m["type"].(string)
	p.ID = m["id"].(string)
	p.Value = m[p.Type]

	return nil
}

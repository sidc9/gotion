package gotion

import (
	"fmt"
	"net/http"
)

type Database struct {
	ID             string             `json:"id"`
	CreatedTime    string             `json:"created_time"`
	LastEditedTime string             `json:"last_edited_time"`
	Object         string             `json:"object"`
	Properties     DatabaseProperties `json:"properties"`
	Title          []*RichText        `json:"title"`

	c *Client
}

func (d *Database) Query(query *Query) (*PageIter, error) {
	return d.c.QueryDatabase(d.ID, query)
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

	db := &Database{c: c}
	if err := c.doRequest(http.MethodGet, fmt.Sprintf("databases/%s", id), nil, db); err != nil {
		return nil, err
	}

	return db, nil
}

type PageList struct {
	Response
	Results []*Page `json:"results"`
}

func (c *Client) QueryDatabase(id string, query *Query) (*PageIter, error) {
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

	pageIter := &PageIter{
		client:     c,
		hasNext:    pgList.HasMore,
		nextCursor: pgList.NextCursor,
		pages:      pgList.Results,
		index:      0,
	}

	return pageIter, nil
}

func (d *Database) NewIterator(query *Query) *PageIter {
	return &PageIter{
		client:     d.c,
		hasNext:    true, // default
		nextCursor: "",
		index:      0,
		counter:    0,
		pages:      make([]*Page, 0),
		query:      query,
		dbID:       d.ID,
	}
}

func (c *Client) queryDatabase(id string, query *Query) (*PageList, error) {
	if id == "" {
		return nil, fmt.Errorf("id: %w", ErrParamRequired)
	}

	var pgList PageList
	err := c.doRequest(http.MethodPost, fmt.Sprintf("databases/%s/query", id), query, &pgList)
	if err != nil {
		return nil, err
	}

	return &pgList, nil
}

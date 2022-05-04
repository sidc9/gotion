package gotion

import (
	"fmt"
	"net/http"
)

const (
	valueDatabase  = "database"
	propertyObject = "object"
)

type SearchRequest struct {
	PageSize int          `json:"page_size"`
	Filter   searchFilter `json:"filter"`
	Query    string       `json:"query"`
}

type searchFilter struct {
	Value    string `json:"value"`
	Property string `json:"property"`
}

type SearchResponse DatabaseList

func (c *Client) SearchDatabaseByTitle(title string) (*Database, error) {
	req := SearchRequest{
		PageSize: 3,
		Filter: searchFilter{
			Value:    valueDatabase,
			Property: propertyObject,
		},
		Query: title,
	}

	var resp SearchResponse
	if err := c.doRequest(http.MethodPost, "search", req, &resp); err != nil {
		return nil, fmt.Errorf("failed to do request: %v", err)
	}

	if len(resp.Results) == 0 {
		return nil, ErrNotFound
	}

	db := resp.Results[0]
	db.c = c
	return db, nil
}

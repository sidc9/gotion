package gotion

import "fmt"

type PageIter struct {
	client     *Client
	hasNext    bool
	nextCursor string
	index      int
	pages      []*Page
	counter    int
	query      *Query
	dbID       string
}

func (pi *PageIter) HasNext() bool {
	if pi.index >= pi.query.Limit {
		return false
	}
	return pi.hasNext
}

func (pi *PageIter) GetNext() (*Page, error) {
	if pi.index >= pi.query.Limit {
		return nil, ErrMaxLimit
	}

	if pi.index >= len(pi.pages) {
		// make request
		if pi.index+pi.query.PageSize > pi.query.Limit {
			pi.query.PageSize = pi.query.Limit - pi.index
		}
		pages, err := pi.client.queryDatabase(pi.dbID, pi.query)
		if err != nil {
			return nil, fmt.Errorf("failed to query: %w", err)
		}
		pi.pages = append(pi.pages, pages.Results...)
		pi.hasNext = pages.HasMore
		pi.query.StartCursor = pages.NextCursor
	}

	pg := pi.pages[pi.index]
	pi.index++
	return pg, nil
}

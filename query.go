package gotion

import "github.com/sidc9/gotion/filter"

type Query struct {
	Filter      filter.Filter `json:"filter,omitempty"`
	Sorts       []*Sort       `json:"sorts,omitempty"`
	PageSize    int           `json:"page_size,omitempty"`
	StartCursor string        `json:"start_cursor,omitempty"`

	Limit int `json:"-"`
}

func NewQuery() *Query {
	return &Query{
		Limit: 10, // default
	}
}

func (q *Query) WithFilter(filter filter.Filter) *Query {
	q.Filter = filter
	return q
}

func (q *Query) WithSorts(sorts []*Sort) *Query {
	q.Sorts = sorts
	return q
}

func (q *Query) WithLimit(limit int) *Query {
	q.Limit = limit
	return q
}

func (q *Query) WithPageSize(size int) *Query {
	q.PageSize = size
	return q
}

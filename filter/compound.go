package filter

import "fmt"

type OrFilter struct {
	Or []Filter `json:"or,omitempty"`
}

func NewOrFilter(filters ...Filter) (*OrFilter, error) {
	for _, f := range filters {
		if !f.IsValid() {
			// TODO add an interface method to identify a filter
			return nil, fmt.Errorf("a filter is not valid")
		}
	}
	return &OrFilter{Or: filters}, nil
}

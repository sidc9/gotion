package filter

import "fmt"

type OrFilter struct {
	Or []Filter `json:"or,omitempty"`
}

func NewOrFilter(filters ...Filter) (*OrFilter, error) {
	for i, f := range filters {
		if !f.IsValid() {
			return nil, fmt.Errorf("filter[%d] is not valid", i)
		}
	}
	return &OrFilter{Or: filters}, nil
}

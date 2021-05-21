package filter

import (
	"encoding/json"
	"fmt"
)

type OrFilter struct {
	baseFilter
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

func (o *OrFilter) Type() string {
	return "compound OR filter"
}

func (of *OrFilter) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"or": of.Or,
	}
	return json.Marshal(m)
}

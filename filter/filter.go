package filter

import (
	"encoding/json"
)

type Filter interface {
	IsValid() bool
	Condition() string
	Type() string
	Property() string
	Value() interface{}
	json.Marshaler
}

func marshalJSON(f Filter) ([]byte, error) {
	m := make(map[string]interface{})
	m["property"] = f.Property()
	m[f.Type()] = map[string]interface{}{
		f.Condition(): f.Value(),
	}

	return json.Marshal(m)
}

type baseFilter struct {
	property  string
	isSet     bool
	value     interface{}
	condition string
}

func (bf *baseFilter) Condition() string {
	return bf.condition
}

func (bf *baseFilter) Value() interface{} {
	return bf.value
}

func (bf *baseFilter) IsValid() bool {
	return bf.isSet
}

func (bf *baseFilter) Property() string {
	return bf.property
}

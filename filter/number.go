package filter

import (
	"encoding/json"
	"fmt"
)

type NumberFilter struct {
	Property string             `json:"property"`
	Number3  *numberFilterParam `json:"number,omitempty"`
	isSet    bool               `json:"-"`
	param    string             `json:"-"`

	number    int    `json:"-"`
	condition string `json:"-"`
}

func NewNumberFilter(property string) *NumberFilter {
	return &NumberFilter{
		Property: property,
	}
}

func (nf *NumberFilter) IsValid() bool {
	return nf.isSet
}

func (nf *NumberFilter) String() string {
	return fmt.Sprintf("NumberFilter:%s, %s:%d", nf.Property, nf.condition, nf.number)
}

func (nf *NumberFilter) Number() int {
	return nf.number
}

func (nf *NumberFilter) Condition() string {
	return nf.condition
}

const (
	CondEquals             string = "equals"
	CondDoesNotEqual       string = "does_not_equal"
	CondGreaterThan        string = "greater_than"
	CondLessThan           string = "less_than"
	CondGreaterThanOrEqual string = "greater_than_or_equal_to"
	CondLessThanOrEqual    string = "less_than_or_equal_to"
	CondIsEmpty            string = "is_empty"
	CondIsNotEmpty         string = "is_not_empty"
)

type numberFilterParam struct {
	Equals       *int `json:"equals,omitempty"`
	DoesNotEqual *int `json:"does_not_equal,omitempty"`

	GreaterThan *int `json:"greater_than,omitempty"`
	LessThan    *int `json:"less_than,omitempty"`

	GreaterThanOrEqual *int `json:"greater_than_or_equal_to,omitempty"`
	LessThanOrEqual    *int `json:"less_than_or_equal_to,omitempty"`

	IsEmpty    *bool `json:"is_empty,omitempty"`
	IsNotEmpty *bool `json:"is_not_empty,omitempty"`
}

func (nf *NumberFilter) Equals(n int) *NumberFilter {
	nf.isSet = true
	nf.number = n
	nf.condition = CondEquals
	return nf
}

func (nf *NumberFilter) DoesNotEqual(n int) *NumberFilter {
	nf.isSet = true
	nf.number = n
	nf.condition = CondDoesNotEqual
	return nf
}

func (nf *NumberFilter) GreaterThan(n int) *NumberFilter {
	nf.Number3 = &numberFilterParam{
		GreaterThan: &n,
	}
	nf.isSet = true
	return nf
}

func (nf *NumberFilter) LessThan(n int) *NumberFilter {
	nf.Number3 = &numberFilterParam{
		LessThan: &n,
	}
	nf.isSet = true
	return nf
}

func (nf *NumberFilter) GreaterThanOrEqual(n int) *NumberFilter {
	nf.Number3 = &numberFilterParam{
		GreaterThanOrEqual: &n,
	}
	nf.isSet = true
	return nf
}

func (nf *NumberFilter) LessThanOrEqual(n int) *NumberFilter {
	nf.Number3 = &numberFilterParam{
		LessThanOrEqual: &n,
	}
	nf.isSet = true
	return nf
}

func (nf *NumberFilter) IsEmpty() *NumberFilter {
	b := true
	nf.Number3 = &numberFilterParam{
		IsEmpty: &b,
	}
	nf.isSet = true
	return nf
}

func (nf *NumberFilter) IsNotEmpty() *NumberFilter {
	b := true
	nf.Number3 = &numberFilterParam{
		IsNotEmpty: &b,
	}
	nf.isSet = true
	return nf
}

func (nf *NumberFilter) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m["property"] = nf.Property
	m["number"] = map[string]interface{}{
		nf.condition: nf.number,
	}

	return json.Marshal(m)
}

// TODO
func (nf *NumberFilter) UnmarshalJSON(b []byte) error {
	return fmt.Errorf("unimplemented")
}

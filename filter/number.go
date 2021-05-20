package filter

import (
	"encoding/json"
	"fmt"
)

type NumberFilter struct {
	property string `json:"property"`

	isSet     bool
	value     interface{}
	condition string
}

func NewNumberFilter(property string) *NumberFilter {
	return &NumberFilter{
		property: property,
	}
}

func (nf *NumberFilter) IsValid() bool {
	return nf.isSet
}

func (nf *NumberFilter) String() string {
	return fmt.Sprintf("NumberFilter:%s, %s:%d", nf.property, nf.condition, nf.value)
}

func (nf *NumberFilter) Condition() string {
	return nf.condition
}

func (nf *NumberFilter) Property() string {
	return nf.property
}

func (nf *NumberFilter) Type() string {
	return "number"
}

func (nf *NumberFilter) GreaterThanOrEqual(n int) *NumberFilter {
	nf.isSet = true
	nf.value = n
	nf.condition = CondGreaterThanOrEqual
	return nf
}

func (nf *NumberFilter) LessThanOrEqual(n int) *NumberFilter {
	nf.isSet = true
	nf.value = n
	nf.condition = CondLessThanOrEqual
	return nf
}

func (nf *NumberFilter) GreaterThan(n int) *NumberFilter {
	nf.isSet = true
	nf.value = n
	nf.condition = CondGreaterThan
	return nf
}

func (nf *NumberFilter) LessThan(n int) *NumberFilter {
	nf.isSet = true
	nf.value = n
	nf.condition = CondLessThan
	return nf
}

func (nf *NumberFilter) Equals(n int) *NumberFilter {
	nf.isSet = true
	nf.value = n
	nf.condition = CondEquals
	return nf
}

func (nf *NumberFilter) DoesNotEqual(n int) *NumberFilter {
	nf.isSet = true
	nf.value = n
	nf.condition = CondDoesNotEqual
	return nf
}

func (nf *NumberFilter) IsEmpty() *NumberFilter {
	nf.isSet = true
	nf.value = true
	nf.condition = CondIsEmpty
	return nf
}

func (nf *NumberFilter) IsNotEmpty() *NumberFilter {
	nf.isSet = true
	nf.value = true
	nf.condition = CondIsNotEmpty
	return nf
}

func (nf *NumberFilter) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m["property"] = nf.Property()
	m[nf.Type()] = map[string]interface{}{
		nf.condition: nf.value,
	}

	return json.Marshal(m)
}

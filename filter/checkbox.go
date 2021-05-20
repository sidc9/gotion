package filter

import "encoding/json"

type CheckboxFilter struct {
	property  string               `json:"property"`
	Checkbox  *checkboxFilterParam `json:"checkbox,omitempty"`
	isSet     bool                 `json:"-"`
	value     interface{}
	condition string
}

type checkboxFilterParam struct {
	Equals       *bool `json:"equals,omitempty"`
	DoesNotEqual *bool `json:"does_not_equal,omitempty"`
}

func NewCheckboxFilter(property string) *CheckboxFilter {
	return &CheckboxFilter{property: property}
}

func (cf *CheckboxFilter) Condition() string {
	return cf.condition
}

func (cf *CheckboxFilter) Equals(b bool) *CheckboxFilter {
	cf.isSet = true
	cf.condition = CondEquals
	cf.value = b
	return cf
}

func (cf *CheckboxFilter) DoesNotEqual(b bool) *CheckboxFilter {
	cf.Checkbox = &checkboxFilterParam{
		DoesNotEqual: &b,
	}
	cf.isSet = true
	return cf
}

func (cf *CheckboxFilter) IsValid() bool {
	return cf.isSet
}

func (cf *CheckboxFilter) Property() string {
	return cf.property
}

func (cf *CheckboxFilter) Type() string {
	return "checkbox"
}

func (cf *CheckboxFilter) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m["property"] = cf.Property()
	m[cf.Type()] = map[string]interface{}{
		cf.Condition(): cf.value,
	}

	return json.Marshal(m)
}

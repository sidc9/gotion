package filter

type CheckboxFilter struct {
	baseFilter
}

type checkboxFilterParam struct {
	Equals       *bool `json:"equals,omitempty"`
	DoesNotEqual *bool `json:"does_not_equal,omitempty"`
}

func NewCheckboxFilter(property string) *CheckboxFilter {
	return &CheckboxFilter{baseFilter{property: property}}
}

func (cf *CheckboxFilter) Equals(b bool) *CheckboxFilter {
	cf.isSet = true
	cf.condition = CondEquals
	cf.value = b
	return cf
}

func (cf *CheckboxFilter) DoesNotEqual(b bool) *CheckboxFilter {
	cf.isSet = true
	cf.condition = CondDoesNotEqual
	cf.value = b
	return cf
}

func (cf *CheckboxFilter) Type() string {
	return "checkbox"
}

func (cf *CheckboxFilter) MarshalJSON() ([]byte, error) {
	return marshalJSON(cf)
}

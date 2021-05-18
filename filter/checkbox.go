package filter

type CheckboxFilter struct {
	Property string               `json:"property"`
	Checkbox *checkboxFilterParam `json:"checkbox,omitempty"`
}

type checkboxFilterParam struct {
	Equals       *bool `json:"equals,omitempty"`
	DoesNotEqual *bool `json:"does_not_equal,omitempty"`
}

func NewCheckboxFilter(property string) *CheckboxFilter {
	return &CheckboxFilter{Property: property}
}

func (cf *CheckboxFilter) Equals(b bool) *CheckboxFilter {
	cf.Checkbox = &checkboxFilterParam{
		Equals: &b,
	}
	return cf
}

func (cf *CheckboxFilter) DoesNotEqual(b bool) *CheckboxFilter {
	cf.Checkbox = &checkboxFilterParam{
		DoesNotEqual: &b,
	}
	return cf
}

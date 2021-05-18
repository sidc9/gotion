package filter

type MultiSelectFilter struct {
	Property    string                  `json:"property"`
	MultiSelect *multiSelectFilterParam `json:"multiSelect,omitempty"`
}

type multiSelectFilterParam struct {
	Contains       *string `json:"contains,omitempty"`
	DoesNotContain *string `json:"does_not_contain,omitempty"`
	IsEmpty        *bool   `json:"is_empty"`
	IsNotEmpty     *bool   `json:"is_not_empty"`
}

func NewMultiSelectFilter(property string) *MultiSelectFilter {
	return &MultiSelectFilter{Property: property}
}

func (sf *MultiSelectFilter) Contains(s string) *MultiSelectFilter {
	sf.MultiSelect = &multiSelectFilterParam{
		Contains: &s,
	}
	return sf
}

func (sf *MultiSelectFilter) DoesNotContain(s string) *MultiSelectFilter {
	sf.MultiSelect = &multiSelectFilterParam{
		DoesNotContain: &s,
	}
	return sf
}

func (sf *MultiSelectFilter) IsEmpty() *MultiSelectFilter {
	b := true
	sf.MultiSelect = &multiSelectFilterParam{
		IsEmpty: &b,
	}
	return sf
}

func (sf *MultiSelectFilter) IsNotEmpty() *MultiSelectFilter {
	b := true
	sf.MultiSelect = &multiSelectFilterParam{
		IsNotEmpty: &b,
	}
	return sf
}

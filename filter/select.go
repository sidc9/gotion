package filter

type SelectFilter struct {
	Property string             `json:"property"`
	Select   *selectFilterParam `json:"select,omitempty"`
}

type selectFilterParam struct {
	Equals       *string `json:"equals,omitempty"`
	DoesNotEqual *string `json:"does_not_equal,omitempty"`
	IsEmpty      *bool   `json:"is_empty"`
	IsNotEmpty   *bool   `json:"is_not_empty"`
}

func NewSelectFilter(property string) *SelectFilter {
	return &SelectFilter{Property: property}
}

func (sf *SelectFilter) Equals(s string) *SelectFilter {
	sf.Select = &selectFilterParam{
		Equals: &s,
	}
	return sf
}

func (sf *SelectFilter) DoesNotEqual(s string) *SelectFilter {
	sf.Select = &selectFilterParam{
		DoesNotEqual: &s,
	}
	return sf
}

func (sf *SelectFilter) IsEmpty() *SelectFilter {
	b := true
	sf.Select = &selectFilterParam{
		IsEmpty: &b,
	}
	return sf
}

func (sf *SelectFilter) IsNotEmpty() *SelectFilter {
	b := true
	sf.Select = &selectFilterParam{
		IsNotEmpty: &b,
	}
	return sf
}

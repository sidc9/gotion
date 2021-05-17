package gotion

type NumberFilter struct {
	Property string             `json:"property"`
	Number   *numberFilterParam `json:"number,omitempty"`
}

func NewNumberFilter(property string) *NumberFilter {
	return &NumberFilter{
		Property: property,
	}
}

type numberFilterParam struct {
	Equals       *int `json:"equals,omitempty"`
	DoesNotEqual *int `json:"does_not_equal,omitempty"`

	GreaterThan *int `json:"greater_than,omitempty"`
	LessThan    *int `json:"less_than,omitempty"`

	GreaterThanOrEqual *int `json:"greater_than_or_equal_to,omitempty"`
	LessThanOrEqual    *int `json:"less_than_or_equal_to,omitempty"`

	IsEmpty    *bool `json:"is_empty"`
	IsNotEmpty *bool `json:"is_not_empty"`
}

func (nf *NumberFilter) Equals(n int) *NumberFilter {
	nf.Number = &numberFilterParam{
		Equals: &n,
	}
	return nf
}

func (nf *NumberFilter) DoesNotEqual(n int) *NumberFilter {
	nf.Number = &numberFilterParam{
		DoesNotEqual: &n,
	}
	return nf
}

func (nf *NumberFilter) GreaterThan(n int) *NumberFilter {
	nf.Number = &numberFilterParam{
		GreaterThan: &n,
	}
	return nf
}

func (nf *NumberFilter) LessThan(n int) *NumberFilter {
	nf.Number = &numberFilterParam{
		LessThan: &n,
	}
	return nf
}

func (nf *NumberFilter) GreaterThanOrEqual(n int) *NumberFilter {
	nf.Number = &numberFilterParam{
		GreaterThanOrEqual: &n,
	}
	return nf
}

func (nf *NumberFilter) LessThanOrEqual(n int) *NumberFilter {
	nf.Number = &numberFilterParam{
		LessThanOrEqual: &n,
	}
	return nf
}

func (nf *NumberFilter) IsEmpty() *NumberFilter {
	nf.Number = &numberFilterParam{
		IsEmpty: boolptr(true),
	}
	return nf
}

func (nf *NumberFilter) IsNotEmpty() *NumberFilter {
	nf.Number = &numberFilterParam{
		IsNotEmpty: boolptr(true),
	}
	return nf
}

func boolptr(b bool) *bool {
	return &b
}

/* type CheckboxFilter struct {
	Property    string `json:"property"`
	ParamEquals bool   `json:"equals"`
}

func NewCheckboxFilter(property string) *CheckboxFilter {
	return &CheckboxFilter{Property: property}
}

func (c *CheckboxFilter) Equals(eq bool) *CheckboxFilter {
	c.ParamEquals = eq
	return c
} */

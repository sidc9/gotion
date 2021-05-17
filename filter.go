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
	GreaterThanOrEqual *int `json:"greater_than_or_equal_to,omitempty"`
	GreaterThan        *int `json:"greater_than,omitempty"`
	Equals             *int `json:"equals,omitempty"`
}

func (nf *NumberFilter) GreaterThanOrEqual(n int) *NumberFilter {
	nf.Number = &numberFilterParam{
		GreaterThanOrEqual: &n,
	}
	return nf
}

func (nf *NumberFilter) GreaterThan(n int) *NumberFilter {
	nf.Number = &numberFilterParam{
		GreaterThan: &n,
	}
	return nf
}

func (nf *NumberFilter) Equals(n int) *NumberFilter {
	nf.Number = &numberFilterParam{
		Equals: &n,
	}
	return nf
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

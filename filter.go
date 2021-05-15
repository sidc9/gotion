package main

type Filter struct {
	Property string          `json:"property"`
	Number   NumberFilter    `json:"number,omitempty"`
	Checkbox map[string]bool `json:"checkbox,omitempty"`
}

type NumberFilter struct {
	Property                string `json:"property"`
	ParamGreaterThanOrEqual *int   `json:"greater_than_or_equal_to,omitempty"`
	ParamGreaterThan        *int   `json:"greater_than_to,omitempty"`
	ParamEquals             *int   `json:"equals,omitempty"`
	paramSet                bool   `json:"-"`
}

func NewNumberFilter(propName string) *NumberFilter {
	return &NumberFilter{Property: propName}
}

func (nf *NumberFilter) GreaterThanOrEqual(n int) *NumberFilter {
	if nf.paramSet {
		return nf
	}

	nf.ParamGreaterThanOrEqual = &n
	nf.paramSet = true
	return nf
}

func (nf *NumberFilter) GreaterThan(n int) *NumberFilter {
	if nf.paramSet {
		return nf
	}

	nf.ParamGreaterThan = &n
	nf.paramSet = true
	return nf
}

func (nf *NumberFilter) Equals(n int) *NumberFilter {
	if nf.paramSet {
		return nf
	}

	nf.ParamEquals = &n
	nf.paramSet = true
	return nf
}

type CheckboxFilter struct {
	Property    string `json:"property"`
	ParamEquals bool   `json:"equals"`
}

func NewCheckboxFilter(property string) *CheckboxFilter {
	return &CheckboxFilter{Property: property}
}

func (c *CheckboxFilter) Equals(eq bool) *CheckboxFilter {
	c.ParamEquals = eq
	return c
}

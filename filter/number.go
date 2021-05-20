package filter

type NumberFilter struct {
	Property string             `json:"property"`
	Number   *numberFilterParam `json:"number,omitempty"`
	isSet    bool               `json:"-"`
	param    string             `json:"-"`
}

func NewNumberFilter(property string) *NumberFilter {
	return &NumberFilter{
		Property: property,
	}
}

func (nf *NumberFilter) IsValid() bool {
	return nf.isSet
}

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
	nf.Number = &numberFilterParam{
		Equals: &n,
	}
	nf.isSet = true
	return nf
}

func (nf *NumberFilter) DoesNotEqual(n int) *NumberFilter {
	nf.Number = &numberFilterParam{
		DoesNotEqual: &n,
	}
	nf.isSet = true
	return nf
}

func (nf *NumberFilter) GreaterThan(n int) *NumberFilter {
	nf.Number = &numberFilterParam{
		GreaterThan: &n,
	}
	nf.isSet = true
	return nf
}

func (nf *NumberFilter) LessThan(n int) *NumberFilter {
	nf.Number = &numberFilterParam{
		LessThan: &n,
	}
	nf.isSet = true
	return nf
}

func (nf *NumberFilter) GreaterThanOrEqual(n int) *NumberFilter {
	nf.Number = &numberFilterParam{
		GreaterThanOrEqual: &n,
	}
	nf.isSet = true
	return nf
}

func (nf *NumberFilter) LessThanOrEqual(n int) *NumberFilter {
	nf.Number = &numberFilterParam{
		LessThanOrEqual: &n,
	}
	nf.isSet = true
	return nf
}

func (nf *NumberFilter) IsEmpty() *NumberFilter {
	b := true
	nf.Number = &numberFilterParam{
		IsEmpty: &b,
	}
	nf.isSet = true
	return nf
}

func (nf *NumberFilter) IsNotEmpty() *NumberFilter {
	b := true
	nf.Number = &numberFilterParam{
		IsNotEmpty: &b,
	}
	nf.isSet = true
	return nf
}

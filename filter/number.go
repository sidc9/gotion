package filter

type NumberFilter struct {
	baseFilter
}

func NewNumberFilter(property string) *NumberFilter {
	return &NumberFilter{
		baseFilter{
			property: property,
		},
	}
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
	return marshalJSON(nf)
}

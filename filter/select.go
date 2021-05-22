package filter

type SelectFilter struct {
	baseFilter
}

func (sf *SelectFilter) Type() string {
	return "select"
}

func NewSelectFilter(property string) *SelectFilter {
	return &SelectFilter{
		baseFilter{
			property: property,
		},
	}
}

func (sf *SelectFilter) Equals(s string) *SelectFilter {
	sf.isSet = true
	sf.value = s
	sf.condition = CondEquals
	return sf
}

func (sf *SelectFilter) DoesNotEqual(s string) *SelectFilter {
	sf.isSet = true
	sf.value = s
	sf.condition = CondDoesNotEqual
	return sf
}

func (sf *SelectFilter) IsEmpty() *SelectFilter {
	sf.isSet = true
	sf.value = true
	sf.condition = CondIsEmpty
	return sf
}

func (sf *SelectFilter) IsNotEmpty() *SelectFilter {
	sf.isSet = true
	sf.value = true
	sf.condition = CondIsNotEmpty
	return sf
}

func (sf *SelectFilter) MarshalJSON() ([]byte, error) {
	return marshalJSON(sf)
}

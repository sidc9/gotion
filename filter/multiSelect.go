package filter

type MultiSelectFilter struct {
	baseFilter
}

func NewMultiSelectFilter(property string) *MultiSelectFilter {
	return &MultiSelectFilter{baseFilter{property: property}}
}

func (ms *MultiSelectFilter) Type() string {
	return "multi_select"
}

func (ms *MultiSelectFilter) MarshalJSON() ([]byte, error) {
	return marshalJSON(ms)
}

func (sf *MultiSelectFilter) Contains(s string) *MultiSelectFilter {
	sf.isSet = true
	sf.value = s
	sf.condition = CondContains
	return sf
}

func (sf *MultiSelectFilter) DoesNotContain(s string) *MultiSelectFilter {
	sf.isSet = true
	sf.value = s
	sf.condition = CondDoesNotContain
	return sf
}

func (sf *MultiSelectFilter) IsEmpty() *MultiSelectFilter {
	sf.isSet = true
	sf.value = true
	sf.condition = CondIsEmpty
	return sf
}

func (sf *MultiSelectFilter) IsNotEmpty() *MultiSelectFilter {
	sf.isSet = true
	sf.value = true
	sf.condition = CondIsNotEmpty
	return sf
}

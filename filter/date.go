package filter

type DateFilter struct {
	baseFilter
}

func (df *DateFilter) Type() string {
	return "date"
}

func (df *DateFilter) MarshalJSON() ([]byte, error) {
	return marshalJSON(df)
}

func NewDateFilter(property string) *DateFilter {
	return &DateFilter{baseFilter{property: property}}
}

func (df *DateFilter) Equals(s string) *DateFilter {
	df.isSet = true
	df.value = s
	df.condition = CondEquals
	return df
}

func (df *DateFilter) Before(s string) *DateFilter {
	df.isSet = true
	df.value = s
	df.condition = CondBefore
	return df
}

func (df *DateFilter) After(s string) *DateFilter {
	df.isSet = true
	df.value = s
	df.condition = CondAfter
	return df
}

func (df *DateFilter) IsEmpty() *DateFilter {
	df.isSet = true
	df.value = true
	df.condition = CondIsEmpty
	return df
}

func (df *DateFilter) IsNotEmpty() *DateFilter {
	df.isSet = true
	df.value = true
	df.condition = CondIsNotEmpty
	return df
}

func (df *DateFilter) OnOrBefore(s string) *DateFilter {
	df.isSet = true
	df.value = s
	df.condition = CondOnOrBefore
	return df
}

func (df *DateFilter) OnOrAfter(s string) *DateFilter {
	df.isSet = true
	df.value = s
	df.condition = CondOnOrAfter
	return df
}

func (df *DateFilter) PastWeek(s string) *DateFilter {
	df.isSet = true
	df.value = struct{}{}
	df.condition = CondPastWeek
	return df
}

func (df *DateFilter) PastMonth(s string) *DateFilter {
	df.isSet = true
	df.value = struct{}{}
	df.condition = CondPastMonth
	return df
}

func (df *DateFilter) PastYear(s string) *DateFilter {
	df.isSet = true
	df.value = struct{}{}
	df.condition = CondPastYear
	return df
}

func (df *DateFilter) NextWeek(s string) *DateFilter {
	df.isSet = true
	df.value = struct{}{}
	df.condition = CondNextWeek
	return df
}

func (df *DateFilter) NextMonth(s string) *DateFilter {
	df.isSet = true
	df.value = struct{}{}
	df.condition = CondNextMonth
	return df
}

func (df *DateFilter) NextYear(s string) *DateFilter {
	df.isSet = true
	df.value = struct{}{}
	df.condition = CondNextYear
	return df
}

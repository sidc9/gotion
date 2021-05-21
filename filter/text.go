package filter

type TextFilter struct {
	baseFilter
}

func NewTextFilter(property string) *TextFilter {
	return &TextFilter{
		baseFilter{
			property: property,
		},
	}
}

func (tf *TextFilter) Type() string {
	return "text"
}

func (tf *TextFilter) Equals(s string) *TextFilter {
	tf.isSet = true
	tf.value = s
	tf.condition = CondEquals
	return tf
}

func (tf *TextFilter) DoesNotEqual(s string) *TextFilter {
	tf.isSet = true
	tf.value = s
	tf.condition = CondDoesNotEqual
	return tf
}

func (tf *TextFilter) IsEmpty() *TextFilter {
	tf.isSet = true
	tf.value = true
	tf.condition = CondIsEmpty
	return tf
}

func (tf *TextFilter) IsNotEmpty() *TextFilter {
	tf.isSet = true
	tf.value = true
	tf.condition = CondIsNotEmpty
	return tf
}

func (tf *TextFilter) Contains(s string) *TextFilter {
	tf.isSet = true
	tf.value = s
	tf.condition = CondContains
	return tf
}

func (tf *TextFilter) DoesNotContain(s string) *TextFilter {
	tf.isSet = true
	tf.value = s
	tf.condition = CondDoesNotContain
	return tf
}

func (tf *TextFilter) StartsWith(s string) *TextFilter {
	tf.isSet = true
	tf.value = s
	tf.condition = CondStartsWith
	return tf
}

func (tf *TextFilter) EndsWith(s string) *TextFilter {
	tf.isSet = true
	tf.value = s
	tf.condition = CondEndsWith
	return tf
}

func (tf *TextFilter) MarshalJSON() ([]byte, error) {
	return marshalJSON(tf)
}

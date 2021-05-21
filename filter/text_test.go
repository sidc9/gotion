package filter_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/sidc9/gotion/filter"
)

func TestTextFilter(t *testing.T) {
	is := is.New(t)

	t.Run("set Equals", func(t *testing.T) {
		tf := filter.NewTextFilter("age").Equals("hello")
		is.Equal(tf.Condition(), "equals")
		is.Equal(tf.Property(), "age")

		checkJSON(t, tf, `{"property":"age","text":{"equals":"hello"}}`)
	})

	t.Run("set NotEqual", func(t *testing.T) {
		tf := filter.NewTextFilter("age").DoesNotEqual("hello")
		is.Equal(tf.Condition(), "does_not_equal")
		is.Equal(tf.Property(), "age")

		checkJSON(t, tf, `{"property":"age","text":{"does_not_equal":"hello"}}`)
	})

	t.Run("set IsEmpty", func(t *testing.T) {
		tf := filter.NewTextFilter("age").IsEmpty()
		is.Equal(tf.Condition(), "is_empty")
		is.Equal(tf.Property(), "age")

		checkJSON(t, tf, `{"property":"age","text":{"is_empty":true}}`)
	})

	t.Run("set IsNotEmpty", func(t *testing.T) {
		tf := filter.NewTextFilter("age").IsNotEmpty()
		is.Equal(tf.Condition(), "is_not_empty")
		is.Equal(tf.Property(), "age")

		checkJSON(t, tf, `{"property":"age","text":{"is_not_empty":true}}`)
	})
}

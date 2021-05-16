package main

/* func TestFilter(t *testing.T) {
	is := is.New(t)

	t.Run("number filter", func(t *testing.T) {
		t.Run("set greaterThanOrEqual", func(t *testing.T) {
			nf := NewNumberFilter("age").GreaterThanOrEqual(2)
			is.Equal(*nf.ParamGreaterThanOrEqual, 2)
			is.Equal(nf.Property, "age")
		})

		t.Run("set greaterThan", func(t *testing.T) {
			nf := NewNumberFilter("age").GreaterThan(2)
			is.Equal(*nf.ParamGreaterThan, 2)
			is.Equal(nf.Property, "age")
		})

		t.Run("set Equals", func(t *testing.T) {
			nf := NewNumberFilter("age").Equals(2)
			is.Equal(*nf.ParamEquals, 2)
			is.Equal(nf.Property, "age")
		})

		t.Run("does not set greaterThanOrEqual if already set", func(t *testing.T) {
			nf := NewNumberFilter("age").GreaterThanOrEqual(5)
			nf.GreaterThanOrEqual(2)
			is.Equal(*nf.ParamGreaterThanOrEqual, 5)
		})

		t.Run("does not set greaterThan if already set", func(t *testing.T) {
			nf := NewNumberFilter("age").GreaterThan(5)
			nf.GreaterThan(2)
			is.Equal(*nf.ParamGreaterThan, 5)
		})

		t.Run("does not set Equals if already set", func(t *testing.T) {
			nf := NewNumberFilter("age").Equals(5)
			nf.Equals(2)
			is.Equal(*nf.ParamEquals, 5)
		})
	})

	t.Run("checkbox filter", func(t *testing.T) {
		c := NewCheckboxFilter("available").Equals(true)
		is.Equal(c.Property, "available")
		is.True(c.ParamEquals)
	})
} */

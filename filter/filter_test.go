package filter_test

import (
	"encoding/json"
	"testing"

	"github.com/matryer/is"
)

func checkJSON(t *testing.T, val interface{}, want string) {
	t.Helper()
	is := is.NewRelaxed(t)
	b, err := json.Marshal(val)
	is.NoErr(err)
	is.Equal(want, string(b))
}

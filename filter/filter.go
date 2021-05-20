package filter

import (
	"encoding/json"
	"fmt"
)

type Filter interface {
	IsValid() bool
	Condition() string
	Type() string
	Property() string
	fmt.Stringer
}

func MarshalJSON(f Filter) ([]byte, error) {
	m := make(map[string]interface{})
	m["property"] = f.Property()
	m[f.Type()] = map[string]interface{}{
		f.Condition(): f.value,
	}

	return json.Marshal(m)
}

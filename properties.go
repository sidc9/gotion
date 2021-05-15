package main

import "encoding/json"

type Properties []*Property

type Property struct {
	Name        string          `json:"-"`
	ID          string          `json:"id"`
	Type        PropertyType    `json:"type"`
	Number      PropNumber      `json:"number"`
	MultiSelect PropMultiSelect `json:"multi_select"`
}

type PropNumber struct {
	Format string `json:"format"`
}

type PropMultiSelect struct {
	Options []interface{} `json:"options"`
}

type PropertyType string

func (p *Properties) UnmarshalJSON(b []byte) error {
	var m map[string]json.RawMessage
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}

	*p = []*Property{}
	for k, v := range m {
		var prop Property
		if err := json.Unmarshal(v, &prop); err != nil {
			return err
		}

		prop.Name = k
		*p = append(*p, &prop)
	}

	return nil
}

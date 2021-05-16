package main

type DatabaseProperties map[string]*DatabaseProperty

type DatabaseProperty struct {
	Name        string            `json:"-"`
	ID          string            `json:"id"`
	Type        string            `json:"type"`
	Number      DBPropNumber      `json:"number"`
	MultiSelect DBPropMultiSelect `json:"multi_select"`
}

type DBPropNumber struct {
	Format string `json:"format"`
}

type DBPropMultiSelect struct {
	Options []interface{} `json:"options"`
}

type PageProperties map[string]*PageProperty

type PageProperty struct {
	Name   string `json:"-"`
	ID     string `json:"id"`
	Type   string `json:"type"`
	Number int    `json:"number"`
	// MultiSelect DBPropMultiSelect `json:"multi_select"`
}

package gotion

type DatabaseProperties map[string]*DatabaseProperty

type DatabaseProperty struct {
	Name        string            `json:"-"`
	ID          string            `json:"id"`
	Type        string            `json:"type"`
	Number      DBPropNumber      `json:"number"`
	MultiSelect DBPropMultiSelect `json:"multi_select"`
	Select      DBPropSelect      `json:"select"`
}

type DBPropNumber struct {
	Format string `json:"format"`
}

type DBPropMultiSelect struct {
	Options []*Option `json:"options"`
}

type DBPropSelect struct {
	Options []*Option `json:"options"`
}

type Option struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type PageProperties map[string]*PageProperty

type PageProperty struct {
	Name   string `json:"-"`
	ID     string `json:"id"`
	Type   string `json:"type"`
	Number int    `json:"number"`
}

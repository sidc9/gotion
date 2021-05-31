package gotion

type DatabaseProperties map[string]*DatabaseProperty

type DatabaseProperty struct {
	Name        string            `json:"-"`
	Title       RichText          `json:"title"`
	ID          string            `json:"id"`
	Type        string            `json:"type"`
	Number      DBPropNumber      `json:"number"`
	MultiSelect DBPropMultiSelect `json:"multi_select"`
	Select      DBPropSelect      `json:"select"`
	RichText    struct{}          `json:"rich_text"`
	Checkbox    struct{}          `json:"checkbox"`
	CreatedTime struct{}          `json:"created_time"`
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

type DBPropText struct {
	Content string `json:"content"`
}

type Option struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type PageProperties map[string]*PageProperty

type PageProperty struct {
	Title       []*RichText            `json:"title,omitempty"`
	ID          string                 `json:"id,omitempty"`
	Type        string                 `json:"type,omitempty"`
	Number      int                    `json:"number,omitempty"`
	Select      *PagePropSelect        `json:"select,omitempty"`
	MultiSelect []*PagePropMultiSelect `json:"multi_select,omitempty"`
	RichText    []*RichText            `json:"rich_text,omitempty"`
	Checkbox    bool                   `json:"checkbox,omitempty"`
	CreatedTime string                 `json:"created_time,omitempty"`
}

type PagePropName struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type PagePropMultiSelect struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type PagePropSelect struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type RichText struct {
	Type        string      `json:"type"`
	Text        Text        `json:"text"`
	Annotations *Annotation `json:"annotations,omitempty"`
	PlainText   string      `json:"plain_text"`
	Href        string      `json:"href"`
}

type Text struct {
	Content string `json:"content"`
	Link    string `json:"link,omitempty"`
}

type Annotation struct {
	Bold          bool   `json:"bold"`
	Italic        bool   `json:"italic"`
	Strikethrough bool   `json:"strikethrough"`
	Underline     bool   `json:"underline"`
	Code          bool   `json:"code"`
	Color         string `json:"color"`
}

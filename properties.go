package gotion

type DatabaseProperties map[string]*DatabaseProperty

type DatabaseProperty struct {
	Name        string                 `json:"-"`
	Title       []Title                `json:"title"`
	ID          string                 `json:"id"`
	Type        string                 `json:"type"`
	Number      DBPropNumber           `json:"number"`
	MultiSelect DBPropMultiSelect      `json:"multi_select"`
	Select      DBPropSelect           `json:"select"`
	RichText    map[string]interface{} `json:"rich_text"`
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
	Title       []Title               `json:"title"`
	ID          string                `json:"id"`
	Type        string                `json:"type"`
	Number      int                   `json:"number"`
	Select      PagePropSelect        `json:"select"`
	MultiSelect []PagePropMultiSelect `json:"multi_select"`
	RichText    []*PagePropRichText   `json:"rich_text"`
}

type PagePropName struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type Title struct {
	Type        string     `json:"type"`
	Text        Text       `json:"text"`
	Annotations Annotation `json:"annotations"`
	PlainText   string     `json:"plain_text"`
	Href        string     `json:"href"`
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

type PagePropRichText struct {
	Type        string     `json:"type"`
	Text        Text       `json:"text"`
	Annotations Annotation `json:"annotations"`
	PlainText   string     `json:"plain_text"`
	Href        string     `json:"href"`
}

type Text struct {
	Content string `json:"content"`
	Link    string `json:"link"`
}

type Annotation struct {
	Bold          bool   `json:"bold"`
	Italic        bool   `json:"italic"`
	Strikethrough bool   `json:"strikethrough"`
	Underline     bool   `json:"underline"`
	Code          bool   `json:"code"`
	Color         string `json:"color"`
}

// "name":{"id":"title","type":"title","title":[{"type":"text","text":{"content":"mary","link":null},"annotations":{"bold":false,"italic":false,"strikethrough":false,"underline":false,"code":false,"color":"default"},"plain_text":"mary","href":null}]}

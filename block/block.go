package block

import "encoding/json"

type Block struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

type HeaderData struct {
	Text  string `json:"text"`
	Level int    `json:"level"`
}

type ParagraphData struct {
	Text string `json:"text"`
}

type RawData struct {
	HTML string `json:"html"`
}

type CheckListData struct {
	Items []CheckListItem `json:"items"`
}

type CheckListItem struct {
	Text    string `json:"text"`
	Checked bool   `json:"checked"`
}

type ListData struct {
	Style string     `json:"style"`
	Items []ListItem `json:"items"`
}

type ListItem struct {
	Content string `json:"content"`
}

type QuoteData struct {
	Text      string `json:"text"`
	Caption   string `json:"caption"`
	Alignment string `json:"alignment"`
}

type TableData struct {
	WithHeadings bool       `json:"withHeadings"`
	Content      [][]string `json:"content"`
}

type CodeData struct {
	Code string `json:"code"`
}

type DelimiterData struct{}

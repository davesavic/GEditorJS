package converter

import (
	"encoding/json"
	"fmt"
	"github.com/davesavic/GEditorJS/block"
	"strings"
)

type HTMLBlockConverter struct{}

func (h HTMLBlockConverter) Convert(b block.Block) (string, error) {
	converter := GetBlockConverter(b.Type)
	if converter == nil {
		return "", fmt.Errorf("unsupported block type: %s", b.Type)
	}

	return converter.Convert(b)
}

func GetBlockConverter(blockType string) BlockConverter {
	switch blockType {
	case "header":
		return HTMLHeaderBlockConverter{}
	case "paragraph":
		return HTMLParagraphBlockConverter{}
	case "raw":
		return HTMLRawBlockConverter{}
	case "checklist":
		return HTMLChecklistBlockConverter{}
	case "list":
		return HTMLListBlockConverter{}
	case "quote":
		return HTMLQuoteBlockConverter{}
	case "table":
		return HTMLTableBlockConverter{}
	case "code":
		return HTMLCodeBlockConverter{}
	case "delimiter":
		return HTMLDelimiterBlockConverter{}
	default:
		return nil
	}
}

type HTMLHeaderBlockConverter struct{}

func (h HTMLHeaderBlockConverter) Convert(b block.Block) (string, error) {
	var data block.HeaderData
	err := json.Unmarshal(b.Data, &data)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("<h%d>%s</h%d>", data.Level, data.Text, data.Level), nil
}

type HTMLParagraphBlockConverter struct{}

func (h HTMLParagraphBlockConverter) Convert(b block.Block) (string, error) {
	var data block.ParagraphData
	err := json.Unmarshal(b.Data, &data)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("<p>%s</p>", data.Text), nil
}

type HTMLRawBlockConverter struct{}

func (h HTMLRawBlockConverter) Convert(b block.Block) (string, error) {
	var data block.RawData
	err := json.Unmarshal(b.Data, &data)
	if err != nil {
		return "", err
	}

	return data.HTML, nil
}

type HTMLChecklistBlockConverter struct{}

func (h HTMLChecklistBlockConverter) Convert(b block.Block) (string, error) {
	var data block.CheckListData
	err := json.Unmarshal(b.Data, &data)
	if err != nil {
		return "", err
	}

	var o strings.Builder
	o.WriteString("<ul>")
	for _, item := range data.Items {
		o.WriteString(fmt.Sprintf("<li><input type=\"checkbox\" disabled%s>%s</li>", func() string {
			if item.Checked {
				return " checked"
			}

			return ""
		}(), item.Text))
	}
	o.WriteString("</ul>")

	return o.String(), nil
}

type HTMLListBlockConverter struct{}

func (h HTMLListBlockConverter) Convert(b block.Block) (string, error) {
	var data block.ListData
	err := json.Unmarshal(b.Data, &data)
	if err != nil {
		return "", err
	}

	var o strings.Builder
	o.WriteString(fmt.Sprintf("<ul style=\"list-style-type: %s\">", data.Style))
	for _, item := range data.Items {
		o.WriteString(fmt.Sprintf("<li>%s</li>", item.Content))
	}
	o.WriteString("</ul>")

	return o.String(), nil
}

type HTMLQuoteBlockConverter struct{}

func (h HTMLQuoteBlockConverter) Convert(b block.Block) (string, error) {
	var data block.QuoteData
	err := json.Unmarshal(b.Data, &data)
	if err != nil {
		return "", err
	}

	if data.Alignment == "" {
		data.Alignment = "center"
	}

	return fmt.Sprintf("<blockquote style=\"text-align: %s\"><p>%s</p><footer>%s</footer></blockquote>", data.Alignment, data.Text, data.Caption), nil
}

type HTMLTableBlockConverter struct{}

func (h HTMLTableBlockConverter) Convert(b block.Block) (string, error) {
	var data block.TableData
	err := json.Unmarshal(b.Data, &data)
	if err != nil {
		return "", err
	}

	var o strings.Builder
	o.WriteString("<table>")

	tbodyOpened := false

	for i, row := range data.Content {
		if i == 0 && data.WithHeadings {
			o.WriteString("<thead>")
		} else if i == 0 || !tbodyOpened {
			o.WriteString("<tbody>")
			tbodyOpened = true
		}

		o.WriteString("<tr>")
		for _, cell := range row {
			if i == 0 && data.WithHeadings {
				o.WriteString(fmt.Sprintf("<th>%s</th>", cell))
			} else {
				o.WriteString(fmt.Sprintf("<td>%s</td>", cell))
			}
		}
		o.WriteString("</tr>")

		if i == 0 && data.WithHeadings {
			o.WriteString("</thead>")
		}
	}

	if tbodyOpened {
		o.WriteString("</tbody>")
	}

	o.WriteString("</table>")

	return o.String(), nil
}

type HTMLCodeBlockConverter struct{}

func (h HTMLCodeBlockConverter) Convert(b block.Block) (string, error) {
	var data block.CodeData
	err := json.Unmarshal(b.Data, &data)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("<pre><code>%s</code></pre>", data.Code), nil
}

type HTMLDelimiterBlockConverter struct{}

func (h HTMLDelimiterBlockConverter) Convert(b block.Block) (string, error) {
	return "<hr>", nil
}

package GEditorJS

import (
	"encoding/json"
	"github.com/davesavic/GEditorJS/block"
	"github.com/davesavic/GEditorJS/converter"
	"strings"
)

type EditorJSOutput struct {
	Time    int64         `json:"time"`
	Blocks  []block.Block `json:"blocks"`
	Version string        `json:"version"`
}

type GEditorJS struct {
	RawContent string
}

func New() *GEditorJS {
	return &GEditorJS{}
}

func (g *GEditorJS) SetRawContent(rawContent string) {
	g.RawContent = rawContent
}

func (g *GEditorJS) Parse(converter converter.BlockConverter) (string, error) {
	if g.RawContent == "" {
		return "", nil
	}

	var output EditorJSOutput
	err := json.Unmarshal([]byte(g.RawContent), &output)
	if err != nil {
		return "", err
	}

	var o strings.Builder
	for _, b := range output.Blocks {
		convertedBlock, err := converter.Convert(b)
		if err != nil {
			return "", err
		}

		o.WriteString(convertedBlock + "\n")
	}

	return o.String(), nil
}

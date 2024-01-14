package converter_test

import (
	"encoding/json"
	"github.com/davesavic/GEditorJS/block"
	"github.com/davesavic/GEditorJS/converter"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func loadBlockData(path string, blockData any) error {
	d, err := os.ReadFile(filepath.Join("testdata", path))
	if err != nil {
		return err
	}

	return json.Unmarshal(d, blockData)
}

func TestHTMLBlockConverter_Convert(t *testing.T) {
	testCases := []struct {
		name      string
		blockType string
		filePath  string
		expected  string
	}{
		{
			name:      "converts header block to html",
			blockType: "header",
			filePath:  "header.json",
			expected:  "<h1>Header</h1>",
		},
		{
			name:      "converts paragraph block to html",
			blockType: "paragraph",
			filePath:  "paragraph.json",
			expected:  "<p>Paragraph</p>",
		},
		{
			name:      "converts raw block to html",
			blockType: "raw",
			filePath:  "raw.json",
			expected:  "Raw",
		},
		{
			name:      "converts checklist block to html",
			blockType: "checklist",
			filePath:  "checklist.json",
			expected:  `<ul><li><input type="checkbox" disabled>Unchecked</li><li><input type="checkbox" disabled checked>Checked</li></ul>`,
		},
		{
			name:      "converts list block to html",
			blockType: "list",
			filePath:  "list.json",
			expected:  `<ul style="list-style-type: ordered"><li>Ordered</li><li>List</li></ul>`,
		},
		{
			name:      "converts quote block to html",
			blockType: "quote",
			filePath:  "quote.json",
			expected:  `<blockquote style="text-align: left"><p>Quote</p><footer>Caption</footer></blockquote>`,
		},
		{
			name:      "converts delimiter block to html",
			blockType: "delimiter",
			filePath:  "delimiter.json",
			expected:  `<hr>`,
		},
		{
			name:      "converts code block to html",
			blockType: "code",
			filePath:  "code.json",
			expected:  `<pre><code>Code</code></pre>`,
		},
		{
			name: "converts table block to html",
			blockType: "table",
			filePath: "table.json",
			expected: `<table><thead><tr><th>Header 1</th><th>Header 2</th></tr></thead><tbody><tr><td>Cell 1</td><td>Cell 2</td></tr></tbody></table>`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := converter.GetBlockConverter(tc.blockType)
			assert.NotNil(t, c, "Converter should not be nil")

			var b block.Block
			b.Type = tc.blockType
			err := loadBlockData(tc.filePath, &b.Data)
			assert.NoError(t, err, "Failed to load block data")

			result, err := c.Convert(b)
			assert.NoError(t, err, "Conversion should not error out")
			assert.Equal(t, tc.expected, result, "Conversion result mismatch")
		})
	}
}

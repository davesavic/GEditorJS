package GEditorJS_test

import (
	"github.com/davesavic/GEditorJS"
	"github.com/davesavic/GEditorJS/converter"
	"os"
	"testing"
)

func TestGEditorJS_Parse(t *testing.T) {
	sb, err := os.ReadFile("testdata/supported_blocks.json")
	if err != nil {
		t.Fatal(err)
	}

	ge := GEditorJS.New()
	ge.SetRawContent(string(sb))

	o, err := ge.Parse(converter.HTMLBlockConverter{})
	if err != nil {
		t.Fatal(err)
	}

	if o == "" {
		t.Fatal("output is empty")
	}
}

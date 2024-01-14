package main

import (
	"fmt"
	"github.com/davesavic/GEditorJS"
	"github.com/davesavic/GEditorJS/converter"
	"os"
)

func main() {
	ge := GEditorJS.New()

	output, err := os.ReadFile("converter/testdata/supported_blocks.json")
	if err != nil {
		panic(err)
	}

	ge.SetRawContent(string(output))

	o, err := ge.Parse(converter.HTMLBlockConverter{})
	if err != nil {
		panic(err)
	}

	fmt.Println(o)
}

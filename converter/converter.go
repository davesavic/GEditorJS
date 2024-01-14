package converter

import (
	"github.com/davesavic/GEditorJS/block"
)

type BlockConverter interface {
	Convert(b block.Block) (string, error)
}

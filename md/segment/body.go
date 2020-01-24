package segment

import (
	"github.com/yuin/goldmark/ast"
	"tix/md"
)

type BodySegmentParser interface {
	Parse(state *md.State, node ast.Node)
}
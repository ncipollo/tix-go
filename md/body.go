package md

import (
	"github.com/yuin/goldmark/ast"
)

type BodySegmentParser interface {
	Parse(state *State, node ast.Node) error
}
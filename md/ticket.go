package md

import (
	"github.com/yuin/goldmark/ast"
)

type TicketParser struct {
	bodySegmentParsers []BodySegmentParser
}

func (p TicketParser) Parse(state *State, node ast.Node)  {

}

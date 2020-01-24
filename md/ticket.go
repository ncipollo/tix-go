package md

import (
	"github.com/yuin/goldmark/ast"
	"tix/md/segment"
)

type TicketParser struct {
	bodySegmentParsers []segment.BodySegmentParser
}

func (p TicketParser) Parse(state *State, node ast.Node)  {

}

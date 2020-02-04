package md

import (
	"github.com/yuin/goldmark/ast"
	"tix/ticket/body"
)

type ThematicBreakSegmentParser struct{}

func NewThematicBreakSegmentParser() *ThematicBreakSegmentParser {
	return &ThematicBreakSegmentParser{}
}

func (b ThematicBreakSegmentParser) Parse(state *State, node ast.Node) error {
	currentTicket := state.CurrentTicket()

	currentTicket.AddBodyLineBreak()
	currentTicket.AddBodySegment(body.NewThematicBreakSegment())
	currentTicket.AddBodyLineBreak()

	return nil
}

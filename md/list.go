package md

import (
	"github.com/yuin/goldmark/ast"
	"tix/ticket/body"
)

type ListSegmentParser struct {
}

func NewListSegmentParser() *ListSegmentParser {
	return &ListSegmentParser{}
}

func (l ListSegmentParser) Parse(state *State, node ast.Node) error {
	currentTicket := state.CurrentTicket()
	listItemNode := node.(*ast.List)

	lineBreak := body.NewLineBreakSegment()
	currentTicket.AddBodySegment(lineBreak)
	err := ParseBodyChildren(state, listItemNode)

	return err
}
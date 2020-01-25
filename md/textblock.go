package md

import (
	"github.com/yuin/goldmark/ast"
	"tix/ticket/body"
)

type TextBlockSegmentParser struct {
}

func NewTextBlockSegmentParser() *TextBlockSegmentParser {
	return &TextBlockSegmentParser{}
}

func (t TextBlockSegmentParser) Parse(state *State, node ast.Node) error {
	currentTicket := state.CurrentTicket()
	paragraph := node.(*ast.TextBlock)

	block := body.NewTextBlockSegment()
	currentTicket.AddBodySegment(block)
	err := ParseBodyChildren(state, paragraph)

	return err
}
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
	textBlock := node.(*ast.TextBlock)

	block := body.NewTextBlockSegment()
	currentTicket.AddBodySegment(block)
	err := ParseBodyChildren(state, textBlock)
	if textBlock.HasChildren() {
		currentTicket.AddBodyLineBreak()
	}
	return err
}
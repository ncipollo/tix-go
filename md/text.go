package md

import (
	"github.com/yuin/goldmark/ast"
	"tix/ticket/body"
)

type TextSegmentParser struct {
	
}

func NewTextSegmentParser() *TextSegmentParser {
	return &TextSegmentParser{}
}

func (t TextSegmentParser) Parse(state *State, node ast.Node) error {
	currentTicket := state.CurrentTicket()
	text := node.(*ast.Text)
	data := text.Segment.Value(state.SourceData)

	bodyText := body.NewTextSegment(string(data))
	currentTicket.AddBodySegment(bodyText)

	return nil
}

package md

import (
	"errors"
	"github.com/yuin/goldmark/ast"
	"tix/ticket/body"
)

type CodeSpanSegmentParser struct {
}

func NewCodeSpanSegmentParser() *CodeSpanSegmentParser {
	return &CodeSpanSegmentParser{}
}

func (c CodeSpanSegmentParser) Parse(state *State, node ast.Node) error {
	currentTicket := state.CurrentTicket()
	codeSpanNode := node.(*ast.CodeSpan)
	if codeSpanNode.ChildCount() == 0 || codeSpanNode.FirstChild().Kind() != ast.KindText {
		return errors.New("emphasis must have text contents")
	}

	textNode := codeSpanNode.FirstChild().(*ast.Text)
	data := textNode.Segment.Value(state.SourceData)

	codeSpan := body.NewCodeSpanSegment(string(data))
	currentTicket.AddBodySegment(codeSpan)

	return nil
}

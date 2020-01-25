package md

import (
	"github.com/yuin/goldmark/ast"
	"tix/ticket/body"
)

type ParagraphSegmentParser struct {
}

func NewParagraphSegmentParser() *ParagraphSegmentParser {
	return &ParagraphSegmentParser{}
}

func (t ParagraphSegmentParser) Parse(state *State, node ast.Node) error {
	currentTicket := state.CurrentTicket()
	paragraph := node.(*ast.Paragraph)

	lineBreak := body.NewLineBreakSegment()
	currentTicket.AddBodySegment(lineBreak)
	err := ParseBodyChildren(state, paragraph)

	return err
}
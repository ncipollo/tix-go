package md

import (
	"github.com/yuin/goldmark/ast"
)

type ParagraphSegmentParser struct {
}

func NewParagraphSegmentParser() *ParagraphSegmentParser {
	return &ParagraphSegmentParser{}
}

func (t ParagraphSegmentParser) Parse(state *State, node ast.Node) error {
	currentTicket := state.CurrentTicket()
	paragraph := node.(*ast.Paragraph)

	if paragraph.HasBlankPreviousLines() {
		currentTicket.AddBodyLineBreak()
	}
	err := ParseBodyChildren(state, paragraph)
	currentTicket.AddBodyLineBreak()

	return err
}
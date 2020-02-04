package md

import (
	"github.com/yuin/goldmark/ast"
	"tix/ticket/body"
)

type BlockQuoteSegmentParser struct{}

func NewBlockQuoteSegmentParser() *BlockQuoteSegmentParser {
	return &BlockQuoteSegmentParser{}
}

func (b BlockQuoteSegmentParser) Parse(state *State, node ast.Node) error {
	currentTicket := state.CurrentTicket()
	block := node.(*ast.Blockquote)

	currentTicket.AddBodySegment(body.NewBlockQuoteSegment())
	err := ParseBodyChildren(state, block)

	return err
}

package md

import (
	"github.com/yuin/goldmark/ast"
	"tix/ticket"
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
	text.SoftLineBreak()
	data := text.Segment.Value(state.SourceData)

	bodyText := body.NewTextSegment(string(data))
	currentTicket.AddBodySegment(bodyText)

	t.AddLineBreaks(currentTicket, text)

	return nil
}

func (t TextSegmentParser) AddLineBreaks(currentTicket *ticket.Ticket, text *ast.Text) {
	if text.SoftLineBreak() || text.HardLineBreak() {
		currentTicket.AddBodyLineBreak()
	}
}
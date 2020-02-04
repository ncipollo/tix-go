package md

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"tix/ticket/body"
)

func TestBlockQuoteSegmentParser_Parse(t *testing.T) {
	text := `
> line1
> line2
`
	parser := NewBlockQuoteSegmentParser()
	state, rootNode := setupTextParser(text)
	state.StartTicket()
	node := rootNode.FirstChild()

	err := parser.Parse(state, node)

	expectedBody := []body.Segment{
		body.NewBlockQuoteSegment(),
		body.NewTextSegment("line1"),
		body.NewLineBreakSegment(),
		body.NewTextSegment("line2"),
		body.NewLineBreakSegment(),
	}
	ticketBody := state.CurrentTicket().Body
	assert.NoError(t, err)
	assert.Equal(t, expectedBody, ticketBody)
}
package md

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"tix/ticket/body"
)

func TestTextBlockSegmentParser_Parse(t *testing.T) {
	text := `
- Item 1
`
	parser := NewListItemSegmentParser()
	state, rootNode := setupTextParser(text)
	state.StartTicket()
	state.ListState.StartBulletList("-")
	node := rootNode.FirstChild()

	err := parser.Parse(state, node.FirstChild())

	expectedBody := []body.Segment{
		body.NewBulletListItemSegment(1, "-"),
		body.NewLineBreakSegment(),
		body.NewTextBlockSegment(),
		body.NewTextSegment("Item 1"),
	}
	ticketBody := state.CurrentTicket().Body
	assert.NoError(t, err)
	assert.Equal(t, expectedBody, ticketBody)
}

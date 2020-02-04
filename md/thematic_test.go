package md

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"tix/ticket/body"
)

func TestThematicBreakSegmentParser_Parse(t *testing.T) {
	text := `**********`
	parser := NewThematicBreakSegmentParser()
	state, rootNode := setupTextParser(text)
	state.StartTicket()
	node := rootNode.FirstChild()

	err := parser.Parse(state, node)

	expectedBody := []body.Segment{
		body.NewLineBreakSegment(),
		body.NewThematicBreakSegment(),
		body.NewLineBreakSegment(),
	}
	ticketBody := state.CurrentTicket().Body
	assert.NoError(t, err)
	assert.Equal(t, expectedBody, ticketBody)
}
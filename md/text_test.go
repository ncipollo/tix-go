package md

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"tix/ticket/body"
)

func TestTextSegmentParser_Parse_NoLineBreak(t *testing.T) {
	text := `
text
`
	parser := NewTextSegmentParser()
	state, rootNode := setupTextParser(text)
	state.StartTicket()
	node := rootNode.FirstChild()

	err := parser.Parse(state, node.FirstChild())

	expectedBody := []body.Segment{
		body.NewTextSegment("text"),
	}
	ticketBody := state.CurrentTicket().Body
	assert.NoError(t, err)
	assert.Equal(t, expectedBody, ticketBody)
}

func TestTextSegmentParser_Parse_WithLineBreak(t *testing.T) {
	text := `
text
text
`
	parser := NewTextSegmentParser()
	state, rootNode := setupTextParser(text)
	state.StartTicket()
	node := rootNode.FirstChild()

	err := parser.Parse(state, node.FirstChild())

	expectedBody := []body.Segment{
		body.NewTextSegment("text"),
		body.NewLineBreakSegment(),
	}
	ticketBody := state.CurrentTicket().Body
	assert.NoError(t, err)
	assert.Equal(t, expectedBody, ticketBody)
}
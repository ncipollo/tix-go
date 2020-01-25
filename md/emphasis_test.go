package md

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"tix/ticket/body"
)

func TestEmphasisSegmentParser_Emphasis(t *testing.T) {
	text := `
*text*
`
	parser := NewEmphasisSegmentParser()
	state, rootNode := setupTextParser(text)
	state.StartTicket()
	node := rootNode.FirstChild()

	err := parser.Parse(state, node.FirstChild())

	expectedBody := []body.Segment{
		body.NewEmphasisSegment("text"),
	}
	ticketBody := state.CurrentTicket().Body
	assert.NoError(t, err)
	assert.Equal(t, expectedBody, ticketBody)
}

func  TestEmphasisParser_Parse_StrongEmphasis(t *testing.T) {
	text := `
**text**
`
	parser := NewEmphasisSegmentParser()
	state, rootNode := setupTextParser(text)
	state.StartTicket()
	node := rootNode.FirstChild()

	err := parser.Parse(state, node.FirstChild())

	expectedBody := []body.Segment{
		body.NewStrongEmphasisSegment("text"),
	}
	ticketBody := state.CurrentTicket().Body
	assert.NoError(t, err)
	assert.Equal(t, expectedBody, ticketBody)
}
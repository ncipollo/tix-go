package md

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"tix/ticket/body"
)

func TestLinkSegmentParser_Parse(t *testing.T) {
	text := "[link](http://api.example.com)"
	parser := NewLinkSegmentParser()
	state, rootNode := setupTextParser(text)
	state.StartTicket()
	node := rootNode.FirstChild()

	err := parser.Parse(state, node.FirstChild())

	expectedBody := []body.Segment{
		body.NewLinkSegment("link", "http://api.example.com"),
	}
	ticketBody := state.CurrentTicket().Body
	assert.NoError(t, err)
	assert.Equal(t, expectedBody, ticketBody)
}
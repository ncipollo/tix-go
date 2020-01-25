package md

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"tix/ticket/body"
)

func TestListItemSegmentParser_Parse_BulletList(t *testing.T) {
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

func TestListItemSegmentParser_Parse_ErrorWhenNoList(t *testing.T) {
	text := `
- Item 1
`
	parser := NewListItemSegmentParser()
	state, rootNode := setupTextParser(text)
	state.StartTicket()
	node := rootNode.FirstChild()

	err := parser.Parse(state, node.FirstChild())

	assert.Error(t, err)
}

func TestListItemSegmentParser_Parse_OrderedList(t *testing.T) {
	text := `
1. First
`
	parser := NewListItemSegmentParser()
	state, rootNode := setupTextParser(text)
	state.StartTicket()
	state.ListState.StartOrderedList(1)
	node := rootNode.FirstChild()

	err := parser.Parse(state, node.FirstChild())

	expectedBody := []body.Segment{
		body.NewOrderedListItemSegment(1, 1),
		body.NewLineBreakSegment(),
		body.NewTextBlockSegment(),
		body.NewTextSegment("First"),
	}
	ticketBody := state.CurrentTicket().Body
	assert.NoError(t, err)
	assert.Equal(t, expectedBody, ticketBody)
	assert.Equal(t, 2, state.ListState.CurrentList().CurrentNumber)
}

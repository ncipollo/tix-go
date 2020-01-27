package md

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"tix/ticket/body"
)

func TestParagraphSegmentParser_Parse(t *testing.T) {
	text := `
text1
*emphasis*
text2
`
	parser := NewParagraphSegmentParser()
	state, rootNode := setupTextParser(text)
	state.StartTicket()
	node := rootNode.FirstChild()

	err := parser.Parse(state, node)

	expectedBody := []body.Segment{
		body.NewTextSegment("text1"),
		body.NewLineBreakSegment(),
		body.NewEmphasisSegment("emphasis"),
		body.NewTextSegment(""),
		body.NewLineBreakSegment(),
		body.NewTextSegment("text2"),
		body.NewLineBreakSegment(),
	}
	ticketBody := state.CurrentTicket().Body
	assert.NoError(t, err)
	assert.Equal(t, expectedBody, ticketBody)
}
package md

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"tix/ticket/body"
)

func TestParagraphSegmentParser_Parse_WithBlankPreviousLine(t *testing.T) {
	text :=
`
text1
*emphasis* text2 *emphasis*
text3
`
	parser := NewParagraphSegmentParser()
	state, rootNode := setupTextParser(text)
	state.StartTicket()
	node := rootNode.FirstChild()

	err := parser.Parse(state, node)

	expectedBody := []body.Segment{
		body.NewLineBreakSegment(),
		body.NewTextSegment("text1"),
		body.NewLineBreakSegment(),
		body.NewEmphasisSegment("emphasis"),
		body.NewTextSegment(" text2 "),
		body.NewEmphasisSegment("emphasis"),
		body.NewTextSegment(""),
		body.NewLineBreakSegment(),
		body.NewTextSegment("text3"),
		body.NewLineBreakSegment(),
	}
	ticketBody := state.CurrentTicket().Body
	assert.NoError(t, err)
	assert.Equal(t, expectedBody, ticketBody)
}

func TestParagraphSegmentParser_ParseWithoutBlankPreviousLine(t *testing.T) {
	text := `# Heading 1
text1
*emphasis* text2 *emphasis*
text3
`
	parser := NewParagraphSegmentParser()
	state, rootNode := setupTextParser(text)
	state.StartTicket()
	node := rootNode.FirstChild().NextSibling()

	err := parser.Parse(state, node)

	expectedBody := []body.Segment{
		body.NewTextSegment("text1"),
		body.NewLineBreakSegment(),
		body.NewEmphasisSegment("emphasis"),
		body.NewTextSegment(" text2 "),
		body.NewEmphasisSegment("emphasis"),
		body.NewTextSegment(""),
		body.NewLineBreakSegment(),
		body.NewTextSegment("text3"),
		body.NewLineBreakSegment(),
	}
	ticketBody := state.CurrentTicket().Body
	assert.NoError(t, err)
	assert.Equal(t, expectedBody, ticketBody)
}
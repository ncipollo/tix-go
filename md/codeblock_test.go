package md

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"tix/ticket/body"
)

func TestCodeBlockSegmentParser_Parse_Normal(t *testing.T) {
	text := `
	code1
	code2
`
	parser := NewCodeBlockSegmentParser(false)
	state, rootNode := setupTextParser(text)
	state.StartTicket()
	node := rootNode.FirstChild()

	err := parser.Parse(state, node)

	expectedBody := []body.Segment{
		body.NewCodeBlockSegment("code1\ncode2\n", ""),
		body.NewLineBreakSegment(),
	}
	ticketBody := state.CurrentTicket().Body
	assert.NoError(t, err)
	assert.Equal(t, expectedBody, ticketBody)
}

func TestCodeBlockSegmentParser_Parse_Fenced(t *testing.T) {
	text := "```go\n" +
		"code1\n" +
		"code2\n" +
		"```"
	parser := NewCodeBlockSegmentParser(true)
	state, rootNode := setupTextParser(text)
	state.StartTicket()
	node := rootNode.FirstChild()

	err := parser.Parse(state, node)

	expectedBody := []body.Segment{
		body.NewCodeBlockSegment("code1\ncode2\n", "go"),
		body.NewLineBreakSegment(),
	}
	ticketBody := state.CurrentTicket().Body
	assert.NoError(t, err)
	assert.Equal(t, expectedBody, ticketBody)
}

func TestCodeBlockSegmentParser_Parse_TicketMetadata(t *testing.T) {
	text := "```tix\n" +
		"meta\n" +
		"```"
	parser := NewCodeBlockSegmentParser(true)
	state, rootNode := setupTextParser(text)
	state.StartTicket()
	node := rootNode.FirstChild()

	err := parser.Parse(state, node)

	expectedMetadata := "meta\n"
	ticketBody := state.CurrentTicket().Body
	ticketMetadata := state.CurrentTicket().Metadata
	assert.NoError(t, err)
	assert.Empty(t, ticketBody)
	assert.Equal(t, expectedMetadata, ticketMetadata)
}

func TestCodeBlockSegmentParser_Parse_No_Language(t *testing.T) {
	text := "```\n" +
		"code1\n" +
		"code2\n" +
		"```"
	parser := NewCodeBlockSegmentParser(true)
	state, rootNode := setupTextParser(text)
	state.StartTicket()
	node := rootNode.FirstChild()

	err := parser.Parse(state, node)

	expectedBody := []body.Segment{
		body.NewCodeBlockSegment("code1\ncode2\n", ""),
		body.NewLineBreakSegment(),
	}
	ticketBody := state.CurrentTicket().Body
	assert.NoError(t, err)
	assert.Equal(t, expectedBody, ticketBody)
}
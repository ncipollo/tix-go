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

func TestCodeBlockSegmentParser_Parse_Fields_Default(t *testing.T) {
	text := "```tix\n" +
		"foo: bar\n" +
		"```"
	parser := NewCodeBlockSegmentParser(true)
	state, rootNode := setupTextParser(text)
	state.StartTicket()
	node := rootNode.FirstChild()

	err := parser.Parse(state, node)

	expected := map[string]interface{}{"foo" : "bar"}
	ticketBody := state.CurrentTicket().Body
	ticketFields := state.CurrentTicket().Fields("jira")
	assert.NoError(t, err)
	assert.Empty(t, ticketBody)
	assert.Equal(t, expected, ticketFields)
}

func TestCodeBlockSegmentParser_Parse_Fields_Error(t *testing.T) {
	text := "```github\n" +
		"\t\t" +
		"```"
	parser := NewCodeBlockSegmentParser(true)
	state, rootNode := setupTextParser(text)
	state.StartTicket()
	node := rootNode.FirstChild()

	err := parser.Parse(state, node)

	assert.Error(t, err)
}

func TestCodeBlockSegmentParser_Parse_Fields_Github(t *testing.T) {
	text := "```github\n" +
		"foo: bar\n" +
		"```"
	parser := NewCodeBlockSegmentParser(true)
	state, rootNode := setupTextParser(text)
	state.StartTicket()
	node := rootNode.FirstChild()

	err := parser.Parse(state, node)

	expected := map[string]interface{}{"foo" : "bar"}
	ticketBody := state.CurrentTicket().Body
	ticketFields := state.CurrentTicket().Fields("github")
	assert.NoError(t, err)
	assert.Empty(t, ticketBody)
	assert.Equal(t, expected, ticketFields)
}

func TestCodeBlockSegmentParser_Parse_Fields_Jira(t *testing.T) {
	text := "```jira\n" +
		"foo: bar\n" +
		"```"
	parser := NewCodeBlockSegmentParser(true)
	state, rootNode := setupTextParser(text)
	state.StartTicket()
	node := rootNode.FirstChild()

	err := parser.Parse(state, node)

	expected := map[string]interface{}{"foo" : "bar"}
	ticketBody := state.CurrentTicket().Body
	ticketFields := state.CurrentTicket().Fields("jira")
	assert.NoError(t, err)
	assert.Empty(t, ticketBody)
	assert.Equal(t, expected, ticketFields)
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
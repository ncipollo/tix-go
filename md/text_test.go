package md

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"tix/ticket/body"
)

func TestTextSegmentParser_Parse(t *testing.T) {
	text := `
text
`
	parser := NewTextSegmentParser()
	state, rootNode := setupTextParser(text)
	state.StartTicket()
	node := rootNode.FirstChild()

	err := parser.Parse(state, node.FirstChild())

	bodySegment := state.CurrentTicket().Body[0]
	assert.NoError(t, err)
	assert.IsType(t, &body.TextSegment{}, bodySegment)
	assert.Equal(t, "text", bodySegment.Value())
}
package md

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"tix/ticket/body"
)

func TestListSegmentParser_Parse_BulletList(t *testing.T) {
	text := `
- Root
	- Sub 1
		+ Deep
	- Sub 2
`
	parser := NewListSegmentParser()
	state, rootNode := setupTextParser(text)
	state.StartTicket()
	node := rootNode.FirstChild()

	err := parser.Parse(state, node)

	var expectedBody []body.Segment
	expectedBody = appendBulletList(expectedBody, 1, "-")
	expectedBody = appendBulletLineSegments(expectedBody, 1, "-", "Root")
	expectedBody = appendBulletList(expectedBody, 2, "-")
	expectedBody = appendBulletLineSegments(expectedBody, 2, "-", "Sub 1")
	expectedBody = appendBulletList(expectedBody, 3, "+")
	expectedBody = appendBulletLineSegments(expectedBody, 3, "+", "Deep")
	expectedBody = appendBulletLineSegments(expectedBody, 2, "-", "Sub 2")
	ticketBody := state.CurrentTicket().Body
	assert.NoError(t, err)
	assert.Equal(t, expectedBody, ticketBody)
}

func appendBulletList(expectedBody []body.Segment, level int, marker string) []body.Segment {
	list := body.NewBulletListStartSegment(level, marker)
	return append(expectedBody, list)
}

func appendBulletLineSegments(expectedBody []body.Segment, level int, marker string, text string) []body.Segment {
	segments := []body.Segment{
		body.NewBulletListItemSegment(level, marker),
		body.NewTextBlockSegment(),
		body.NewTextSegment(text),
		body.NewLineBreakSegment(),
	}
	return append(expectedBody, segments...)
}

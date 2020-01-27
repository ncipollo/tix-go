package md

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"tix/ticket/body"
)

func TestListSegmentParser_Parse_MixedList(t *testing.T) {
	text := `
- Bullet 1
	1. Sub 1
	2. Sub 2
- Bullet 2
`
	parser := NewListSegmentParser()
	state, rootNode := setupTextParser(text)
	state.StartTicket()
	node := rootNode.FirstChild()

	err := parser.Parse(state, node)

	var expectedBody []body.Segment
	expectedBody = appendBulletList(expectedBody, 1, "-")
	expectedBody = appendBulletLineSegments(expectedBody, 1, "-", "Bullet 1")
	expectedBody = appendOrderedList(expectedBody, 2, 1)
	expectedBody = appendOrderedLineSegments(expectedBody, 2, 1, "Sub 1")
	expectedBody = appendOrderedLineSegments(expectedBody, 2, 2, "Sub 2")
	expectedBody = appendBulletLineSegments(expectedBody, 1, "-", "Bullet 2")
	ticketBody := state.CurrentTicket().Body
	assert.NoError(t, err)
	assert.Equal(t, expectedBody, ticketBody)
}

func TestListSegmentParser_Parse_NestedBulletList(t *testing.T) {
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

func TestListSegmentParser_Parse_NestedOrderedList(t *testing.T) {
	text := `
1. Root
	1. Sub 1
		1. Deep
	2. Sub 2
`
	parser := NewListSegmentParser()
	state, rootNode := setupTextParser(text)
	state.StartTicket()
	node := rootNode.FirstChild()

	err := parser.Parse(state, node)

	var expectedBody []body.Segment
	expectedBody = appendOrderedList(expectedBody, 1, 1)
	expectedBody = appendOrderedLineSegments(expectedBody, 1, 1, "Root")
	expectedBody = appendOrderedList(expectedBody, 2, 1)
	expectedBody = appendOrderedLineSegments(expectedBody, 2, 1, "Sub 1")
	expectedBody = appendOrderedList(expectedBody, 3, 1)
	expectedBody = appendOrderedLineSegments(expectedBody, 3, 1, "Deep")
	expectedBody = appendOrderedLineSegments(expectedBody, 2, 2, "Sub 2")
	ticketBody := state.CurrentTicket().Body
	assert.NoError(t, err)
	assert.Equal(t, expectedBody, ticketBody)
}

func TestListSegmentParser_Parse_WithInlineElements(t *testing.T) {
	text := `
- **Strong**
- [link](api.example.com)
`
	parser := NewListSegmentParser()
	state, rootNode := setupTextParser(text)
	state.StartTicket()
	node := rootNode.FirstChild()

	err := parser.Parse(state, node)

	var expectedBody []body.Segment
	expectedBody = appendBulletList(expectedBody, 1, "-")
	expectedBody = appendBulletEmphasisSegments(expectedBody, 1, "-", "Strong")
	expectedBody = appendBulletLinkSegments(expectedBody, 1, "-", "link", "api.example.com")

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

func appendBulletEmphasisSegments(expectedBody []body.Segment, level int, marker string, text string) []body.Segment {
	segments := []body.Segment{
		body.NewBulletListItemSegment(level, marker),
		body.NewTextBlockSegment(),
		body.NewStrongEmphasisSegment(text),
		body.NewLineBreakSegment(),
	}
	return append(expectedBody, segments...)
}

func appendBulletLinkSegments(expectedBody []body.Segment, level int, marker string, text string, url string) []body.Segment {
	segments := []body.Segment{
		body.NewBulletListItemSegment(level, marker),
		body.NewTextBlockSegment(),
		body.NewLinkSegment(text, url),
		body.NewLineBreakSegment(),
	}
	return append(expectedBody, segments...)
}

func appendOrderedList(expectedBody []body.Segment, level int, start int) []body.Segment {
	list := body.NewOrderedListStartSegment(level, start)
	return append(expectedBody, list)
}

func appendOrderedLineSegments(expectedBody []body.Segment, level int, number int, text string) []body.Segment {
	segments := []body.Segment{
		body.NewOrderedListItemSegment(level, number),
		body.NewTextBlockSegment(),
		body.NewTextSegment(text),
		body.NewLineBreakSegment(),
	}
	return append(expectedBody, segments...)
}

package render

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"tix/ticket/body"
)

func TestJiraBodyRenderer_RenderSegment_BulletListItem_LevelOne(t *testing.T) {
	segment := body.NewBulletListItemSegment(1, "-")
	renderer := NewJiraBodyRenderer()

	text := renderer.RenderSegment(segment)

	expected := "- "
	assert.Equal(t, expected, text)
}

func TestJiraBodyRenderer_RenderSegment_BulletListItem_LevelTwo(t *testing.T) {
	segment := body.NewBulletListItemSegment(2, "-")
	renderer := NewJiraBodyRenderer()

	text := renderer.RenderSegment(segment)

	expected := "-- "
	assert.Equal(t, expected, text)
}

func TestJiraBodyRenderer_RenderSegment_BlockQuote(t *testing.T) {
	segment := body.NewBlockQuoteSegment()
	renderer := NewJiraBodyRenderer()

	text := renderer.RenderSegment(segment)

	expected := "{quote}"
	assert.Equal(t, expected, text)
}

func TestJiraBodyRenderer_RenderSegment_CodeBlockSegment_NoLanguage(t *testing.T) {
	segment := body.NewCodeBlockSegment("println()\n", "")
	renderer := NewJiraBodyRenderer()

	text := renderer.RenderSegment(segment)

	expected := `{code}
println()
{code}`
	assert.Equal(t, expected, text)
}

func TestJiraBodyRenderer_RenderSegment_CodeBlockSegment_WithLanguage(t *testing.T) {
	segment := body.NewCodeBlockSegment("println()\n", "go")
	renderer := NewJiraBodyRenderer()

	text := renderer.RenderSegment(segment)

	expected := `{code:go}
println()
{code}`
	assert.Equal(t, expected, text)
}

func TestJiraBodyRenderer_RenderSegment_CodeSpan_NoSibling_NoSuffix(t *testing.T) {
	segment := body.NewCodeSpanSegment("code")
	renderer := NewJiraBodyRenderer()

	text := renderer.RenderSegment(segment)

	expected := "{{code}}"
	assert.Equal(t, expected, text)
}

func TestJiraBodyRenderer_RenderSegment_CodeSpan_WithLineBreakSibling_NoSuffix(t *testing.T) {
	segment := body.NewCodeSpanSegment("code")
	segment.SetNext(body.NewLineBreakSegment())
	renderer := NewJiraBodyRenderer()

	text := renderer.RenderSegment(segment)

	expected := "{{code}}"
	assert.Equal(t, expected, text)
}

func TestJiraBodyRenderer_RenderSegment_CodeSpan_WithTextSibling_NoSuffix(t *testing.T) {
	segment := body.NewCodeSpanSegment("code")
	segment.SetNext(body.NewTextSegment(" more text"))
	renderer := NewJiraBodyRenderer()

	text := renderer.RenderSegment(segment)

	expected := "{{code}}"
	assert.Equal(t, expected, text)
}

func TestJiraBodyRenderer_RenderSegment_CodeSpan_WithTextSibling_WithSuffix(t *testing.T) {
	segment := body.NewCodeSpanSegment("code")
	segment.SetNext(body.NewTextSegment("more text"))
	renderer := NewJiraBodyRenderer()

	text := renderer.RenderSegment(segment)

	expected := "{{code}} "
	assert.Equal(t, expected, text)
}

func TestJiraBodyRenderer_RenderSegment_Emphasis(t *testing.T) {
	segment := body.NewEmphasisSegment("words")
	renderer := NewJiraBodyRenderer()

	text := renderer.RenderSegment(segment)

	expected := "_words_"
	assert.Equal(t, expected, text)
}

func TestJiraBodyRenderer_RenderSegment_Link(t *testing.T) {
	segment := body.NewLinkSegment("text", "https://api.example.com")
	renderer := NewJiraBodyRenderer()

	text := renderer.RenderSegment(segment)

	expected := "[text|https://api.example.com]"
	assert.Equal(t, expected, text)
}

func TestJiraBodyRenderer_RenderSegment_List(t *testing.T) {
	segment := body.NewOrderedListStartSegment(0, 0)
	renderer := NewJiraBodyRenderer()

	text := renderer.RenderSegment(segment)

	expected := ""
	assert.Equal(t, expected, text)
}

func TestJiraBodyRenderer_RenderSegment_ListEnd_NoNewline(t *testing.T) {
	segment := body.NewBulletListEndSegment(2)
	renderer := NewJiraBodyRenderer()

	text := renderer.RenderSegment(segment)

	expected := ""
	assert.Equal(t, expected, text)
}

func TestJiraBodyRenderer_RenderSegment_ListEnd_WithNewline(t *testing.T) {
	segment := body.NewBulletListEndSegment(1)
	renderer := NewJiraBodyRenderer()

	text := renderer.RenderSegment(segment)

	expected := "\n"
	assert.Equal(t, expected, text)
}

func TestJiraBodyRenderer_RenderSegment_LineBreaks(t *testing.T) {
	segment := body.NewLineBreakSegment()
	renderer := NewJiraBodyRenderer()

	text := renderer.RenderSegment(segment)

	expected := "\n"
	assert.Equal(t, expected, text)
}

func TestJiraBodyRenderer_RenderSegment_OrderedListItem_LevelOne(t *testing.T) {
	segment := body.NewOrderedListItemSegment(1, 1)
	renderer := NewJiraBodyRenderer()

	text := renderer.RenderSegment(segment)

	expected := "# "
	assert.Equal(t, expected, text)
}

func TestJiraBodyRenderer_RenderSegment_OrderedListItem_LevelTwo(t *testing.T) {
	segment := body.NewOrderedListItemSegment(2, 1)
	renderer := NewJiraBodyRenderer()

	text := renderer.RenderSegment(segment)

	expected := "## "
	assert.Equal(t, expected, text)
}

func TestJiraBodyRenderer_RenderSegment_StrongEmphasis(t *testing.T) {
	segment := body.NewStrongEmphasisSegment("words")
	renderer := NewJiraBodyRenderer()

	text := renderer.RenderSegment(segment)

	expected := "*words*"
	assert.Equal(t, expected, text)
}

func TestJiraBodyRenderer_RenderSegment_TextBlock(t *testing.T) {
	segment := body.NewTextBlockSegment()
	renderer := NewJiraBodyRenderer()

	text := renderer.RenderSegment(segment)

	expected := ""
	assert.Equal(t, expected, text)
}

func TestJiraBodyRenderer_RenderSegment_Text(t *testing.T) {
	segment := body.NewTextSegment("text")
	renderer := NewJiraBodyRenderer()

	text := renderer.RenderSegment(segment)

	expected := "text"
	assert.Equal(t, expected, text)
}

func TestJiraBodyRenderer_RenderSegment_ThematicBreak(t *testing.T) {
	segment := body.NewThematicBreakSegment()
	renderer := NewJiraBodyRenderer()

	text := renderer.RenderSegment(segment)

	expected := "----"
	assert.Equal(t, expected, text)
}
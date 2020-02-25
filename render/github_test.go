package render

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"tix/ticket/body"
)

func TestGithubBodyRenderer_RenderSegment_BulletListItem_LevelOne(t *testing.T) {
	segment := body.NewBulletListItemSegment(1, "-")
	renderer := NewGithubBodyRenderer()

	text := renderer.RenderSegment(segment)

	expected := "- "
	assert.Equal(t, expected, text)
}

func TestGithubBodyRenderer_RenderSegment_BulletListItem_LevelTwo(t *testing.T) {
	segment := body.NewBulletListItemSegment(2, "-")
	renderer := NewGithubBodyRenderer()

	text := renderer.RenderSegment(segment)

	expected := "  - "
	assert.Equal(t, expected, text)
}

func TestGithubBodyRenderer_RenderSegment_BlockQuote(t *testing.T) {
	segment := body.NewBlockQuoteSegment()
	renderer := NewGithubBodyRenderer()

	text := renderer.RenderSegment(segment)

	expected := "> "
	assert.Equal(t, expected, text)
}

func TestGithubBodyRenderer_RenderSegment_CodeBlockSegment_NoLanguage(t *testing.T) {
	segment := body.NewCodeBlockSegment("println()\n", "")
	renderer := NewGithubBodyRenderer()

	text := renderer.RenderSegment(segment)

	expected := "```\n" +
		"println()\n" +
		"```"
	assert.Equal(t, expected, text)
}

func TestGithubBodyRenderer_RenderSegment_CodeBlockSegment_WithLanguage(t *testing.T) {
	segment := body.NewCodeBlockSegment("println()\n", "go")
	renderer := NewGithubBodyRenderer()

	text := renderer.RenderSegment(segment)

	expected := "```go\n" +
		"println()\n" +
		"```"
	assert.Equal(t, expected, text)
}

func TestGithubBodyRenderer_RenderSegment_CodeSpan(t *testing.T) {
	segment := body.NewCodeSpanSegment("code")
	renderer := NewGithubBodyRenderer()

	text := renderer.RenderSegment(segment)

	expected := "`code`"
	assert.Equal(t, expected, text)
}

func TestGithubBodyRenderer_RenderSegment_Emphasis(t *testing.T) {
	segment := body.NewEmphasisSegment("words")
	renderer := NewGithubBodyRenderer()

	text := renderer.RenderSegment(segment)

	expected := "*words*"
	assert.Equal(t, expected, text)
}

func TestGithubBodyRenderer_RenderSegment_Link(t *testing.T) {
	segment := body.NewLinkSegment("text", "https://api.example.com")
	renderer := NewGithubBodyRenderer()

	text := renderer.RenderSegment(segment)

	expected := "[text](https://api.example.com)"
	assert.Equal(t, expected, text)
}

func TestGithubBodyRenderer_RenderSegment_List(t *testing.T) {
	segment := body.NewOrderedListStartSegment(0, 0)
	renderer := NewGithubBodyRenderer()

	text := renderer.RenderSegment(segment)

	expected := ""
	assert.Equal(t, expected, text)
}

func TestGithubBodyRenderer_RenderSegment_ListEnd_NoNewline(t *testing.T) {
	segment := body.NewBulletListEndSegment(2)
	renderer := NewGithubBodyRenderer()

	text := renderer.RenderSegment(segment)

	expected := ""
	assert.Equal(t, expected, text)
}

func TestGithubBodyRenderer_RenderSegment_ListEnd_WithNewline(t *testing.T) {
	segment := body.NewBulletListEndSegment(1)
	renderer := NewGithubBodyRenderer()

	text := renderer.RenderSegment(segment)

	expected := "\n"
	assert.Equal(t, expected, text)
}

func TestGithubBodyRenderer_RenderSegment_LineBreaks(t *testing.T) {
	segment := body.NewLineBreakSegment()
	renderer := NewGithubBodyRenderer()

	text := renderer.RenderSegment(segment)

	expected := "\n"
	assert.Equal(t, expected, text)
}

func TestGithubBodyRenderer_RenderSegment_OrderedListItem_LevelOne(t *testing.T) {
	segment := body.NewOrderedListItemSegment(1, 1)
	renderer := NewGithubBodyRenderer()

	text := renderer.RenderSegment(segment)

	expected := "1. "
	assert.Equal(t, expected, text)
}

func TestGithubBodyRenderer_RenderSegment_OrderedListItem_LevelTwo(t *testing.T) {
	segment := body.NewOrderedListItemSegment(2, 1)
	renderer := NewGithubBodyRenderer()

	text := renderer.RenderSegment(segment)

	expected := "  1. "
	assert.Equal(t, expected, text)
}

func TestGithubBodyRenderer_RenderSegment_StrongEmphasis(t *testing.T) {
	segment := body.NewStrongEmphasisSegment("words")
	renderer := NewGithubBodyRenderer()

	text := renderer.RenderSegment(segment)

	expected := "**words**"
	assert.Equal(t, expected, text)
}

func TestGithubBodyRenderer_RenderSegment_TextBlock(t *testing.T) {
	segment := body.NewTextBlockSegment()
	renderer := NewGithubBodyRenderer()

	text := renderer.RenderSegment(segment)

	expected := ""
	assert.Equal(t, expected, text)
}

func TestGithubBodyRenderer_RenderSegment_Text(t *testing.T) {
	segment := body.NewTextSegment("text")
	renderer := NewGithubBodyRenderer()

	text := renderer.RenderSegment(segment)

	expected := "text"
	assert.Equal(t, expected, text)
}

func TestGithubBodyRenderer_RenderSegment_ThematicBreak(t *testing.T) {
	segment := body.NewThematicBreakSegment()
	renderer := NewGithubBodyRenderer()

	text := renderer.RenderSegment(segment)

	expected := "---"
	assert.Equal(t, expected, text)
}

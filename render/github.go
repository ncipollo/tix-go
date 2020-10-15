package render

import (
	"fmt"
	"strings"
	"tix/ticket/body"
)

type GithubBodyRenderer struct{}

func NewGithubBodyRenderer() *GithubBodyRenderer {
	return &GithubBodyRenderer{}
}

func (g GithubBodyRenderer) RenderSegment(bodySegment body.Segment) string {
	switch segment := bodySegment.(type) {
	case *body.BlockQuoteSegment:
		return g.renderBlockQuoteItem()
	case *body.BulletListItemSegment:
		return g.renderBulletListItem(segment)
	case *body.CodeBlockSegment:
		return g.renderCodeBlock(segment)
	case *body.CodeSpanSegment:
		return g.renderCodeSpan(segment)
	case *body.EmphasisSegment:
		return g.renderEmphasis(segment)
	case *body.LinkSegment:
		return g.renderLink(segment)
	case *body.ListEndSegment:
		return g.renderListEnd(segment)
	case *body.ListStartSegment:
		return segment.Value()
	case *body.LineBreakSegment:
		return segment.Value()
	case *body.OrderedListItemSegment:
		return g.renderOrderedListItem(segment)
	case *body.StrongEmphasisSegment:
		return g.renderStrongEmphasis(segment)
	case *body.TextBlockSegment:
		return segment.Value()
	case *body.TextSegment:
		return segment.Value()
	case *body.ThematicBreakSegment:
		return g.renderThematicBreak()
	default:
		return segment.Value()
	}
}

func (g GithubBodyRenderer) renderBlockQuoteItem() string {
	return "> "
}

func (g GithubBodyRenderer) renderBulletListItem(segment *body.BulletListItemSegment) string {
	var builder strings.Builder
	level := segment.Attributes().Level

	for ii := 0; ii < level - 1; ii++ {
		builder.WriteString("  ")
	}
	builder.WriteString(segment.Value())
	builder.WriteString(" ")

	return builder.String()
}

func (g GithubBodyRenderer) renderCodeBlock(segment *body.CodeBlockSegment) string {
	var builder strings.Builder
	lang := segment.Attributes().Language

	if len(lang) > 0 {
		marker := fmt.Sprintf("```%s\n", lang)
		builder.WriteString(marker)
	} else {
		builder.WriteString("```\n")
	}
	builder.WriteString(segment.Value())
	builder.WriteString("```")

	return builder.String()
}

func (g GithubBodyRenderer) renderCodeSpan(segment *body.CodeSpanSegment) string {
	return fmt.Sprintf("`%s`", segment.Value())
}

func (g GithubBodyRenderer) renderEmphasis(segment *body.EmphasisSegment) string {
	return fmt.Sprintf("*%s*", segment.Value())
}

func (g GithubBodyRenderer) renderLink(segment *body.LinkSegment) string {
	url := segment.Attributes().Url
	return fmt.Sprintf("[%s](%s)", segment.Value(), url)
}

func (g GithubBodyRenderer) renderOrderedListItem(segment *body.OrderedListItemSegment) string {
	var builder strings.Builder
	level := segment.Attributes().Level

	for ii := 0; ii < level - 1; ii++ {
		builder.WriteString("  ")
	}
	builder.WriteString("1.")
	builder.WriteString(" ")

	return builder.String()
}

func (g GithubBodyRenderer) renderListEnd(segment *body.ListEndSegment) string {
	if segment.Attributes().Level == 1 {
		return "\n"
	} else {
		return ""
	}
}

func (g GithubBodyRenderer) renderStrongEmphasis(segment *body.StrongEmphasisSegment) string {
	return fmt.Sprintf("**%s**", segment.Value())
}

func (g GithubBodyRenderer) renderThematicBreak() string {
	return "---"
}

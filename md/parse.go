package md

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
	"tix/ticket"
)

type Parser interface {
	Parse(source []byte) ([]*ticket.Ticket, error)
}

func NewParser(fieldState *FieldState) Parser {
	return &markdownParser{
		fieldState: fieldState,
		ticketParser: NewTicketParser(),
	}
}

type markdownParser struct {
	fieldState *FieldState
	ticketParser *TicketParser
}

func (m *markdownParser) Parse(source []byte) ([]*ticket.Ticket, error) {
	state, rootNode := setupParser(source, m.fieldState)

	err := m.ticketParser.Parse(state, rootNode)

	return state.RootTickets, err
}

func setupParser(source []byte, fieldState *FieldState) (*State, ast.Node) {
	state := newState(source, fieldState)

	reader := text.NewReader(source)
	markdownParser := createMarkdownParser()
	rootNode := markdownParser.Parse(reader)

	return state, rootNode
}

func createMarkdownParser() parser.Parser {
	return parser.NewParser(blockParsers(),
		inlineParsers(),
		parser.WithParagraphTransformers(parser.DefaultParagraphTransformers()...),
	)
}

func blockParsers() parser.Option  {
	parsers := []util.PrioritizedValue{
		util.Prioritized(parser.NewSetextHeadingParser(), 100),
		util.Prioritized(parser.NewThematicBreakParser(), 200),
		util.Prioritized(parser.NewListParser(), 300),
		util.Prioritized(parser.NewListItemParser(), 400),
		util.Prioritized(parser.NewCodeBlockParser(), 500),
		util.Prioritized(parser.NewATXHeadingParser(), 600),
		util.Prioritized(parser.NewFencedCodeBlockParser(), 700),
		util.Prioritized(parser.NewBlockquoteParser(), 800),
		util.Prioritized(parser.NewParagraphParser(), 900),
	}
	return parser.WithBlockParsers(parsers...)
}

func inlineParsers() parser.Option  {
	parsers := []util.PrioritizedValue{
		util.Prioritized(parser.NewCodeSpanParser(), 100),
		util.Prioritized(parser.NewLinkParser(), 200),
		util.Prioritized(parser.NewEmphasisParser(), 300),
	}
	return parser.WithInlineParsers(parsers...)
}

func setupTextParser(text string) (*State, ast.Node) {
	source := []byte(text)
	return setupParser(source, NewFieldState())
}

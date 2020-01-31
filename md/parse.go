package md

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
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
	parser := goldmark.DefaultParser()
	rootNode := parser.Parse(reader)

	return state, rootNode
}

func setupTextParser(text string) (*State, ast.Node) {
	source := []byte(text)
	return setupParser(source, NewFieldState())
}

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

func NewParser() Parser {
	return &markdownParser{
		ticketParser: NewTicketParser(),
	}
}

type markdownParser struct {
	ticketParser *TicketParser
}

func (m *markdownParser) Parse(source []byte) ([]*ticket.Ticket, error) {
	state, rootNode := setupParser(source)

	err := m.ticketParser.Parse(state, rootNode)

	return state.RootTickets, err
}

func setupParser(source []byte) (*State, ast.Node) {
	state := newState(source, nil)

	reader := text.NewReader(source)
	parser := goldmark.DefaultParser()
	rootNode := parser.Parse(reader)

	return state, rootNode
}

func setupTextParser(text string)  (*State, ast.Node) {
	source := []byte(text)
	return setupParser(source)
}

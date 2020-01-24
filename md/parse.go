package md

import (
	"github.com/yuin/goldmark"
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
	reader := text.NewReader(source)
	parser := goldmark.DefaultParser()
	rootNode := parser.Parse(reader)
	state := newState(source)

	err := m.ticketParser.Parse(state, rootNode)

	return state.RootTickets, err
}

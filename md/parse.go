package md

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
	"tix/ticket"
)

type Parser interface {
	Parse(source []byte) ([]ticket.Ticket, error)
}

func NewParser() Parser {
	return &markdownParser{}
}

type markdownParser struct{}

func (m markdownParser) Parse(source []byte) ([]ticket.Ticket, error) {
	reader := text.NewReader(source)
	parser := goldmark.DefaultParser()
	rootNode := parser.Parse(reader)
	state := newState()

	for node := rootNode.FirstChild(); node != nil; node = node.NextSibling() {
		kind := node.Kind()
		if shouldSkip(state, node) {
			continue
		}
		println(kind.String())
	}

	return make([]ticket.Ticket, 1), nil
}

func shouldSkip(state *State, node ast.Node) bool {
	if state.CurrentTicket() != nil {
		return false
	}

	if node.Kind() == ast.KindHeading {
		heading := node.(*ast.Heading)
		if heading.Level == 1 {
			return false
		} else {
			return true
		}
	} else {
		return true
	}
}
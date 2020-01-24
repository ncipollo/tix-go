package md

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
	"tix/ticket"
)

type Parser interface {
	Parse(source []byte) []ticket.Ticket
}

func NewParser() Parser {
	return &markdownParser{}
}

type markdownParser struct{}

func (m markdownParser) Parse(source []byte) []ticket.Ticket {
	reader := text.NewReader(source)
	parser := goldmark.DefaultParser()
	node := parser.Parse(reader)
	ast.Walk(node, func(visited ast.Node, entering bool) (ast.WalkStatus, error) {
		status := ast.WalkStatus(ast.WalkContinue)
		kind := visited.Kind()
		if entering {
			println(kind.String())
		}
		return status, nil
	})
	return make([]ticket.Ticket, 1)
}

package md

import (
	"errors"
	"fmt"
	"github.com/yuin/goldmark/ast"
)

type BodySegmentParser interface {
	Parse(state *State, node ast.Node) error
}

func BodyParserForKind(kind ast.NodeKind) (BodySegmentParser, error) {
	switch kind {
	case ast.KindText:
		return NewTextSegmentParser(), nil
	case ast.KindParagraph:
		return NewParagraphSegmentParser(), nil
	default:
		message := fmt.Sprintf("no body parser for markdown element type %v", kind)
		return nil, errors.New(message)
	}
}

func ParseBodyChildren(state *State, rootNode ast.Node) error {
	for node := rootNode.FirstChild(); node != nil; node = node.NextSibling() {
		kind := node.Kind()
		parser,err := BodyParserForKind(kind)
		if err != nil {
			return nil
		}
		err = parser.Parse(state, node)
		if err != nil {
			return nil
		}
	}
	return nil
}
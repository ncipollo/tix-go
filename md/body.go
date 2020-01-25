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
	case ast.KindCodeBlock:
		return NewCodeBlockSegmentParser(false), nil
	case ast.KindFencedCodeBlock:
		return NewCodeBlockSegmentParser(true), nil
	case ast.KindCodeSpan:
		return NewCodeSpanSegmentParser(), nil
	case ast.KindEmphasis:
		return NewEmphasisSegmentParser(), nil
	case ast.KindText:
		return NewTextSegmentParser(), nil
	case ast.KindLink:
		return NewLinkSegmentParser(), nil
	case ast.KindParagraph:
		return NewParagraphSegmentParser(), nil
	default:
		message := fmt.Sprintf("no body parser for markdown element type %v", kind)
		return nil, errors.New(message)
	}
}

func UnsupportedMarkdownKinds() []ast.NodeKind {
	return []ast.NodeKind {
		ast.KindBlockquote,
		ast.KindHTMLBlock,
		ast.KindList,
		ast.KindListItem,
		ast.KindTextBlock,
		ast.KindThematicBreak,
		ast.KindAutoLink,
		ast.KindImage,
		ast.KindRawHTML,
		ast.KindString,
	}
}

func ParseBodyChildren(state *State, rootNode ast.Node) error {
	for node := rootNode.FirstChild(); node != nil; node = node.NextSibling() {
		kind := node.Kind()
		parser,err := BodyParserForKind(kind)
		if err != nil {
			return err
		}
		err = parser.Parse(state, node)
		if err != nil {
			return err
		}
	}
	return nil
}
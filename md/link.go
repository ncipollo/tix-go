package md

import (
	"errors"
	"github.com/yuin/goldmark/ast"
	"tix/ticket/body"
)

type LinkSegmentParser struct {
}

func NewLinkSegmentParser() *LinkSegmentParser {
	return &LinkSegmentParser{}
}

func (c LinkSegmentParser) Parse(state *State, node ast.Node) error {
	currentTicket := state.CurrentTicket()
	linkNode := node.(*ast.Link)
	if linkNode.ChildCount() == 0 || linkNode.FirstChild().Kind() != ast.KindText {
		return errors.New("link must have text contents")
	}

	textNode := linkNode.FirstChild().(*ast.Text)
	data := textNode.Segment.Value(state.SourceData)

	codeSpan := body.NewLinkSegment(string(data), string(linkNode.Destination))
	currentTicket.AddBodySegment(codeSpan)

	return nil
}

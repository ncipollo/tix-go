package md

import (
	"errors"
	"github.com/yuin/goldmark/ast"
	"tix/ticket/body"
)

type EmphasisSegmentParser struct {
}

func NewEmphasisSegmentParser() *EmphasisSegmentParser {
	return &EmphasisSegmentParser{}
}

func (e EmphasisSegmentParser) Parse(state *State, node ast.Node) error {
	currentTicket := state.CurrentTicket()
	emphasisNode := node.(*ast.Emphasis)
	if emphasisNode.ChildCount() == 0 || emphasisNode.FirstChild().Kind() != ast.KindText {
		return errors.New("emphasis must have text contents")
	}

	textNode := emphasisNode.FirstChild().(*ast.Text)
	data := textNode.Segment.Value(state.SourceData)

	if emphasisNode.Level == 2 {
		emphasis := body.NewStrongEmphasisSegment(string(data))
		currentTicket.AddBodySegment(emphasis)
	} else {
		emphasis := body.NewEmphasisSegment(string(data))
		currentTicket.AddBodySegment(emphasis)
	}

	return nil
}

package md

import (
	"github.com/yuin/goldmark/ast"
	"tix/ticket/body"
)

type ListItemSegmentParser struct {
	isOrdered bool
	level     int
	marker    string
	number    int
}

func NewListItemSegmentParser(isOrdered bool, level int, marker string, number int) *ListItemSegmentParser {
	return &ListItemSegmentParser{isOrdered: isOrdered, level: level, marker: marker, number: number}
}

func (l ListItemSegmentParser) Parse(state *State, node ast.Node) error {
	currentTicket := state.CurrentTicket()
	listItemNode := node.(*ast.ListItem)

	if l.isOrdered {
		listItem := body.NewOrderedListItemSegment(l.level, l.number)
		currentTicket.AddBodySegment(listItem)
	} else {
		listItem := body.NewBulletListItemSegment(l.level, l.marker)
		currentTicket.AddBodySegment(listItem)
	}

	lineBreak := body.NewLineBreakSegment()
	currentTicket.AddBodySegment(lineBreak)
	err := ParseBodyChildren(state, listItemNode)

	return err
}

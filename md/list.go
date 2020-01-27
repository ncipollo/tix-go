package md

import (
	"github.com/yuin/goldmark/ast"
	"tix/ticket/body"
)

type ListSegmentParser struct{}

func NewListSegmentParser() *ListSegmentParser {
	return &ListSegmentParser{}
}

func (l ListSegmentParser) Parse(state *State, node ast.Node) error {
	listNode := node.(*ast.List)

	if listNode.IsOrdered() {
		return l.parseOrderedList(state, listNode)
	} else {
		return l.parseBulletList(state, listNode)
	}
}

func (l ListSegmentParser) parseOrderedList(state *State, node *ast.List) error {
	currentTicket := state.CurrentTicket()
	listState := state.ListState
	list := body.NewOrderedListStartSegment(listState.ListLevel(), node.Start)

	listState.StartOrderedList(node.Start)
	currentTicket.AddBodySegment(list)
	err := ParseBodyChildren(state, node)
	listState.CompleteList()

	return err
}

func (l ListSegmentParser) parseBulletList(state *State, node *ast.List) error {
	currentTicket := state.CurrentTicket()
	listState := state.ListState
	marker := string(node.Marker)
	list := body.NewBulletListItemSegment(listState.ListLevel(), marker)

	listState.StartBulletList(marker)
	currentTicket.AddBodySegment(list)
	err := ParseBodyChildren(state, node)
	listState.CompleteList()

	return err
}

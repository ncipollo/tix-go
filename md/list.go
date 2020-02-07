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

	listState.StartOrderedList(node.Start)

	listStart := body.NewOrderedListStartSegment(listState.ListLevel(), node.Start)
	currentTicket.AddBodySegment(listStart)

	err := ParseBodyChildren(state, node)

	listEnd := body.NewOrderedListEndSegment(listState.ListLevel())
	currentTicket.AddBodySegment(listEnd)

	listState.CompleteList()

	return err
}

func (l ListSegmentParser) parseBulletList(state *State, node *ast.List) error {
	currentTicket := state.CurrentTicket()
	listState := state.ListState
	marker := string(node.Marker)

	listState.StartBulletList(marker)

	listStart := body.NewBulletListStartSegment(listState.ListLevel(), marker)
	currentTicket.AddBodySegment(listStart)

	err := ParseBodyChildren(state, node)

	listEnd := body.NewBulletListEndSegment(listState.ListLevel())
	currentTicket.AddBodySegment(listEnd)

	listState.CompleteList()

	return err
}

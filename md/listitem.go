package md

import (
	"errors"
	"github.com/yuin/goldmark/ast"
	"tix/ticket/body"
)

type ListItemSegmentParser struct{}

func NewListItemSegmentParser() *ListItemSegmentParser {
	return &ListItemSegmentParser{}
}

func (l ListItemSegmentParser) Parse(state *State, node ast.Node) error {
	currentList := state.ListState.CurrentList()
	if currentList == nil {
		return errors.New("a list item must be within a list")
	}

	currentTicket := state.CurrentTicket()
	currentLevel := state.ListState.ListLevel()

	listItemNode := node.(*ast.ListItem)

	if currentList.IsOrdered {
		listItem := body.NewOrderedListItemSegment(currentLevel, currentList.CurrentNumber)
		currentTicket.AddBodySegment(listItem)
		currentList.CurrentNumber++
	} else {
		listItem := body.NewBulletListItemSegment(currentLevel, currentList.Marker)
		currentTicket.AddBodySegment(listItem)
	}

	lineBreak := body.NewLineBreakSegment()
	currentTicket.AddBodySegment(lineBreak)
	err := ParseBodyChildren(state, listItemNode)

	return err
}

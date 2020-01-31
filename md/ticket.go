package md

import (
	"errors"
	"github.com/yuin/goldmark/ast"
)

type TicketParser struct {
	bodySegmentParsers map[ast.NodeKind]BodySegmentParser
}

func NewTicketParser() *TicketParser  {
	return &TicketParser{}
}

func (p *TicketParser) Parse(state *State, rootNode ast.Node) error {
	for node := rootNode.FirstChild(); node != nil; node = node.NextSibling() {
		kind := node.Kind()
		if kind == ast.KindHeading {
			heading := node.(*ast.Heading)
			err :=  p.parseHeading(state, heading)
			if err != nil {
				return err
			}
		} else {
			if state.TicketLevel() == 0 {
				return errors.New(":scream: ticket information must be under a heading")
			}
			ParseBody(state, node)
		}
	}

	// Complete remaining open tickets
	state.CompleteAllTickets()

	return nil
}

func (p *TicketParser) parseHeading(state *State, heading *ast.Heading) error {
	levelDifference := state.TicketLevel() - heading.Level
	switch {
	case levelDifference == 0:
		state.CompleteTicket()
		state.StartTicket()
		return p.parseTicketTitle(state, heading)
	case levelDifference > 0:
		for ii := 0; ii <= levelDifference; ii++ {
			state.CompleteTicket()
		}
		state.StartTicket()
		return p.parseTicketTitle(state, heading)
	case levelDifference > -2:
		state.StartTicket()
		return p.parseTicketTitle(state, heading)
	default:
		return errors.New("new heading level can not be one greater than previous level")
	}
}


func (p *TicketParser) parseTicketTitle(state *State, heading *ast.Heading) error {
	node := heading.FirstChild()
	if node == nil {
		return errors.New("heading must have a title")
	}
	if node.Kind() == ast.KindText {
		ticket := state.CurrentTicket()
		text := node.(*ast.Text)
		value := text.Segment.Value(state.SourceData)
		ticket.Title = string(value)
		return nil
	} else {
		return errors.New("non-text heading titles are not supported")
	}
}

package md

import (
	"tix/ticket"
	"tix/ticket/body"
)

type State struct {
	SourceData  []byte
	RootTickets []*ticket.Ticket
	TicketPath  []*ticket.Ticket
	WorkingBody body.Segment
}

func newState(sourceData []byte) *State {
	return &State{SourceData: sourceData}
}

func (s *State) CurrentTicket() *ticket.Ticket {
	index := len(s.TicketPath) - 1
	if index < 0 {
		return nil
	}

	return s.TicketPath[index]
}

func (s *State) TicketLevel() int {
	return len(s.TicketPath)
}

func (s *State) NeedsTicketTitle() bool {
	currentTicket := s.CurrentTicket()
	if currentTicket == nil {
		return false
	}
	return currentTicket.Title == ""
}

func (s *State) StartTicket() {
	currentTicket := s.CurrentTicket()
	newTicket := ticket.NewTicket()

	if currentTicket != nil {
		currentTicket.AddSubticket(newTicket)
	}

	s.TicketPath = append(s.TicketPath, newTicket)
}

func (s *State) CompleteTicket() {
	index := len(s.TicketPath) - 1

	if index > 0 {
		s.TicketPath = s.TicketPath[:index]
	} else if index == 0 {
		s.RootTickets = append(s.RootTickets, s.CurrentTicket())
		s.TicketPath = s.TicketPath[:index]
	}
}

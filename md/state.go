package md

import (
	"tix/ticket"
)

type State struct {
	ListState   *ListState
	RootTickets []*ticket.Ticket
	SourceData  []byte
	FieldState  *FieldState
	TicketPath  []*ticket.Ticket
}

func newState(sourceData []byte, fieldState *FieldState) *State {
	return &State{ListState: NewListState(), SourceData: sourceData, FieldState: fieldState}
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
	level := s.TicketLevel()
	fields := s.FieldState.FieldsForLevel(level)
	newTicket := ticket.NewTicketWithFields(fields)

	if currentTicket != nil {
		currentTicket.AddSubticket(newTicket)
	}

	s.TicketPath = append(s.TicketPath, newTicket)
}

func (s *State) CompleteTicket() {
	s.buildTraversal()

	index := len(s.TicketPath) - 1

	if index > 0 {
		s.TicketPath = s.TicketPath[:index]
	} else if index == 0 {
		s.RootTickets = append(s.RootTickets, s.CurrentTicket())
		s.TicketPath = s.TicketPath[:index]
	}
}

func (s *State) buildTraversal() {
	currentTicket := s.CurrentTicket()
	if currentTicket == nil {
		return
	}
	currentTicket.BuildTraversal()
}

func (s *State) CompleteAllTickets() {
	ticketCount := len(s.TicketPath)
	for ii := 0; ii < ticketCount; ii++ {
		s.CompleteTicket()
	}
}

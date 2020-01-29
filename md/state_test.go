package md

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"tix/ticket"
)

func Test_state_CompleteAllTickets(t *testing.T) {
	state := newState(nil, nil)
	state.StartTicket()
	state.StartTicket()
	state.StartTicket()

	state.CompleteAllTickets()

	assert.Empty(t, state.TicketPath)
}

func Test_state_CompleteTicket_AddsToRootTickets(t *testing.T) {
	state := newState(nil, nil)
	state.StartTicket()
	state.CurrentTicket().Title = "1"
	state.CompleteTicket()
	state.StartTicket()
	state.CurrentTicket().Title = "2"
	state.CompleteTicket()

	assert.Len(t, state.RootTickets, 2)
	assert.Equal(t, &ticket.Ticket{Title: "1"}, state.RootTickets[0])
	assert.Equal(t, &ticket.Ticket{Title: "2"}, state.RootTickets[1])
}

func Test_state_CompleteTicket_PopsAllTickets(t *testing.T) {
	state := newState(nil, nil)
	for i := 0; i < 10; i++ {
		state.StartTicket()
	}
	for i := 0; i < 10; i++ {
		state.CompleteTicket()
	}

	assert.Empty(t, state.TicketPath)
}

func Test_state_CurrentTicket_ReturnsNilWhenEmpty(t *testing.T) {
	state := newState(nil, nil)

	currentTicket := state.CurrentTicket()

	assert.Nil(t, currentTicket)
}

func Test_state_CurrentTicket_ReturnsSubticket(t *testing.T) {
	state := newState(nil, nil)
	state.StartTicket()
	state.StartTicket()
	state.TicketPath[1].Title = "subticket"

	currentTicket := state.CurrentTicket()

	assert.Equal(t, state.TicketPath[1], currentTicket)
}

func Test_state_CurrentTicket_ReturnsTicket(t *testing.T) {
	state := newState(nil, nil)
	state.StartTicket()
	state.TicketPath[0].Title = "ticket"

	currentTicket := state.CurrentTicket()

	assert.Equal(t, state.TicketPath[0], currentTicket)
}

func Test_state_NeedsTicketTitle_FalseWhenNoTicket(t *testing.T) {
	state := newState(nil, nil)

	assert.False(t, state.NeedsTicketTitle())
}

func Test_state_NeedsTicketTitle_FalseTicketHasTitle(t *testing.T) {
	state := newState(nil, nil)
	state.StartTicket()
	state.CurrentTicket().Title = "title"

	assert.False(t, state.NeedsTicketTitle())
}

func Test_state_NeedsTicketTitle_TrueWhenTicketMissingTitle(t *testing.T) {
	state := newState(nil, nil)
	state.StartTicket()

	assert.True(t, state.NeedsTicketTitle())
}

func Test_state_StartTicket_AddsFields(t *testing.T) {
	fields := map[string]interface{}{"foo" : "map"}
	state := newState(nil, fields)
	state.StartTicket()

	assert.Equal(t,fields, state.CurrentTicket().Fields)
}

func Test_state_StartTicket_AddsTickets(t *testing.T) {
	state := newState(nil, nil)
	state.StartTicket()
	state.StartTicket()
	state.StartTicket()

	assert.Len(t, state.TicketPath, 3)
}

func Test_state_StartTicket_LinksTickets(t *testing.T) {
	state := newState(nil, nil)
	state.StartTicket()
	state.CurrentTicket().Title = "root ticket"
	state.StartTicket()
	state.CurrentTicket().Title = "1"
	state.CompleteTicket()
	state.StartTicket()
	state.CurrentTicket().Title = "2"

	rootTicket := state.TicketPath[0]
	assert.Len(t, rootTicket.Subtickets, 2)
	assert.Equal(t, &ticket.Ticket{Title: "1"}, rootTicket.Subtickets[0])
	assert.Equal(t, &ticket.Ticket{Title: "2"}, rootTicket.Subtickets[1])
}

func TestState_TicketLevel(t *testing.T) {
	state := newState(nil, nil)
	state.StartTicket()
	state.StartTicket()

	assert.Equal(t, 2, state.TicketLevel())
}
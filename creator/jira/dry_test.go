package jira

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"tix/ticket"
	"tix/ticket/body"
)

func TestDryRunCreator_CreateTickets_WithEpics(t *testing.T) {
	var builder strings.Builder

	dryCreator := NewDryRunCreatorWithEpics(&builder)
	dryCreator.CreateTickets(createDryRunTickets())

	const expected = `Would have created tickets: :point_down:

-----------------
:rocket:Epic - ticket 1
body 1
-----------------

-----------------
:rocket:Epic - ticket 1
body 2
-----------------

-----------------
:rocket:Issue - sub-ticket
sub-body
-----------------

`
	assert.Equal(t, expected, builder.String())
}

func TestDryRunCreator_CreateTickets_WithoutEpics(t *testing.T) {
	var builder strings.Builder

	dryCreator := NewDryRunCreatorWithoutEpics(&builder)
	dryCreator.CreateTickets(createDryRunTickets())

	const expected = `Would have created tickets: :point_down:

-----------------
:rocket:Issue - ticket 1
body 1
-----------------

-----------------
:rocket:Issue - ticket 1
body 2
-----------------

-----------------
:rocket:Task - sub-ticket
sub-body
-----------------

`
	assert.Equal(t, expected, builder.String())
}

func createDryRunTickets() []*ticket.Ticket {
	ticketFields := map[string]interface{}{}

	ticket1 := ticket.NewTicketWithFields(ticketFields)
	ticket1.Title = "ticket 1"
	ticket1.AddBodySegment(body.NewTextSegment("body 1"))

	subTicket := ticket.NewTicketWithFields(ticketFields)
	subTicket.Title = "sub-ticket"
	subTicket.AddBodySegment(body.NewTextSegment("sub-body"))

	ticket2 := ticket.NewTicketWithFields(ticketFields)
	ticket2.Title = "ticket 1"
	ticket2.AddBodySegment(body.NewTextSegment("body 2"))
	ticket2.AddSubticket(subTicket)

	return []*ticket.Ticket{ticket1, ticket2}
}

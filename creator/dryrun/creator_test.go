package dryrun

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"tix/creator"
	"tix/render"
	"tix/ticket"
	"tix/ticket/body"
)

func TestCreator_CreateTickets_StartingLevelZero(t *testing.T) {
	var builder strings.Builder
	dryCreator := createDryRunCreator(&builder, 0)
	dryCreator.CreateTickets(createDryRunTickets())

	const expected = `Would have created tickets: :point_down:

-----------------
:rocket:Epic - ticket 1

Jira Fields:
- field1: 1
- field2: 2

body 1
-----------------

-----------------
:rocket:Epic - ticket 2
body 2
-----------------

-----------------
:rocket:Story - sub-ticket
sub-body
-----------------

Ticket Stats:
- Total Tickets: 3
- Epics: 2
- Stories: 1
- Tasks: 0
`
	assert.Equal(t, expected, builder.String())
}

func TestCreator_CreateTickets_StartingLevelOne(t *testing.T) {
	var builder strings.Builder
	dryCreator := createDryRunCreator(&builder, 1)
	dryCreator.CreateTickets(createDryRunTickets())

	const expected = `Would have created tickets: :point_down:

-----------------
:rocket:Story - ticket 1

Jira Fields:
- field1: 1
- field2: 2

body 1
-----------------

-----------------
:rocket:Story - ticket 2
body 2
-----------------

-----------------
:rocket:Task - sub-ticket
sub-body
-----------------

Ticket Stats:
- Total Tickets: 3
- Epics: 0
- Stories: 2
- Tasks: 1
`
	assert.Equal(t, expected, builder.String())
}

func createDryRunCreator(builder *strings.Builder, startingTicketLevel int) creator.TicketCreator {
	labels := []*LevelLabel{
		NewLevelLabel("epic", "epics"),
		NewLevelLabel("story", "stories"),
		NewLevelLabel("task", "tasks"),
	}
	return NewCreator(
		builder,
		labels,
		render.NewJiraBodyRenderer(),
		startingTicketLevel,
		"jira",
	)
}

func createDryRunTickets() []*ticket.Ticket {
	ticketFields := map[string]interface{}{
		"field1": 1,
		"field2": 2,
	}

	ticket1 := ticket.NewTicketWithFields(ticketFields)
	ticket1.Title = "ticket 1"
	ticket1.AddBodySegment(body.NewTextSegment("body 1"))

	subTicket := ticket.NewTicketWithFields(map[string]interface{}{})
	subTicket.Title = "sub-ticket"
	subTicket.AddBodySegment(body.NewTextSegment("sub-body"))

	ticket2 := ticket.NewTicketWithFields(map[string]interface{}{})
	ticket2.Title = "ticket 2"
	ticket2.AddBodySegment(body.NewTextSegment("body 2"))
	ticket2.AddSubticket(subTicket)

	return []*ticket.Ticket{ticket1, ticket2}
}

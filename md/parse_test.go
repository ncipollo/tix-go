package md

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"tix/ticket"
	"tix/ticket/body"
)

const md = `
# Epic 1

Body 1

# Epic 2
Body 2
`

func TestParse(t *testing.T) {
	source := []byte(md)
	fields := map[string]interface{}{
		"foo": "bar",
	}
	fieldState := NewFieldState()
	fieldState.SetDefaultFields(fields)
	parser := NewParser(fieldState)
	tickets, err := parser.Parse(source)

	expectedTicket1 := ticket.NewTicketWithFields(fields)
	expectedTicket1.Title = "Epic 1"
	expectedTicket1.AddBodyLineBreak()
	expectedTicket1.AddBodySegment(body.NewTextSegment("Body 1"))
	expectedTicket1.AddBodyLineBreak()
	expectedTicket1.BuildTraversal()

	expectedTicket2 := ticket.NewTicketWithFields(fields)
	expectedTicket2.Title = "Epic 2"
	expectedTicket2.AddBodySegment(body.NewTextSegment("Body 2"))
	expectedTicket2.AddBodyLineBreak()
	expectedTicket2.BuildTraversal()

	expectedTickets := []*ticket.Ticket{expectedTicket1,expectedTicket2}
	assert.NoError(t, err)
	assert.Equal(t, expectedTickets, tickets)
}

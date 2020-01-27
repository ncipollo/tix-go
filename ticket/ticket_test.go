package ticket

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"tix/ticket/body"
)

func TestTicket_AddBodySegment(t *testing.T) {
	ticket := NewTicket()
	segment := body.NewTextSegment("")

	ticket.AddBodySegment(segment)

	assert.Equal(t, []body.Segment{segment}, ticket.Body)
}

func TestTicket_AddLineBreakSegment(t *testing.T) {
	ticket := NewTicket()

	ticket.AddBodyLineBreak()

	assert.Equal(t, []body.Segment{body.NewLineBreakSegment()}, ticket.Body)
}

func TestTicket_AddSubticket(t *testing.T) {
	ticket := NewTicket()
	subticket := NewTicket()

	ticket.AddSubticket(subticket)

	assert.Equal(t, []*Ticket{subticket}, ticket.Subtickets)
}

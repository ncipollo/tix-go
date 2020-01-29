package render

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"tix/ticket"
	"tix/ticket/body"
)

func TestTicket(t *testing.T) {
	currentTicket := ticket.NewTicket()
	currentTicket.Body = []body.Segment{
		body.NewTextSegment("line 1"),
		body.NewLineBreakSegment(),
		body.NewTextSegment("line 2"),
	}
	renderer := NewJiraBodyRenderer()

	text := Ticket(currentTicket, renderer)

	expected := "line 1\nline 2"
	assert.Equal(t, expected, text)
}
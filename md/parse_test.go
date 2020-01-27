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
	parser := NewParser()
	tickets, err := parser.Parse(source)

	expectedTickets := []*ticket.Ticket{
		{
			Title: "Epic 1",
			Body: []body.Segment{
				body.NewTextSegment("Body 1"),
				body.NewLineBreakSegment(),
			},
		},
		{
			Title: "Epic 2",
			Body: []body.Segment{
				body.NewTextSegment("Body 2"),
				body.NewLineBreakSegment(),
			},
		},
	}
	assert.NoError(t, err)
	assert.Equal(t, expectedTickets, tickets)
}

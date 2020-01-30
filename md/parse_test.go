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
	parser := NewParser(fields)
	tickets, err := parser.Parse(source)

	expectedTickets := []*ticket.Ticket{
		{
			Fields: fields,
			Title:  "Epic 1",
			Body: []body.Segment{
				body.NewTextSegment("Body 1"),
				body.NewLineBreakSegment(),
			},
		},
		{
			Fields: fields,
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

package jira

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestTicketStats_Render(t *testing.T) {
	var builder strings.Builder
	stats := NewTicketStats()

	stats.CountTicket(0)
	stats.CountTicket(0)
	stats.CountTicket(0)
	stats.CountTicket(1)
	stats.CountTicket(1)
	stats.CountTicket(2)
	stats.Render(&builder)

	expected := `Ticket Stats:
- Total Tickets: 6
- Epics: 3
- Issues: 2
- Tasks: 1
`
	assert.Equal(t, expected, builder.String())
}

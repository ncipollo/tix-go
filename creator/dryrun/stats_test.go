package dryrun

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestTicketStats_RenderTicket(t *testing.T) {
	var builder strings.Builder
	stats := NewTicketStats([]*LevelLabel{
		NewLevelLabel("epic", "epics"),
		NewLevelLabel("story", "stories"),
		NewLevelLabel("task", "tasks"),
	})

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
- Stories: 2
- Tasks: 1
`
	assert.Equal(t, expected, builder.String())
}

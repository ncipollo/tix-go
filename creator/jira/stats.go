package jira

import (
	"fmt"
	"strings"
)

type TicketStats struct {
	epics  int
	issues int
	tasks  int
}

func NewTicketStats()  *TicketStats{
	return &TicketStats{}
}

func (t *TicketStats) CountTicket(level int) {
	switch level {
	case 0:
		t.epics++
	case 1:
		t.issues++
	case 2:
		t.tasks++
	}
}

func (t TicketStats) Render(builder *strings.Builder) {
	builder.WriteString("Ticket Stats:\n")
	builder.WriteString(fmt.Sprintf("- Total Tickets: %d\n", t.total()))
	builder.WriteString(fmt.Sprintf("- Epics: %d\n", t.epics))
	builder.WriteString(fmt.Sprintf("- Issues: %d\n", t.issues))
	builder.WriteString(fmt.Sprintf("- Tasks: %d\n", t.tasks))
}

func (t *TicketStats) total() int {
	return t.epics + t.issues + t.tasks
}

package jira

import (
	"fmt"
	"strings"
	"tix/render"
	"tix/ticket"
)

type DryRunCreator struct {
	builder             *strings.Builder
	startingTicketLevel int
}

func NewDryRunCreatorWithEpics(builder *strings.Builder) *DryRunCreator {
	return &DryRunCreator{builder: builder, startingTicketLevel: 0}
}

func NewDryRunCreatorWithoutEpics(builder *strings.Builder) *DryRunCreator {
	return &DryRunCreator{builder: builder, startingTicketLevel: 1}
}

func (d DryRunCreator) CreateTickets(tickets []*ticket.Ticket) {
	var renderer = render.NewJiraBodyRenderer()

	d.builder.WriteString("Would have created tickets: :point_down:\n\n")
	d.renderTicketsForLevel(tickets, d.startingTicketLevel, renderer)
}

func (d DryRunCreator) renderTicketsForLevel(tickets []*ticket.Ticket,
	level int,
	renderer render.BodyRenderer) {
	for _, currentTicket := range tickets {
		ticketString := render.Ticket(currentTicket, renderer)

		d.builder.WriteString("-----------------\n")
		d.builder.WriteString(d.title(currentTicket, level))
		d.builder.WriteString(ticketString)
		d.builder.WriteString("\n")
		d.builder.WriteString("-----------------\n\n")

		d.renderTicketsForLevel(currentTicket.Subtickets, level+1, renderer)
	}
}

func (d DryRunCreator) title(ticket *ticket.Ticket, level int) string {
	return fmt.Sprintf(":rocket:%s - %s\n", d.ticketType(level), ticket.Title)
}

func (d DryRunCreator) ticketType(level int) string {
	switch level {
	case 0:
		return "Epic"
	case 1:
		return "Issue"
	case 2:
		return "Subtask"
	default:
		return "Unknown Ticket Type"
	}
}

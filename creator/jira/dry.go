package jira

import (
	"fmt"
	"strings"
	"tix/logger"
	"tix/render"
	"tix/ticket"
)

type DryRunCreator struct {
	startingTicketLevel int
}

func NewDryRunCreatorWithEpics() *DryRunCreator {
	return &DryRunCreator{0}
}

func NewDryRunCreatorWithoutEpics() *DryRunCreator {
	return &DryRunCreator{1}
}

func (d DryRunCreator) CreateTickets(tickets []*ticket.Ticket) {
	var builder strings.Builder
	var renderer = render.NewJiraBodyRenderer()

	builder.WriteString("Would have created tickets: :point_down:\n\n")
	d.renderTicketsForLevel(tickets, &builder, d.startingTicketLevel, renderer)
	logger.Message(builder.String())
}

func (d DryRunCreator) renderTicketsForLevel(tickets []*ticket.Ticket,
	builder *strings.Builder,
	level int,
	renderer render.BodyRenderer) {
	for _, currentTicket := range tickets {
		ticketString := render.Ticket(currentTicket, renderer)

		builder.WriteString("-----------------\n")
		builder.WriteString(d.title(currentTicket, level))
		builder.WriteString(ticketString)
		builder.WriteString("\n")
		builder.WriteString("-----------------\n\n")

		d.renderTicketsForLevel(currentTicket.Subtickets, builder, level+1, renderer)
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

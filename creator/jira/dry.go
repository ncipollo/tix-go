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
	stats               *TicketStats
}

func NewDryRunCreatorWithEpics(builder *strings.Builder) *DryRunCreator {
	return &DryRunCreator{
		builder:             builder,
		startingTicketLevel: 0,
		stats: NewTicketStats(),
	}
}

func NewDryRunCreatorWithoutEpics(builder *strings.Builder) *DryRunCreator {
	return &DryRunCreator{
		builder:             builder,
		startingTicketLevel: 1,
		stats: NewTicketStats(),
	}
}

func (d DryRunCreator) CreateTickets(tickets []*ticket.Ticket) {
	var renderer = render.NewJiraBodyRenderer()

	d.builder.WriteString("Would have created tickets: :point_down:\n\n")
	d.renderTicketsForLevel(tickets, d.startingTicketLevel, renderer)
	d.stats.Render(d.builder)
}

func (d DryRunCreator) renderTicketsForLevel(tickets []*ticket.Ticket,
	level int,
	renderer render.BodyRenderer) {
	for _, currentTicket := range tickets {
		ticketString := render.Ticket(currentTicket, renderer)

		d.builder.WriteString("-----------------\n")
		d.builder.WriteString(d.title(currentTicket, level))
		d.renderFields(currentTicket)
		d.builder.WriteString(ticketString)
		d.builder.WriteString("\n")
		d.builder.WriteString("-----------------\n\n")

		d.renderTicketsForLevel(currentTicket.Subtickets, level+1, renderer)
		d.stats.CountTicket(level)
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
		return "Task"
	default:
		return "Unknown Ticket Type"
	}
}

func (d DryRunCreator) renderFields(ticket *ticket.Ticket) {
	fields := ticket.Fields("jira")
	if len(fields) == 0 {
		return
	}

	d.builder.WriteString("\nJira Fields:\n")
	for key, value := range fields {
		fieldString := fmt.Sprintf("- %s: %v\n", key, value)
		d.builder.WriteString(fieldString)
	}
	d.builder.WriteString("\n")
}

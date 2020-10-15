package dryrun

import (
	"fmt"
	"sort"
	"strings"
	"tix/creator"
	"tix/render"
	"tix/ticket"
)

type Creator struct {
	builder             *strings.Builder
	levelLabels         []*LevelLabel
	renderer            render.BodyRenderer
	startingTicketLevel int
	stats               *TicketStats
	ticketSystem        string
}

func NewCreator(
	builder *strings.Builder,
	levelLabels []*LevelLabel,
	renderer render.BodyRenderer,
	startingTicketLevel int,
	ticketSystem string,
) creator.TicketCreator {
	return &Creator{
		builder:             builder,
		levelLabels:         levelLabels,
		renderer:            renderer,
		startingTicketLevel: startingTicketLevel,
		stats:               NewTicketStats(levelLabels),
		ticketSystem:        ticketSystem,
	}
}

func (c *Creator) CreateTickets(tickets []*ticket.Ticket) {
	c.builder.WriteString("Would have created tickets: :point_down:\n\n")
	c.renderTicketsForLevel(tickets, c.startingTicketLevel)
	c.stats.Render(c.builder)
}

func (c Creator) renderTicketsForLevel(tickets []*ticket.Ticket,
	level int) {
	for _, currentTicket := range tickets {
		ticketString := render.Ticket(currentTicket, c.renderer)

		c.builder.WriteString("-----------------\n")
		c.builder.WriteString(c.title(currentTicket, level))
		c.renderFields(currentTicket)
		c.builder.WriteString(ticketString)
		c.builder.WriteString("\n")
		c.builder.WriteString("-----------------\n\n")

		c.renderTicketsForLevel(currentTicket.Subtickets, level+1)
		c.stats.CountTicket(level)
	}
}

func (c Creator) title(ticket *ticket.Ticket, level int) string {
	return fmt.Sprintf(":rocket:%s - %s\n", c.ticketType(level), ticket.Title)
}

func (c Creator) ticketType(level int) string {
	return c.levelLabels[level].Singular()
}

func (c Creator) renderFields(ticket *ticket.Ticket) {
	fields := ticket.Fields(c.ticketSystem)
	if len(fields) == 0 {
		return
	}

	keys := sortedFieldKeys(fields)

	c.builder.WriteString(c.fieldsTitle())
	for _, key := range keys {
		value := fields[key]
		fieldString := fmt.Sprintf("- %s: %v\n", key, value)
		c.builder.WriteString(fieldString)
	}
	c.builder.WriteString("\n")
}

func (c Creator) fieldsTitle() string {
	return fmt.Sprintf("\n%s Fields:\n", strings.Title(c.ticketSystem))
}

func sortedFieldKeys(fields map[string]interface{}) []string {
	keys := make([]string, 0, len(fields))
	for k := range fields {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	return keys
}

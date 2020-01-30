package creator

import "tix/ticket"

type TicketCreator interface {
	CreateTickets(tickets []*ticket.Ticket)
}

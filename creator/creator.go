package creator

import "tix/ticket"

type TicketCreator interface {
	CreateTicket(tickets []*ticket.Ticket)
}

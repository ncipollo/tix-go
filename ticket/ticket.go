package ticket

import "tix/ticket/body"

type Ticket struct {
	Metadata   interface{}
	Title      string
	Body       []body.Segment
	Subtickets []*Ticket
}

func NewTicket() *Ticket {
	return &Ticket{}
}

func (t *Ticket) AddBodySegment(segment body.Segment) {
	t.Body = append(t.Body, segment)
}

func (t *Ticket) AddSubticket(ticket *Ticket) {
	t.Subtickets = append(t.Subtickets, ticket)
}

package ticket

import "tix/ticket/body"

type Ticket struct {
	Fields     map[string]interface{}
	Metadata   interface{}
	Title      string
	Body       []body.Segment
	Subtickets []*Ticket
}

func NewTicketWithFields(fields map[string]interface{}) *Ticket {
	return &Ticket{Fields: fields}
}

func NewTicket() *Ticket {
	return NewTicketWithFields(nil)
}

func (t *Ticket) AddBodySegment(segment body.Segment) {
	t.Body = append(t.Body, segment)
}

func (t *Ticket) AddBodyLineBreak() {
	lineBreak := body.NewLineBreakSegment()
	t.Body = append(t.Body, lineBreak)
}

func (t *Ticket) AddSubticket(ticket *Ticket) {
	t.Subtickets = append(t.Subtickets, ticket)
}

package ticket

import "tix/ticket/body"

type Ticket struct {
	Title      string
	Body       []body.Segment
	Subtickets []Ticket
}

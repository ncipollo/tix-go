package render

import (
	"strings"
	"tix/ticket"
	"tix/ticket/body"
)

type BodyRenderer interface {
	RenderSegment(segment body.Segment) string
}

func Ticket(ticket ticket.Ticket, renderer BodyRenderer) string {
	var builder strings.Builder
	for _, segment := range ticket.Body {
		text := renderer.RenderSegment(segment)
		builder.WriteString(text)
	}
	return builder.String()
}

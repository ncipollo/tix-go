package ticket

import "tix/ticket/body"

type Ticket struct {
	DefaultFields        map[string]interface{}
	FieldsByTicketSystem map[string]map[string]interface{}
	Metadata             interface{}
	Title                string
	Body                 []body.Segment
	Subtickets           []*Ticket
}

func NewTicketWithFields(fields map[string]interface{}) *Ticket {
	return &Ticket{DefaultFields: fields}
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

func (t *Ticket) AddFieldsForTicketSystem(fields map[string]interface{}, ticketSystem string) {
	combinedFields := MergeFields(t.DefaultFields, fields)
	t.FieldsByTicketSystem[ticketSystem] = combinedFields
}

func (t *Ticket) AddSubticket(ticket *Ticket) {
	t.Subtickets = append(t.Subtickets, ticket)
}

func (t *Ticket) Fields(ticketSystem string) map[string]interface{} {
	fields := t.FieldsByTicketSystem[ticketSystem]
	if fields != nil {
		return fields
	}
	return t.DefaultFields
}

func MergeFields(
	baseFields map[string]interface{},
	overlayFields map[string]interface{}) map[string]interface{} {
	combinedFields := make(map[string]interface{})
	if baseFields != nil {
		for key, value := range baseFields {
			combinedFields[key] = value
		}
	}
	if overlayFields != nil {
		for key, value := range overlayFields {
			combinedFields[key] = value
		}
	}

	return combinedFields
}

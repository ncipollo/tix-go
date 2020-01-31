package ticket

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"tix/ticket/body"
)

func TestTicket_AddBodySegment(t *testing.T) {
	ticket := NewTicket()
	segment := body.NewTextSegment("")

	ticket.AddBodySegment(segment)

	assert.Equal(t, []body.Segment{segment}, ticket.Body)
}

func TestTicket_AddLineBreakSegment(t *testing.T) {
	ticket := NewTicket()

	ticket.AddBodyLineBreak()

	assert.Equal(t, []body.Segment{body.NewLineBreakSegment()}, ticket.Body)
}

func TestTicket_AddSubticket(t *testing.T) {
	ticket := NewTicket()
	subticket := NewTicket()

	ticket.AddSubticket(subticket)

	assert.Equal(t, []*Ticket{subticket}, ticket.Subtickets)
}

func TestTicket_Fields_ReturnsDefaultFields(t *testing.T) {
	defaultFields := map[string]interface{}{
		"default": "default",
	}
	jiraFields := map[string]interface{}{
		"jira": "jira",
	}
	ticket := NewTicketWithFields(defaultFields)

	ticket.AddFieldsForTicketSystem(jiraFields, "jira")
	combinedFields := ticket.Fields("github")

	expected := map[string]interface{}{
		"default": "default",
	}
	assert.Equal(t, expected, combinedFields)
}

func TestTicket_Fields_ReturnsTicketSystemSpecificFields(t *testing.T) {
	defaultFields := map[string]interface{}{
		"default": "default",
	}
	jiraFields := map[string]interface{}{
		"jira": "jira",
	}
	ticket := NewTicketWithFields(defaultFields)

	ticket.AddFieldsForTicketSystem(jiraFields, "jira")
	combinedFields := ticket.Fields("jira")

	expected := map[string]interface{}{
		"default": "default",
		"jira":    "jira",
	}
	assert.Equal(t, expected, combinedFields)
}

func TestTicket_UpdateDefaultFields(t *testing.T) {
	defaultFields := map[string]interface{}{
		"original": "foo",
	}
	updatedFields := map[string]interface{}{
		"updated": "bar",
	}
	ticket := NewTicketWithFields(defaultFields)

	ticket.UpdateDefaultFields(updatedFields)

	expected := map[string]interface{}{
		"original": "foo",
		"updated": "bar",
	}
	assert.Equal(t, expected, ticket.DefaultFields)
}
package jira

import (
	"github.com/andygrunwald/go-jira"
	"github.com/stretchr/testify/assert"
	"testing"
	"tix/render"
	"tix/ticket"
	"tix/ticket/body"
)

func TestIssues_FromTicket_Epic(t *testing.T) {
	issues := createIssues()
	ticket := createTicket()

	newIssue := issues.FromTicket(ticket, "", 1)

	expected := &jira.Issue{
		Fields: &jira.IssueFields{
			Components: []*jira.Component{
				{Name: "component1"},
				{Name: "component2"},
			},
			Description: "body",
			Labels:      []string{"label1", "label2"},
			Type:        jira.IssueType{Name: "type"},
			Project:     jira.Project{Key: "project"},
			Summary:     "title",
			Unknowns: map[string]interface{}{
				"field2": "epic",
				"field3": "random",
			},
		},
	}

	assert.Equal(t, expected, newIssue)
}

func createIssues() *Issues {
	jiraFields := []jira.Field{
		{ID: "field1", Name: "Epic Link"},
		{ID: "field2", Name: "Epic Name"},
		{ID: "field3", Name: "Random"},
	}
	return NewIssues(jiraFields, render.NewJiraBodyRenderer())
}

func createTicket() *ticket.Ticket {
	ticketFields := map[string]interface{}{
		"components": []string{"component1", "component2"},
		"labels":     []string{"label1", "label2"},
		"epic name":  "epic",
		"project":    "project",
		"random":     "random",
		"type":       "type",
	}
	newTicket := ticket.NewTicketWithFields(ticketFields)

	newTicket.Title = "title"
	newTicket.AddBodySegment(body.NewTextSegment("body"))

	return newTicket
}

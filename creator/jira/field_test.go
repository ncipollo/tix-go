package jira

import (
	"github.com/andygrunwald/go-jira"
	"github.com/stretchr/testify/assert"
	"testing"
	"tix/ticket"
)

func TestIssueFields_Components_EmptyForInvalidType(t *testing.T) {
	ticketFields := map[string]interface{}{
		"components": "lol nope",
	}
	issueFields := NewIssueFields(nil, ticket.NewTicketWithFields(ticketFields))

	jiraComps := issueFields.Components()

	assert.Empty(t, jiraComps)
}

func TestIssueFields_Components_WithComponents(t *testing.T) {
	ticketFields := map[string]interface{}{
		"components": []string{"one", "two"},
	}
	issueFields := NewIssueFields(nil, ticket.NewTicketWithFields(ticketFields))

	jiraComps := issueFields.Components()

	expected := []*jira.Component{{Name: "one"}, {Name: "two"}}
	assert.Equal(t, expected, jiraComps)
}

func TestIssueFields_EpicType_DefinedType(t *testing.T) {
	ticketFields := map[string]interface{}{
		"type": "test",
	}
	issueFields := NewIssueFields(nil, ticket.NewTicketWithFields(ticketFields))

	issueType := issueFields.EpicType()

	expected := jira.IssueType{Name: "test"}
	assert.Equal(t, expected, issueType)
}

func TestIssueFields_EpicType_DefaultType(t *testing.T) {
	ticketFields := map[string]interface{}{}
	issueFields := NewIssueFields(nil, ticket.NewTicketWithFields(ticketFields))

	issueType := issueFields.EpicType()

	expected := jira.IssueType{Name: "Epic"}
	assert.Equal(t, expected, issueType)
}

func TestIssueFields_IssueType_DefinedType(t *testing.T) {
	ticketFields := map[string]interface{}{
		"type": "test",
	}
	issueFields := NewIssueFields(nil, ticket.NewTicketWithFields(ticketFields))

	issueType := issueFields.IssueType()

	expected := jira.IssueType{Name: "test"}
	assert.Equal(t, expected, issueType)
}

func TestIssueFields_Labels(t *testing.T) {
	ticketFields := map[string]interface{}{}
	issueFields := NewIssueFields(nil, ticket.NewTicketWithFields(ticketFields))

	issueType := issueFields.IssueType()

	expected := jira.IssueType{Name: "Story"}
	assert.Equal(t, expected, issueType)
}

func TestIssueFields_Project(t *testing.T) {
	ticketFields := map[string]interface{}{
		"project": "test",
	}
	issueFields := NewIssueFields(nil, ticket.NewTicketWithFields(ticketFields))

	project := issueFields.Project()

	expected := jira.Project{Key: "test"}
	assert.Equal(t, expected, project)
}

func TestIssueFields_Unknowns(t *testing.T) {
	ticketFields := map[string]interface{}{
		"epic name": "epic",
		"option":    "option",
		"Random":    "random",
		"type":      "type", // should be excluded since this is a known key
	}
	issueFields := NewIssueFields(createJiraFields(), ticket.NewTicketWithFields(ticketFields))

	unknowns := issueFields.Unknowns()

	expected := map[string]interface{}{
		"field1": "epic",
		"field3": map[string]interface{}{
			"value": "option",
		},
		"Random": "random",
	}
	assert.Equal(t, expected, unknowns)
}

func createJiraFields() []jira.Field {
	return []jira.Field{
		{ID: "field1", Name: "Epic Name"},
		{ID: "field2", Name: "Type"},
		{ID: "field3", Name: "Option", Schema: jira.FieldSchema{Type: "option"}},
	}
}

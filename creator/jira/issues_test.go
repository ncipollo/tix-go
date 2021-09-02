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
    newTicket := createTicket(false)

    newIssue := issues.FromTicket(newTicket, nil, 0)

    expected := &jira.Issue{
        Fields: &jira.IssueFields{
            AffectsVersions: []*jira.AffectsVersion{{Name: "1"}},
            Components: []*jira.Component{
                {Name: "component1"},
                {Name: "component2"},
            },
            Description: "body",
            FixVersions: []*jira.FixVersion{{Name: "2"}},
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

func TestIssues_FromTicket_Story_NoParent(t *testing.T) {
    issues := createIssues()
    newTicket := createTicket(false)

    newIssue := issues.FromTicket(newTicket, nil, 1)

    expected := &jira.Issue{
        Fields: &jira.IssueFields{
            AffectsVersions: []*jira.AffectsVersion{{Name: "1"}},
            Components: []*jira.Component{
                {Name: "component1"},
                {Name: "component2"},
            },
            Description: "body",
            FixVersions: []*jira.FixVersion{{Name: "2"}},
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

func TestIssues_FromTicket_Story_WithParent_UsingEpicLink(t *testing.T) {
    issues := createIssues()
    newTicket := createTicket(false)

    newIssue := issues.FromTicket(newTicket, &jira.Issue{Key: "parent"}, 1)

    expected := &jira.Issue{
        Fields: &jira.IssueFields{
            AffectsVersions: []*jira.AffectsVersion{{Name: "1"}},
            Components: []*jira.Component{
                {Name: "component1"},
                {Name: "component2"},
            },
            Description: "body",
            FixVersions: []*jira.FixVersion{{Name: "2"}},
            Labels:      []string{"label1", "label2"},
            Type:        jira.IssueType{Name: "type"},
            Project:     jira.Project{Key: "project"},
            Summary:     "title",
            Unknowns: map[string]interface{}{
                "field1": "parent",
                "field2": "epic",
                "field3": "random",
            },
        },
    }

    assert.Equal(t, expected, newIssue)
}

func TestIssues_FromTicket_Story_WithParent_UsingProject(t *testing.T) {
    issues := createIssues()
    newTicket := createTicket(true)

    newIssue := issues.FromTicket(newTicket, &jira.Issue{Key: "parent"}, 1)

    expected := &jira.Issue{
        Fields: &jira.IssueFields{
            AffectsVersions: []*jira.AffectsVersion{{Name: "1"}},
            Components: []*jira.Component{
                {Name: "component1"},
                {Name: "component2"},
            },
            Description: "body",
            FixVersions: []*jira.FixVersion{{Name: "2"}},
            Labels:      []string{"label1", "label2"},
            Type:        jira.IssueType{Name: "type"},
            Parent:      &jira.Parent{Key: "parent"},
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

func TestIssues_FromTicket_Task(t *testing.T) {
    issues := createIssues()
    newTicket := createTicket(false)

    newIssue := issues.FromTicket(newTicket, &jira.Issue{ID: "1"}, 2)

    expected := &jira.Issue{
        Fields: &jira.IssueFields{
            AffectsVersions: []*jira.AffectsVersion{{Name: "1"}},
            Components: []*jira.Component{
                {Name: "component1"},
                {Name: "component2"},
            },
            Description: "body",
            FixVersions: []*jira.FixVersion{{Name: "2"}},
            Labels:      []string{"label1", "label2"},
            Type:        jira.IssueType{Name: "type"},
            Parent:      &jira.Parent{ID: "1"},
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

func createTicket(useParent bool) *ticket.Ticket {
    ticketFields := map[string]interface{}{
        "affects versions": []interface{}{"1"},
        "components":       []interface{}{"component1", "component2"},
        "labels":           []interface{}{"label1", "label2"},
        "epic name":        "epic",
        "fix versions":     []interface{}{"2"},
        "project":          "project",
        "random":           "random",
        "type":             "type",
        "use_parent":       useParent,
    }
    newTicket := ticket.NewTicketWithFields(ticketFields)

    newTicket.Title = "title"
    newTicket.AddBodySegment(body.NewTextSegment("body"))

    return newTicket
}

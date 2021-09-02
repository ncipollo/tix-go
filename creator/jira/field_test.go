package jira

import (
    "github.com/andygrunwald/go-jira"
    "github.com/stretchr/testify/assert"
    "testing"
    "tix/ticket"
)

func TestIssueFields_AddDefaultEpicName_AddDefaultIfEpicEmpty(t *testing.T) {
    ticketFields := map[string]interface{}{
        "field2": "",
    }
    issueFields := NewIssueFields(createJiraFields(), ticket.NewTicketWithFields(ticketFields))

    unknowns := issueFields.Unknowns()
    issueFields.AddDefaultEpicName(unknowns, "new name")

    expected := map[string]interface{}{
        "field2": "new name",
    }
    assert.Equal(t, expected, unknowns)
}

func TestIssueFields_AddDefaultEpicName_AddDefaultIfEpicMissing(t *testing.T) {
    ticketFields := map[string]interface{}{}
    issueFields := NewIssueFields(createJiraFields(), ticket.NewTicketWithFields(ticketFields))

    unknowns := issueFields.Unknowns()
    issueFields.AddDefaultEpicName(unknowns, "new name")

    expected := map[string]interface{}{
        "field2": "new name",
    }
    assert.Equal(t, expected, unknowns)
}

func TestIssueFields_AddDefaultEpicName_DoNothingIfEpicNameExists(t *testing.T) {
    ticketFields := map[string]interface{}{
        "field2": "epic",
    }
    issueFields := NewIssueFields(createJiraFields(), ticket.NewTicketWithFields(ticketFields))

    unknowns := issueFields.Unknowns()
    issueFields.AddDefaultEpicName(unknowns, "new name")

    expected := map[string]interface{}{
        "field2": "epic",
    }
    assert.Equal(t, expected, unknowns)
}

func TestIssueFields_AffectsVersions_NilForInvalidType(t *testing.T) {
    ticketFields := map[string]interface{}{}
    issueFields := NewIssueFields(nil, ticket.NewTicketWithFields(ticketFields))

    versions := issueFields.AffectsVersions()

    assert.Nil(t, versions)
}

func TestIssueFields_AffectsVersions_WithComponents(t *testing.T) {
    ticketFields := map[string]interface{}{
        "affects versions": []interface{}{"1", "2"},
    }
    issueFields := NewIssueFields(nil, ticket.NewTicketWithFields(ticketFields))

    versions := issueFields.AffectsVersions()

    expected := []*jira.AffectsVersion{{Name: "1"}, {Name: "2"}}
    assert.Equal(t, expected, versions)
}

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
        "components": []interface{}{"one", "two"},
    }
    issueFields := NewIssueFields(nil, ticket.NewTicketWithFields(ticketFields))

    jiraComps := issueFields.Components()

    expected := []*jira.Component{{Name: "one"}, {Name: "two"}}
    assert.Equal(t, expected, jiraComps)
}

func TestIssueFields_FixVersions_NilForInvalidType(t *testing.T) {
    ticketFields := map[string]interface{}{}
    issueFields := NewIssueFields(nil, ticket.NewTicketWithFields(ticketFields))

    versions := issueFields.FixVersions()

    assert.Nil(t, versions)
}

func TestIssueFields_FixVersions_WithComponents(t *testing.T) {
    ticketFields := map[string]interface{}{
        "fix versions": []interface{}{"1", "2"},
    }
    issueFields := NewIssueFields(nil, ticket.NewTicketWithFields(ticketFields))

    versions := issueFields.FixVersions()

    expected := []*jira.FixVersion{{Name: "1"}, {Name: "2"}}
    assert.Equal(t, expected, versions)
}

func TestIssueFields_EpicLinkKey(t *testing.T) {
    ticketFields := map[string]interface{}{}
    issueFields := NewIssueFields(createJiraFields(), ticket.NewTicketWithFields(ticketFields))

    key := issueFields.EpicLinkKey()

    assert.Equal(t, "field1", key)
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

func TestIssueFields_Labels_EmptyForInvalidType(t *testing.T) {
    ticketFields := map[string]interface{}{
        "labels": 42,
    }
    issueFields := NewIssueFields(nil, ticket.NewTicketWithFields(ticketFields))

    labels := issueFields.Labels()

    assert.Empty(t, labels)
}

func TestIssueFields_Labels_WithLabels(t *testing.T) {
    ticketFields := map[string]interface{}{
        "labels": []interface{}{"label1", "label2"},
    }
    issueFields := NewIssueFields(nil, ticket.NewTicketWithFields(ticketFields))

    labels := issueFields.Labels()

    expected := []string{"label1", "label2"}
    assert.Equal(t, expected, labels)
}

func TestIssueFields_Project_EmptyIfMissing(t *testing.T) {
    ticketFields := map[string]interface{}{
    }
    issueFields := NewIssueFields(nil, ticket.NewTicketWithFields(ticketFields))

    project := issueFields.Project()

    expected := jira.Project{Key: ""}
    assert.Equal(t, expected, project)
}

func TestIssueFields_Project_WithProject(t *testing.T) {
    ticketFields := map[string]interface{}{
        "project": "test",
    }
    issueFields := NewIssueFields(nil, ticket.NewTicketWithFields(ticketFields))

    project := issueFields.Project()

    expected := jira.Project{Key: "test"}
    assert.Equal(t, expected, project)
}

func TestIssueFields_TaskType_DefinedType(t *testing.T) {
    ticketFields := map[string]interface{}{
        "type": "test",
    }
    issueFields := NewIssueFields(nil, ticket.NewTicketWithFields(ticketFields))

    issueType := issueFields.TaskType()

    expected := jira.IssueType{Name: "test"}
    assert.Equal(t, expected, issueType)
}

func TestIssueFields_TaskType_DefaultType(t *testing.T) {
    ticketFields := map[string]interface{}{}
    issueFields := NewIssueFields(nil, ticket.NewTicketWithFields(ticketFields))

    issueType := issueFields.TaskType()

    expected := jira.IssueType{Name: "Task"}
    assert.Equal(t, expected, issueType)
}

func TestIssueFields_UseParent_DefinedType(t *testing.T) {
    ticketFields := map[string]interface{}{
        "use_parent": true,
    }
    issueFields := NewIssueFields(nil, ticket.NewTicketWithFields(ticketFields))

    useParent := issueFields.UseParent()

    assert.True(t, useParent)
}

func TestIssueFields_UseParent_DefaultType(t *testing.T) {
    ticketFields := map[string]interface{}{}
    issueFields := NewIssueFields(nil, ticket.NewTicketWithFields(ticketFields))

    useParent := issueFields.UseParent()

    assert.False(t, useParent)
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
        "field2": "epic",
        "field4": map[string]interface{}{
            "value": "option",
        },
        "Random": "random",
    }
    assert.Equal(t, expected, unknowns)
}

func createJiraFields() []jira.Field {
    return []jira.Field{
        {ID: "field1", Name: "Epic Link"},
        {ID: "field2", Name: "Epic Name"},
        {ID: "field3", Name: "Type"},
        {ID: "field4", Name: "Option", Schema: jira.FieldSchema{Type: "option"}},
    }
}

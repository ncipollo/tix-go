package github

import (
	"errors"
	"github.com/google/go-github/v29/github"
	"github.com/stretchr/testify/assert"
	"testing"
	"tix/ticket"
)

func TestFields_Assignees_SingleAssignees_CorrectType(t *testing.T) {
	rawFields := map[string]interface{}{KeyAssignee: "me"}
	fields, _, _, _ := testFieldsSetup(rawFields)

	result := fields.Assignees()

	expected := &[]string{"me"}
	assert.Equal(t, expected, result)
}

func TestFields_Assignees_SingleAssignees_WrongType(t *testing.T) {
	rawFields := map[string]interface{}{KeyAssignee: 42}
	fields, _, _, _ := testFieldsSetup(rawFields)

	result := fields.Assignees()

	assert.Nil(t, result)
}

func TestFields_Assignees_MultipleAssignees_CorrectType(t *testing.T) {
	rawFields := map[string]interface{}{KeyAssignees: []interface{}{"me", "myself", "I"}}
	fields, _, _, _ := testFieldsSetup(rawFields)

	result := fields.Assignees()

	expected := &[]string{"me", "myself", "I"}
	assert.Equal(t, expected, result)
}

func TestFields_Assignees_MultipleAssignees_WrongType(t *testing.T) {
	rawFields := map[string]interface{}{KeyAssignees: 42}
	fields, _, _, _ := testFieldsSetup(rawFields)

	result := fields.Assignees()

	assert.Nil(t, result)
}

func TestFields_IssueColumn_ApiError(t *testing.T) {
	number := 1
	id := int64(number)
	name := "todo"
	rawFields := map[string]interface{}{KeyColumn: name}
	fields, _, _, api := testFieldsSetup(rawFields)
	err := errors.New("project error")
	api.On("ListRepoProjects").Return(nil, err)

	result, err := fields.IssueColumn(id)

	assert.Nil(t, result)
	assert.Error(t, err)
}

func TestFields_IssueColumn_CorrectType(t *testing.T) {
	number := 1
	id := int64(number)
	name := "todo"
	column := &github.ProjectColumn{ID: &id, Name: &name}
	rawFields := map[string]interface{}{KeyColumn: name}
	fields, cache, _, _ := testFieldsSetup(rawFields)
	cache.Project.AddProject(&github.Project{ID: &id, Number: &number})
	columnCache, _ := cache.Project.ColumnCacheById(id)
	columnCache.AddColumn(column)

	result, err := fields.IssueColumn(id)

	assert.Equal(t, column, result)
	assert.NoError(t, err)
}

func TestFields_IssueColumn_MissingProject(t *testing.T) {
	number := 1
	id := int64(number)
	name := "todo"
	rawFields := map[string]interface{}{KeyColumn: name}
	fields, _, _, api := testFieldsSetup(rawFields)
	api.On("ListRepoProjects").Return([]*github.Project{}, nil)

	result, err := fields.IssueColumn(id)

	assert.Nil(t, result)
	assert.Error(t, err)
}

func TestFields_IssueColumn_WrongType(t *testing.T) {
	number := 1
	id := int64(number)
	name := "To do"
	column := &github.ProjectColumn{ID: &id, Name: &name}
	rawFields := map[string]interface{}{KeyColumn: nil}
	fields, cache, _, _ := testFieldsSetup(rawFields)
	cache.Project.AddProject(&github.Project{ID: &id, Number: &number})
	columnCache, _ := cache.Project.ColumnCacheById(id)
	columnCache.AddColumn(column)

	result, err := fields.IssueColumn(1)

	assert.Equal(t, column, result)
	assert.NoError(t, err)
}

func TestFields_Labels_CorrectType(t *testing.T) {
	rawFields := map[string]interface{}{KeyLabels: []interface{}{"label1", "label2"}}
	fields, _, _, _ := testFieldsSetup(rawFields)

	result := fields.Labels()

	expected := &[]string{"label1", "label2"}
	assert.Equal(t, expected, result)
}

func TestFields_Labels_WrongType(t *testing.T) {
	rawFields := map[string]interface{}{KeyLabels: 42}
	fields, _, _, _ := testFieldsSetup(rawFields)

	result := fields.Labels()

	assert.Nil(t, result)
}

func TestFields_Milestone_CorrectType(t *testing.T) {
	id := int64(1)
	name := "milestone1"
	milestone := &github.Milestone{ID: &id, Title: &name}
	rawFields := map[string]interface{}{KeyMilestone: name}
	fields, cache, _, _ := testFieldsSetup(rawFields)
	cache.Milestone.AddMilestone(milestone)

	result, err := fields.Milestone()

	assert.Equal(t, milestone, result)
	assert.NoError(t, err)
}

func TestFields_Milestone_WrongType(t *testing.T) {
	rawFields := map[string]interface{}{KeyMilestone: 42}
	fields, _, _, _ := testFieldsSetup(rawFields)

	result, err := fields.Milestone()

	assert.Nil(t, result)
	assert.Nil(t, err)
}

func TestFields_Project_ApiError(t *testing.T) {
	rawId := 1
	rawFields := map[string]interface{}{KeyProject: rawId}
	fields, _, _, api := testFieldsSetup(rawFields)
	err := errors.New("project error")
	api.On("ListRepoProjects").Return(nil, err)

	result, err := fields.Project()

	assert.Nil(t, result)
	assert.Error(t, err)
}

func TestFields_Project_CorrectType(t *testing.T) {
	number := 1
	id := int64(number)
	rawId := 1
	rawFields := map[string]interface{}{KeyProject: rawId}
	fields, cache, _, _ := testFieldsSetup(rawFields)
	cache.Project.AddProject(&github.Project{ID: &id, Number: &number})

	result, err := fields.Project()

	expectedId := int64(rawId)
	expected := &github.Project{ID: &expectedId, Number: &number}
	assert.Equal(t, expected, result)
	assert.NoError(t, err)
}

func TestFields_Project_MissingProject(t *testing.T) {
	rawId := 1
	rawFields := map[string]interface{}{KeyProject: rawId}
	fields, _, _, api := testFieldsSetup(rawFields)
	api.On("ListRepoProjects").Return([]*github.Project{}, nil)

	result, err := fields.Project()

	assert.Nil(t, result)
	assert.Error(t, err)
}

func TestFields_Project_WrongType(t *testing.T) {
	rawFields := map[string]interface{}{KeyProject: "wrong"}
	fields, _, _, _ := testFieldsSetup(rawFields)

	result, err := fields.Project()

	assert.Nil(t, result)
	assert.NoError(t, err)
}

func TestFields_ProjectColumns_CorrectType(t *testing.T) {
	rawFields := map[string]interface{}{KeyColumns: []interface{}{"col1", "col2"}}
	fields, _, _, _ := testFieldsSetup(rawFields)

	result := fields.ProjectColumns()

	expected := []string{"col1", "col2"}
	assert.Equal(t, expected, result)
}

func TestFields_ProjectColumns_WrongType(t *testing.T) {
	rawFields := map[string]interface{}{KeyColumns: 42}
	fields, _, _, _ := testFieldsSetup(rawFields)

	result := fields.ProjectColumns()

	expected := []string{"To do", "In progress", "Done"}
	assert.Equal(t, expected, result)
}

func testFieldsSetup(rawFields map[string]interface{}) (*Fields, *Cache, *ticket.Ticket, *mockApi) {
	api := newMockApi()
	cache := NewCache(api)
	testTicket := ticketForFieldTest(rawFields)
	fields := NewFields(cache, testTicket)

	return fields, cache, testTicket, api
}

func ticketForFieldTest(rawFields map[string]interface{}) *ticket.Ticket {
	testTicket := ticket.NewTicket()
	testTicket.AddFieldsForTicketSystem(rawFields, "github")
	return testTicket
}

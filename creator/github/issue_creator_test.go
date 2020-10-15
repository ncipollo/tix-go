package github

import (
	"errors"
	"github.com/google/go-github/v29/github"
	"github.com/stretchr/testify/assert"
	"testing"
	"tix/render"
	"tix/ticket"
	"tix/ticket/body"
)

func TestIssueCreator_CreateIssue_NoProject(t *testing.T) {
	issueCreator, api, _, options := setupIssueCreatorTest()
	issueTicket := createIssueCreatorTicket()
	issue := createIssueCreatorIssue()
	issueRequest := options.Issue(issueTicket, nil, nil, nil)

	api.On("CreateIssue", issueRequest).Return(issue, nil)

	result, err := issueCreator.CreateIssue(issueTicket, nil)

	assert.Equal(t, issue, result)
	assert.NoError(t, err)
}

func TestIssueCreator_CreateIssue_WithAddToProject(t *testing.T) {
	issueCreator, api, cache, options := setupIssueCreatorTest()
	issueTicket := createIssueCreatorTicket()
	issue := createIssueCreatorIssue()
	issueRequest := options.Issue(issueTicket, nil, nil, nil)
	project := createIssueCreatorProject()
	card := options.IssueCard(issue)
	column := setupIssueCreatorColumn(cache, project)
	issueTicket.AddFieldsForTicketSystem(map[string]interface{}{"project": *project.Number}, "github")

	api.On("CreateIssue", issueRequest).Return(issue, nil)
	api.On("AddCardToProject", card, column).Return(&github.ProjectCard{}, nil)

	result, err := issueCreator.CreateIssue(issueTicket, nil)

	assert.Equal(t, issue, result)
	assert.NoError(t, err)
}

func TestIssueCreator_CreateIssue_WithParentProject_Error_CardFailure(t *testing.T) {
	issueCreator, api, cache, options := setupIssueCreatorTest()
	issueTicket := createIssueCreatorTicket()
	issue := createIssueCreatorIssue()
	issueRequest := options.Issue(issueTicket, nil, nil, nil)
	project := createIssueCreatorProject()
	card := options.IssueCard(issue)
	column := setupIssueCreatorColumn(cache, project)

	err := errors.New("card failure")
	api.On("CreateIssue", issueRequest).Return(issue, nil)
	api.On("AddCardToProject", card, column).Return(nil, err)

	result, err := issueCreator.CreateIssue(issueTicket, project)

	assert.Nil(t, result)
	assert.Error(t, err)
}

func TestIssueCreator_CreateIssue_WithParentProject_Error_IssueFailure(t *testing.T) {
	issueCreator, api, _, options := setupIssueCreatorTest()
	issueTicket := createIssueCreatorTicket()
	issue := createIssueCreatorIssue()
	issueRequest := options.Issue(issueTicket, nil, nil, nil)
	project := createIssueCreatorProject()

	err := errors.New("issue failure")
	api.On("CreateIssue", issueRequest).Return(issue, err)

	result, err := issueCreator.CreateIssue(issueTicket, project)

	assert.Nil(t, result)
	assert.Error(t, err)
}

func TestIssueCreator_CreateIssue_WithParentProject_Success(t *testing.T) {
	issueCreator, api, cache, options := setupIssueCreatorTest()
	issueTicket := createIssueCreatorTicket()
	issue := createIssueCreatorIssue()
	issueRequest := options.Issue(issueTicket, nil, nil, nil)
	project := createIssueCreatorProject()
	card := options.IssueCard(issue)
	column := setupIssueCreatorColumn(cache, project)

	api.On("CreateIssue", issueRequest).Return(issue, nil)
	api.On("AddCardToProject", card, column).Return(&github.ProjectCard{}, nil)

	result, err := issueCreator.CreateIssue(issueTicket, project)

	assert.Equal(t, issue, result)
	assert.NoError(t, err)
}

func setupIssueCreatorTest() (IssueCreator, *mockApi, *Cache, *Options) {
	api := newMockApi()
	cache := NewCache(api)
	renderer := render.NewGithubBodyRenderer()
	options := NewOptions(renderer)
	projectCreator := NewIssueCreator(api, cache, options)

	return projectCreator, api, cache, options
}

func createIssueCreatorIssue() *github.Issue {
	number := 1
	id := int64(2)
	return &github.Issue{
		ID:     &id,
		Number: &number,
	}
}

func createIssueCreatorProject() *github.Project {
	number := 1
	id := int64(2)
	return &github.Project{
		ID:     &id,
		Number: &number,
	}
}

func createIssueCreatorTicket() *ticket.Ticket {
	newTicket := ticket.NewTicket()
	newTicket.Title = "Issue Title"
	newTicket.AddBodySegment(body.NewTextSegment("body"))

	return newTicket
}

func setupIssueCreatorColumn(cache *Cache,project *github.Project) *github.ProjectColumn {
	columnId := int64(3)
	columnName := "To do"
	column := github.ProjectColumn{ID: &columnId, Name: &columnName}
	cache.Project.AddProject(project)
	columnCache, _ := cache.Project.ColumnCacheById(*project.ID)
	columnCache.AddColumn(&column)

	return &column
}
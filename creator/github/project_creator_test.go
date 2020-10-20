package github

import (
	"errors"
	"github.com/google/go-github/v29/github"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"tix/render"
	"tix/ticket"
	"tix/ticket/body"
)

func TestProjectCreator_CreateProject_Error_ColumnFailure(t *testing.T) {
	columnNames := []string{"To do", "In progress", "Done"}
	projectCreator, api, _, options := setupProjectCreatorTest()
	projectTicket := createProjectCreatorTicket()
	project := createProjectCreatorProject()
	api.On("CreateProject", options.Project(projectTicket)).Return(project, nil)
	err := errors.New("column error")
	for _, name := range columnNames {
		api.On("CreateProjectColumn", options.ProjectColumn(name), project).Return(nil, err)
	}

	result, err := projectCreator.CreateProject(projectTicket)

	assert.Nil(t, result)
	assert.Error(t, err)
}

func TestProjectCreator_CreateProject_Error_ProjectFailure(t *testing.T) {
	projectCreator, api, _, options := setupProjectCreatorTest()
	projectTicket := createProjectCreatorTicket()
	err := errors.New("project error")
	api.On("CreateProject", options.Project(projectTicket)).Return(nil, err)

	result, err := projectCreator.CreateProject(projectTicket)

	assert.Nil(t, result)
	assert.Error(t, err)
}

func TestProjectCreator_CreateProject_Success(t *testing.T) {
	columnNames := []string{"To do", "In progress", "Done"}
	projectCreator, api, cache, options := setupProjectCreatorTest()
	projectTicket := createProjectCreatorTicket()
	project := createProjectCreatorProject()
	api.On("CreateProject", options.Project(projectTicket)).Return(project, nil)
	for _, name := range columnNames {
		// We need to avoid passing in the iterator value since it is a point which mutates. All values end up as "done"
		// if we don't create a copy.
		columnName := name
		api.On(
			"CreateProjectColumn",
			options.ProjectColumn(name),
			project,
		).Return(&github.ProjectColumn{Name: &columnName}, nil)
	}

	result, err := projectCreator.CreateProject(projectTicket)

	for _, name := range columnNames {
		columnCache, columnErr := cache.Project.ColumnCacheById(*project.ID)
		assert.NoError(t, columnErr)

		expected := &github.ProjectColumn{Name: &name}
		column, columnErr := columnCache.GetByName(name)
		assert.NoError(t, columnErr)
		assert.Equal(t, expected, column)
	}
	assert.Equal(t, project, result)
	assert.NoError(t, err)
}

func TestProjectCreator_UpdateProject_ApiError(t *testing.T) {
	projectCreator, api, cache, options := setupProjectCreatorTest()
	projectTicket := createProjectCreatorTicket()
	project := createProjectCreatorProject()
	updateKey := strconv.Itoa(*project.Number)
	cache.Project.AddProject(project)
	err := errors.New("api error")
	api.On("UpdateProject", project, options.Project(projectTicket)).Return(nil, err)

	result, err := projectCreator.UpdateProject(projectTicket, updateKey)

	assert.Nil(t, result)
	assert.Error(t, err)
}

func TestProjectCreator_UpdateProject_NonNumericKeyError(t *testing.T) {
	projectCreator, _, cache, _ := setupProjectCreatorTest()
	projectTicket := createProjectCreatorTicket()
	project := createProjectCreatorProject()
	cache.Project.AddProject(project)

	result, err := projectCreator.UpdateProject(projectTicket, "barf")

	assert.Nil(t, result)
	assert.Error(t, err)
}

func TestProjectCreator_UpdateProject_NoProject(t *testing.T) {
	projectCreator, api, _, _ := setupProjectCreatorTest()
	projectTicket := createProjectCreatorTicket()
	project := createProjectCreatorProject()
	updateKey := strconv.Itoa(*project.Number)
	api.On("ListRepoProjects").Return([]*github.Project{}, nil)

	result, err := projectCreator.UpdateProject(projectTicket, updateKey)

	assert.Nil(t, result)
	assert.Error(t, err)
}

func TestProjectCreator_UpdateProject_Success(t *testing.T) {
	projectCreator, api, cache, options := setupProjectCreatorTest()
	projectTicket := createProjectCreatorTicket()
	project := createProjectCreatorProject()
	updateKey := strconv.Itoa(*project.Number)
	updatedNumber := 42
	updatedProject := createProjectCreatorProject()
	updatedProject.Number = &updatedNumber
	cache.Project.AddProject(project)
	api.On("UpdateProject", project, options.Project(projectTicket)).Return(updatedProject, nil)

	result, err := projectCreator.UpdateProject(projectTicket, updateKey)
	cachedProject, _ := cache.Project.ProjectByNumber(42)

	assert.Equal(t, updatedProject, result)
	assert.Equal(t, updatedProject, cachedProject)
	assert.NoError(t, err)
}

func setupProjectCreatorTest() (ProjectCreator, *mockApi, *Cache, *Options) {
	api := newMockApi()
	cache := NewCache(api)
	renderer := render.NewGithubBodyRenderer()
	options := NewOptions(renderer)
	projectCreator := NewProjectCreator(api, cache, options)

	return projectCreator, api, cache, options
}

func createProjectCreatorProject() *github.Project {
	number := 1
	id := int64(2)
	return &github.Project{
		ID:     &id,
		Number: &number,
	}
}

func createProjectCreatorTicket() *ticket.Ticket {
	newTicket := ticket.NewTicket()
	newTicket.Title = "Project Title"
	newTicket.AddBodySegment(body.NewTextSegment("body"))

	return newTicket
}

package github

import (
	"github.com/google/go-github/v29/github"
	"github.com/stretchr/testify/mock"
)

type mockApi struct {
	mock.Mock
}

func newMockApi() *mockApi {
	return &mockApi{}
}

func (m *mockApi) AddCardToProject(cardOptions *github.ProjectCardOptions,
	column *github.ProjectColumn) (*github.ProjectCard, error) {
	args := m.Called(cardOptions, column)
	result := args.Get(0)
	err := args.Error(1)
	if result != nil {
		return result.(*github.ProjectCard), err
	}
	return nil, err
}

func (m *mockApi) CreateIssue(issueRequest *github.IssueRequest) (*github.Issue, error) {
	args := m.Called(issueRequest)
	result := args.Get(0)
	err := args.Error(1)
	if result != nil {
		return result.(*github.Issue), err
	}
	return nil, err
}

func (m *mockApi) CreateMilestone(milestone *github.Milestone) (*github.Milestone, error) {
	args := m.Called(milestone)
	result := args.Get(0)
	err := args.Error(1)
	if result != nil {
		return result.(*github.Milestone), err
	}
	return nil, err
}

func (m *mockApi) CreateProject(projectOptions *github.ProjectOptions) (*github.Project, error) {
	args := m.Called(projectOptions)
	result := args.Get(0)
	err := args.Error(1)
	if result != nil {
		return result.(*github.Project), err
	}
	return nil, err
}

func (m *mockApi) CreateProjectColumn(columnOptions *github.ProjectColumnOptions,
	project *github.Project) (*github.ProjectColumn, error) {
	args := m.Called(columnOptions, project)
	result := args.Get(0)
	err := args.Error(1)
	if result != nil {
		return result.(*github.ProjectColumn), err
	}
	return nil, err
}

func (m *mockApi)  ListProjectColumns(project *github.Project) ([]*github.ProjectColumn, error) {
	args := m.Called(project)
	result := args.Get(0)
	err := args.Error(1)
	if result != nil {
		return result.([]*github.ProjectColumn), err
	}
	return nil, err
}

func (m *mockApi) ListRepoProjects() ([]*github.Project, error) {
	args := m.Called()
	result := args.Get(0)
	err := args.Error(1)
	if result != nil {
		return result.([]*github.Project), err
	}
	return nil, err
}

func (m *mockApi) ListMilestones() ([]*github.Milestone, error) {
	args := m.Called()
	result := args.Get(0)
	err := args.Error(1)
	if result != nil {
		return result.([]*github.Milestone), err
	}
	return nil, err
}

func (m *mockApi) UpdateIssue(issue *github.Issue, issueRequest *github.IssueRequest) (*github.Issue, error) {
	args := m.Called(issue, issueRequest)
	result := args.Get(0)
	err := args.Error(1)
	if result != nil {
		return result.(*github.Issue), err
	}
	return nil, err
}

func (m *mockApi) UpdateProject(project *github.Project, projectOptions *github.ProjectOptions) (*github.Project, error) {
	args := m.Called(project, projectOptions)
	result := args.Get(0)
	err := args.Error(1)
	if result != nil {
		return result.(*github.Project), err
	}
	return nil, err
}

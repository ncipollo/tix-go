package github

import (
	"context"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type Api interface {
	AddCardToProject(cardOptions *github.ProjectCardOptions, column *github.ProjectColumn) (*github.ProjectCard, error)
	CreateIssue(issueRequest *github.IssueRequest) (*github.Issue, error)
	CreateProject(projectOptions *github.ProjectOptions) (*github.Project, error)
	CreateProjectColumn(columnOptions *github.ProjectColumnOptions, project *github.Project) (*github.ProjectColumn, error)
}

type githubApi struct {
	client *github.Client
	ctx    context.Context
	owner  string
	repo   string
}

func NewGithubApi(token string, owner string, repo string) *githubApi {
	ctx := context.Background()
	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	httpClient := oauth2.NewClient(ctx, tokenSource)

	githubClient := github.NewClient(httpClient)

	return &githubApi{client: githubClient, ctx: ctx, owner: owner, repo: repo}
}

func (g githubApi) AddCardToProject(cardOptions *github.ProjectCardOptions,
	column *github.ProjectColumn) (*github.ProjectCard, error) {
	card, _, err :=  g.client.Projects.CreateProjectCard(g.ctx, *column.ID, cardOptions)
	return card, err
}

func (g githubApi) CreateIssue(issueRequest *github.IssueRequest) (*github.Issue, error) {
	issue, _, err := g.client.Issues.Create(g.ctx, g.owner, g.repo, issueRequest)
	return issue, err
}

func (g githubApi) CreateProject(projectOptions *github.ProjectOptions) (*github.Project, error) {
	project, _, err := g.client.Repositories.CreateProject(g.ctx, g.owner, g.repo, projectOptions)
	return project, err
}

func (g githubApi) CreateProjectColumn(columnOptions *github.ProjectColumnOptions,
	project *github.Project) (*github.ProjectColumn, error) {
	column, _, err := g.client.Projects.CreateProjectColumn(g.ctx, *project.ID, columnOptions)
	return column, err
}

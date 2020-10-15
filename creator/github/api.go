package github

import (
	"context"
	"github.com/google/go-github/v29/github"
	"golang.org/x/oauth2"
)

type Api interface {
	AddCardToProject(cardOptions *github.ProjectCardOptions, column *github.ProjectColumn) (*github.ProjectCard, error)
	CreateIssue(issueRequest *github.IssueRequest) (*github.Issue, error)
	CreateMilestone(milestone *github.Milestone) (*github.Milestone, error)
	CreateProject(projectOptions *github.ProjectOptions) (*github.Project, error)
	CreateProjectColumn(columnOptions *github.ProjectColumnOptions, project *github.Project) (*github.ProjectColumn, error)
	ListProjectColumns(project *github.Project) ([]*github.ProjectColumn, error)
	ListRepoProjects() ([]*github.Project, error)
	ListMilestones() ([]*github.Milestone, error)
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
	card, _, err := g.client.Projects.CreateProjectCard(g.ctx, *column.ID, cardOptions)
	return card, err
}

func (g githubApi) CreateIssue(issueRequest *github.IssueRequest) (*github.Issue, error) {
	issue, _, err := g.client.Issues.Create(g.ctx, g.owner, g.repo, issueRequest)
	return issue, err
}

func (g githubApi) CreateMilestone(milestone *github.Milestone) (*github.Milestone, error) {
	milestone, _, err := g.client.Issues.CreateMilestone(g.ctx, g.owner, g.repo, milestone)
	return milestone, err
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

func (g githubApi) ListProjectColumns(project *github.Project) ([]*github.ProjectColumn, error) {
	var allColumns []*github.ProjectColumn
	opt := &github.ListOptions{
		PerPage: 50,
	}
	for {
		columns, resp, err := g.client.Projects.ListProjectColumns(g.ctx, *project.ID, opt)
		if err != nil {
			return nil, err
		}
		allColumns = append(allColumns, columns...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	return allColumns, nil
}

func (g githubApi) ListRepoProjects() ([]*github.Project, error) {
	var allProjects []*github.Project
	opt := &github.ProjectListOptions{
		ListOptions: github.ListOptions{PerPage: 50},
	}
	for {
		projects, resp, err := g.client.Repositories.ListProjects(g.ctx, g.owner, g.repo, opt)
		if err != nil {
			return nil, err
		}
		allProjects = append(allProjects, projects...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	return allProjects, nil
}

func (g githubApi) ListMilestones() ([]*github.Milestone, error) {
	var allMilestones []*github.Milestone
	opt := &github.MilestoneListOptions{
		ListOptions: github.ListOptions{PerPage: 50},
	}
	for {
		milestones, resp, err := g.client.Issues.ListMilestones(g.ctx, g.owner, g.repo, opt)
		if err != nil {
			return nil, err
		}
		allMilestones = append(allMilestones, milestones...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	return allMilestones, nil
}

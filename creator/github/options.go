package github

import (
	"github.com/google/go-github/v29/github"
	"tix/render"
	"tix/ticket"
)

type Options struct {
	renderer render.BodyRenderer
}

func NewOptions(renderer render.BodyRenderer) *Options {
	return &Options{renderer: renderer}
}

func (o Options) Project(ticket *ticket.Ticket) *github.ProjectOptions {
	body := render.Ticket(ticket, o.renderer)
	return &github.ProjectOptions{
		Name: &ticket.Title,
		Body: &body,
	}
}

func (o Options) ProjectColumn(name string) *github.ProjectColumnOptions {
	return &github.ProjectColumnOptions{Name: name}
}

func (o Options) Issue(ticket *ticket.Ticket,
	assignees *[]string,
	labels *[]string,
	milestoneId *int) *github.IssueRequest {
	body := render.Ticket(ticket, o.renderer)
	return &github.IssueRequest{
		Title:     &ticket.Title,
		Body:      &body,
		Labels:    labels,
		Assignee:  nil,
		State:     nil,
		Milestone: milestoneId,
		Assignees: assignees,
	}
}

func (o Options) IssueCard(issue *github.Issue) *github.ProjectCardOptions {
	return &github.ProjectCardOptions{
		ContentID:   *issue.ID,
		ContentType: "Issue",
	}
}

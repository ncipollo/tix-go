package github

import (
	"github.com/google/go-github/v29/github"
	"tix/ticket"
)

type IssueCreator interface {
	CreateIssue(ticket *ticket.Ticket, parentProject *github.Project) (*github.Issue, error)
}

type ApiIssueCreator struct {
	api     Api
	cache   *Cache
	options *Options
}

func NewIssueCreator(api Api, cache *Cache, options *Options) IssueCreator {
	return &ApiIssueCreator{
		api:     api,
		cache:   cache,
		options: options,
	}
}

func (c ApiIssueCreator) CreateIssue(ticket *ticket.Ticket, parentProject *github.Project) (*github.Issue, error) {
	fields := c.createFields(ticket)
	issue, err := c.createIssue(ticket, fields)
	if err != nil {
		return nil, err
	}

	project, err := c.projectToUse(parentProject, fields)

	if err != nil {
		return nil, err
	}

	if project != nil {
		err = c.addIssueToProject(issue, project, fields)
		if err != nil {
			return nil, err
		}
	}

	return issue, nil
}

func (c ApiIssueCreator) createFields(ticket *ticket.Ticket) *Fields {
	return NewFields(c.cache, ticket)
}

func (c ApiIssueCreator) createIssue(ticket *ticket.Ticket, fields *Fields) (*github.Issue, error) {
	assignees := fields.Assignees()
	labels := fields.Labels()
	milestone, err := fields.Milestone()
	if err != nil {
		return nil, err
	}
	var milestoneId *int
	if milestone != nil && milestone.Number != nil {
		value := *milestone.Number
		milestoneId = &value
	}

	request := c.options.Issue(ticket, assignees, labels, milestoneId)

	return c.api.CreateIssue(request)
}

func (c ApiIssueCreator) addIssueToProject(issue *github.Issue, project *github.Project, fields *Fields) error {
	column, err := fields.IssueColumn(*project.ID)
	if err != nil {
		return err
	}

	card := c.options.IssueCard(issue)

	_, err = c.api.AddCardToProject(card, column)

	return err
}

func (c ApiIssueCreator) projectToUse(parentProject *github.Project, fields *Fields) (*github.Project, error) {
	if parentProject != nil {
		return parentProject, nil
	}
	return fields.Project()
}

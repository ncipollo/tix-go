package jira

import (
	"github.com/andygrunwald/go-jira"
	"tix/render"
	"tix/ticket"
)

type Issues struct {
	jiraFields []jira.Field
	renderer   render.BodyRenderer
}

func NewIssues(jiraFields []jira.Field, renderer render.BodyRenderer) *Issues {
	return &Issues{jiraFields: jiraFields, renderer: renderer}
}

func (i *Issues) FromTicket(ticket *ticket.Ticket, parentTicketId string, ticketLevel int) *jira.Issue {
	switch ticketLevel {
	case 1:
		return i.epic(ticket)
	case 2:
		return i.story(ticket, parentTicketId)
	default:
		return i.subtask(ticket, parentTicketId)
	}
}

func (i *Issues) epic(ticket *ticket.Ticket) *jira.Issue {
	description := i.renderBody(ticket)
	issueFields := NewIssueFields(i.jiraFields, ticket)
	// Add epic name if missing
	unknowns := issueFields.Unknowns()
	issueFields.AddDefaultEpicName(unknowns, ticket.Title)

	return &jira.Issue{
		Fields: &jira.IssueFields{
			Components:  issueFields.Components(),
			Description: description,
			Labels:      issueFields.Labels(),
			Type:        issueFields.EpicType(),
			Project:     issueFields.Project(),
			Summary:     ticket.Title,
			Unknowns:    unknowns,
		},
	}
}

func (i *Issues) story(ticket *ticket.Ticket, parentTicketId string) *jira.Issue {
	description := i.renderBody(ticket)
	issueFields := NewIssueFields(i.jiraFields, ticket)
	// Add epic link to unknowns
	unknowns := issueFields.Unknowns()
	unknowns[issueFields.EpicLinkKey()] = parentTicketId

	return &jira.Issue{
		Fields: &jira.IssueFields{
			Components:  issueFields.Components(),
			Description: description,
			Labels:      issueFields.Labels(),
			Type:        issueFields.IssueType(),
			Project:     issueFields.Project(),
			Summary:     ticket.Title,
			Unknowns:    unknowns,
		},
	}
}

func (i *Issues) subtask(ticket *ticket.Ticket, parentTicketId string) *jira.Issue {
	description := i.renderBody(ticket)
	issueFields := NewIssueFields(i.jiraFields, ticket)

	return &jira.Issue{
		Fields: &jira.IssueFields{
			Components:  issueFields.Components(),
			Description: description,
			Labels:      issueFields.Labels(),
			Type:        issueFields.EpicType(),
			Project:     issueFields.Project(),
			Parent:      &jira.Parent{Key: parentTicketId},
			Summary:     ticket.Title,
			Unknowns:    issueFields.Unknowns(),
		},
	}
}

func (i *Issues) renderBody(ticket *ticket.Ticket) string {
	return render.Ticket(ticket, i.renderer)
}

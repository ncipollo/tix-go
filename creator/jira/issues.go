package jira

import (
	"github.com/andygrunwald/go-jira"
	"tix/render"
	"tix/ticket"
)

type Issues struct {
	renderer    render.BodyRenderer
	issueFields IssueFields
}

func NewIssues(renderer render.BodyRenderer, issueFields IssueFields) *Issues {
	return &Issues{renderer: renderer, issueFields: issueFields}
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
	return &jira.Issue{
		Fields: &jira.IssueFields{
			Components:  i.issueFields.Components(),
			Description: description,
			Labels:      i.issueFields.Labels(),
			Type:        i.issueFields.EpicType(),
			Project:     i.issueFields.Project(),
			Summary:     ticket.Title,
			Unknowns:    i.issueFields.Unknowns(),
		},
	}
}

func (i *Issues) story(ticket *ticket.Ticket, parentTicketId string) *jira.Issue {
	description := i.renderBody(ticket)
	// Add epic link to unknowns
	unknowns := i.issueFields.Unknowns()
	unknowns[i.issueFields.EpicLinkKey()] = parentTicketId

	return &jira.Issue{
		Fields: &jira.IssueFields{
			Components:  i.issueFields.Components(),
			Description: description,
			Labels:      i.issueFields.Labels(),
			Type:        i.issueFields.EpicType(),
			Project:     i.issueFields.Project(),
			Summary:     ticket.Title,
			Unknowns:    unknowns,
		},
	}
}

func (i *Issues) subtask(ticket *ticket.Ticket, parentTicketId string) *jira.Issue {
	description := i.renderBody(ticket)

	return &jira.Issue{
		Fields: &jira.IssueFields{
			Components:  i.issueFields.Components(),
			Description: description,
			Labels:      i.issueFields.Labels(),
			Type:        i.issueFields.EpicType(),
			Project:     i.issueFields.Project(),
			Parent:      &jira.Parent{Key: parentTicketId},
			Summary:     ticket.Title,
			Unknowns:    i.issueFields.Unknowns(),
		},
	}
}

func (i *Issues) renderBody(ticket *ticket.Ticket) string {
	return render.Ticket(ticket, i.renderer)
}

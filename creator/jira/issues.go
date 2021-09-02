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

func (i *Issues) FromTicket(ticket *ticket.Ticket, parentIssue *jira.Issue, ticketLevel int) *jira.Issue {
    switch ticketLevel {
    case 0:
        return i.epic(ticket)
    case 1:
        return i.story(ticket, parentIssue)
    default:
        return i.task(ticket, parentIssue)
    }
}

func (i *Issues) epic(ticket *ticket.Ticket) *jira.Issue {
    description := i.renderBody(ticket)
    issueFields := NewIssueFields(i.jiraFields, ticket)
    // Add epic name if missing
    unknowns := issueFields.Unknowns()

    return &jira.Issue{
        Fields: &jira.IssueFields{
            AffectsVersions: issueFields.AffectsVersions(),
            Components:      issueFields.Components(),
            Description:     description,
            FixVersions:     issueFields.FixVersions(),
            Labels:          issueFields.Labels(),
            Type:            issueFields.EpicType(),
            Project:         issueFields.Project(),
            Summary:         ticket.Title,
            Unknowns:        unknowns,
        },
    }
}

func (i *Issues) story(ticket *ticket.Ticket, parentIssue *jira.Issue) *jira.Issue {
    description := i.renderBody(ticket)
    issueFields := NewIssueFields(i.jiraFields, ticket)
    unknowns := issueFields.Unknowns()
	var parent *jira.Parent
    if parentIssue != nil {
        if issueFields.UseParent() {
            parent = &jira.Parent{Key: parentIssue.Key}
        } else  {
            // Add epic link to unknowns
            unknowns[issueFields.EpicLinkKey()] = parentIssue.Key
        }
    }

    return &jira.Issue{
        Fields: &jira.IssueFields{
            AffectsVersions: issueFields.AffectsVersions(),
            Components:      issueFields.Components(),
            Description:     description,
            FixVersions:     issueFields.FixVersions(),
            Labels:          issueFields.Labels(),
            Type:            issueFields.IssueType(),
            Parent:          parent,
            Project:         issueFields.Project(),
            Summary:         ticket.Title,
            Unknowns:        unknowns,
        },
    }
}

func (i *Issues) task(ticket *ticket.Ticket, parentIssue *jira.Issue) *jira.Issue {
    description := i.renderBody(ticket)
    issueFields := NewIssueFields(i.jiraFields, ticket)

    return &jira.Issue{
        Fields: &jira.IssueFields{
            AffectsVersions: issueFields.AffectsVersions(),
            Components:      issueFields.Components(),
            Description:     description,
            FixVersions:     issueFields.FixVersions(),
            Labels:          issueFields.Labels(),
            Type:            issueFields.TaskType(),
            Project:         issueFields.Project(),
            Parent:          &jira.Parent{ID: parentIssue.ID},
            Summary:         ticket.Title,
            Unknowns:        issueFields.Unknowns(),
        },
    }
}

func (i *Issues) renderBody(ticket *ticket.Ticket) string {
    return render.Ticket(ticket, i.renderer)
}

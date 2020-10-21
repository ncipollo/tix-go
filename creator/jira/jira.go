package jira

import (
	"github.com/andygrunwald/go-jira"
	"tix/creator/reporter"
	"tix/logger"
	"tix/render"
	"tix/ticket"
)

type Creator struct {
	api                 Api
	startingTicketLevel int
}

func NewCreatorWithEpics(api Api) *Creator {
	return &Creator{api, 0}
}

func NewCreatorWithoutEpics(api Api) *Creator {
	return &Creator{api, 1}
}

func (j Creator) CreateTickets(tickets []*ticket.Ticket) {
	issues, err := j.createIssues()
	if err != nil {
		logger.Error("%v", err)
	} else {
		j.createTicketsForLevel(tickets, issues, j.startingTicketLevel, nil)
	}
}

func (j Creator) createIssues() (*Issues, error) {
	fields, err := j.api.GetIssueFieldList()
	if err != nil {
		return nil, err
	}

	renderer := render.NewJiraBodyRenderer()
	return NewIssues(fields, renderer), nil
}

func (j Creator) createTicketsForLevel(tickets []*ticket.Ticket, issues *Issues, level int, parentIssue *jira.Issue) {
	for _, currentTicket := range tickets {
		var resultIssue *jira.Issue
		var err error

		updateKey := currentTicket.TicketUpdateKey("jira")
		issue := issues.FromTicket(currentTicket, parentIssue, level)

		if len(updateKey) > 0 {
			issue.Key = updateKey
			resultIssue, err = j.api.UpdateIssue(issue)
		} else {
			resultIssue, err = j.api.CreateIssue(issue)
		}

		if err != nil {
			reporter.ReportFailedTicket(err, j.startingTicketLevel, level)
		} else {
			reporter.ReportSuccessfulTicket(
				resultIssue.Key,
				j.startingTicketLevel, level,
				currentTicket.Title,
				updateKey,
			)
			j.createTicketsForLevel(currentTicket.Subtickets, issues, level+1, resultIssue)
		}
	}
}

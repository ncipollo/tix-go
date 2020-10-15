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
		issue := issues.FromTicket(currentTicket, parentIssue, level)
		resultIssue, err := j.api.CreateIssue(issue)
		if err != nil {
			reporter.ReportFailedTicketCreate(err, j.startingTicketLevel, level)
		} else {
			reporter.ReportSuccessfulTicketCreate(resultIssue.Key, j.startingTicketLevel, level, currentTicket.Title)
			j.createTicketsForLevel(currentTicket.Subtickets, issues, level+1, resultIssue)
		}
	}
}

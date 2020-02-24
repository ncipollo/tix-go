package jira

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
	"strings"
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
			j.reportFailedTicketCreate(err, level)
		} else {
			j.reportSuccessfulTicketCreate(resultIssue, level, currentTicket.Title)
			j.createTicketsForLevel(currentTicket.Subtickets, issues, level+1, resultIssue)
		}
	}
}

func (j Creator) reportFailedTicketCreate(err error, level int) {
	var builder strings.Builder
	for ii := j.startingTicketLevel; ii < level; ii++ {
		builder.WriteString("\t")
	}
	builder.WriteString("- ")
	builder.WriteString(err.Error())

	logger.Error(builder.String())
}

func (j Creator) reportSuccessfulTicketCreate(issue *jira.Issue, level int, title string) {
	var builder strings.Builder
	for ii := j.startingTicketLevel; ii < level; ii++ {
		builder.WriteString("\t")
	}
	message := fmt.Sprintf("- :tada: %v: %v created", issue.Key, title)
	builder.WriteString(message)

	logger.Message(builder.String())
}

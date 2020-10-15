package github

import (
	"fmt"
	"github.com/google/go-github/v29/github"
	"tix/creator/reporter"
	"tix/ticket"
)

type Creator struct {
	issueCreator        IssueCreator
	projectCreator      ProjectCreator
	startingTicketLevel int
}

func NewCreatorWithProjects(issueCreator IssueCreator, projectCreator ProjectCreator) *Creator {
	return &Creator{issueCreator: issueCreator, projectCreator: projectCreator, startingTicketLevel: 0}
}

func NewCreatorWithoutProjects(issueCreator IssueCreator, projectCreator ProjectCreator) *Creator {
	return &Creator{issueCreator: issueCreator, projectCreator: projectCreator, startingTicketLevel: 1}
}

func (c Creator) CreateTickets(tickets []*ticket.Ticket) {
	if c.hasProjects() {
		c.createProjects(tickets)
	} else {
		c.createIssues(tickets, nil)
	}
}

func (c Creator) createProjects(tickets []*ticket.Ticket) {
	for _, currentTicket := range tickets {
		project, err := c.projectCreator.CreateProject(currentTicket)
		if err != nil {
			reporter.ReportFailedTicketCreate(err, c.startingTicketLevel, 0)
		} else {
			key := fmt.Sprintf("%d", *project.Number)
			reporter.ReportSuccessfulTicketCreate(key, c.startingTicketLevel, 0, currentTicket.Title)
			c.createIssues(currentTicket.Subtickets, project)
		}
	}
}

func (c Creator) createIssues(tickets []*ticket.Ticket, project *github.Project) {
	for _, currentTicket := range tickets {
		issue, err := c.issueCreator.CreateIssue(currentTicket, project)
		if err != nil {
			reporter.ReportFailedTicketCreate(err, c.startingTicketLevel, 1)
		} else {
			key := fmt.Sprintf("%d", *issue.Number)
			reporter.ReportSuccessfulTicketCreate(key, c.startingTicketLevel, 1, currentTicket.Title)
		}
	}
}

func (c Creator) hasProjects() bool {
	return c.startingTicketLevel == 0
}

package github

import (
	"github.com/google/go-github/v29/github"
	"github.com/stretchr/testify/mock"
	"testing"
	"tix/ticket"
)

func TestCreator_CreateTickets_UpdateTickets(t *testing.T) {
	id := int64(1)
	number := 2
	ticketCreator, issueCreator, projectCreator := setupCreatorTests(false)
	tickets := ticketsForUpdater()
	project := &github.Project{ID: &id, Number: &number}
	issue := &github.Issue{ID: &id, Number: &number}
	projectCreator.On(
		"UpdateProject",
		mock.AnythingOfType("*ticket.Ticket"),
		mock.AnythingOfType("string"),
	).Return(project, nil)
	issueCreator.On(
		"UpdateIssue",
		mock.AnythingOfType("*ticket.Ticket"),
		mock.AnythingOfType("string"),
	).Return(issue, nil)

	ticketCreator.CreateTickets(tickets)
	projectCreator.AssertNumberOfCalls(t, "UpdateProject", 2)
	issueCreator.AssertNumberOfCalls(t, "UpdateIssue", 1)
}

func TestCreator_CreateTickets_WithOutProjects(t *testing.T) {
	id := int64(1)
	number := 2
	ticketCreator, issueCreator, projectCreator := setupCreatorTests(true)
	tickets := ticketsForCreator()
	issue := &github.Issue{ID: &id, Number: &number}
	issueCreator.On(
		"CreateIssue",
		mock.AnythingOfType("*ticket.Ticket"),
		mock.AnythingOfType("*github.Project"),
	).Return(issue, nil)

	ticketCreator.CreateTickets(tickets)
	projectCreator.AssertNotCalled(t, "CreateProject")
	issueCreator.AssertNumberOfCalls(t, "CreateIssue", 2)
}

func TestCreator_CreateTickets_WithProjects(t *testing.T) {
	id := int64(1)
	number := 2
	ticketCreator, issueCreator, projectCreator := setupCreatorTests(false)
	tickets := ticketsForCreator()
	project := &github.Project{ID: &id, Number: &number}
	issue := &github.Issue{ID: &id, Number: &number}
	projectCreator.On("CreateProject", mock.AnythingOfType("*ticket.Ticket")).Return(project, nil)
	issueCreator.On(
		"CreateIssue",
		mock.AnythingOfType("*ticket.Ticket"),
		mock.AnythingOfType("*github.Project"),
	).Return(issue, nil)

	ticketCreator.CreateTickets(tickets)

	projectCreator.AssertNumberOfCalls(t, "CreateProject", 2)
	issueCreator.AssertNumberOfCalls(t, "CreateIssue", 1)
}

func setupCreatorTests(noProjects bool) (*Creator, *mockIssueCreator, *mockProjectCreator) {
	issueCreator := &mockIssueCreator{}
	projectCreator := &mockProjectCreator{}

	if noProjects {
		return NewCreatorWithoutProjects(issueCreator, projectCreator), issueCreator, projectCreator
	} else {
		return NewCreatorWithProjects(issueCreator, projectCreator), issueCreator, projectCreator
	}
}

func ticketsForCreator() []*ticket.Ticket {
	ticket1 := ticket.NewTicket()
	ticket1.Title = "ticket 1"

	ticket2 := ticket.NewTicket()
	ticket2.Title = "ticket 1"

	subTicket := ticket.NewTicket()
	subTicket.Title = "subticket"
	ticket1.AddSubticket(subTicket)

	return []*ticket.Ticket{ticket1, ticket2}
}

func ticketsForUpdater() []*ticket.Ticket {
	ticket1 := ticket.NewTicketWithFields(map[string]interface{}{"update_ticket": "1"})
	ticket1.Title = "ticket 1"

	ticket2 := ticket.NewTicketWithFields(map[string]interface{}{"update_ticket": "2"})
	ticket2.Title = "ticket 1"

	subTicket := ticket.NewTicketWithFields(map[string]interface{}{"update_ticket": "3"})
	subTicket.Title = "subticket"
	ticket1.AddSubticket(subTicket)

	return []*ticket.Ticket{ticket1, ticket2}
}

type mockIssueCreator struct {
	mock.Mock
}

func (m *mockIssueCreator) CreateIssue(ticket *ticket.Ticket, parentProject *github.Project) (*github.Issue, error) {
	args := m.Called(ticket, parentProject)
	result := args.Get(0)
	err := args.Error(1)
	if result != nil {
		return result.(*github.Issue), err
	}
	return nil, err
}

func (m *mockIssueCreator) UpdateIssue(ticket *ticket.Ticket, updateKey string) (*github.Issue, error) {
	args := m.Called(ticket, updateKey)
	result := args.Get(0)
	err := args.Error(1)
	if result != nil {
		return result.(*github.Issue), err
	}
	return nil, err
}

type mockProjectCreator struct {
	mock.Mock
}

func (m *mockProjectCreator) CreateProject(ticket *ticket.Ticket) (*github.Project, error) {
	args := m.Called(ticket)
	result := args.Get(0)
	err := args.Error(1)
	if result != nil {
		return result.(*github.Project), err
	}
	return nil, err
}

func (m *mockProjectCreator) UpdateProject(ticket *ticket.Ticket, updateKey string) (*github.Project, error) {
	args := m.Called(ticket, updateKey)
	result := args.Get(0)
	err := args.Error(1)
	if result != nil {
		return result.(*github.Project), err
	}
	return nil, err
}

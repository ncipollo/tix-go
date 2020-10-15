package github

import (
	"github.com/google/go-github/v29/github"
	"github.com/stretchr/testify/assert"
	"testing"
	"tix/render"
	"tix/ticket"
	"tix/ticket/body"
)

func TestOptions_Issue(t *testing.T) {
	options := createTestOptions()
	testTicket := createOptionsTestTicket()
	assignees := []string{"aPerson"}
	labels := []string{"label1", "label2"}
	milestoneId := 42

	result := options.Issue(testTicket, &assignees, &labels, &milestoneId)

	expectedBody := "body"
	expected := &github.IssueRequest{
		Title:     &testTicket.Title,
		Body:      &expectedBody,
		Labels:    &labels,
		Milestone: &milestoneId,
		Assignees: &assignees,
	}
	assert.Equal(t, expected, result)
}

func TestOptions_IssueCard(t *testing.T) {
	options := createTestOptions()
	issueId := int64(1)
	issue := &github.Issue{ID: &issueId}

	result := options.IssueCard(issue)

	expected := &github.ProjectCardOptions{
		ContentID:   issueId,
		ContentType: "Issue",
	}
	assert.Equal(t, expected, result)
}

func TestOptions_Project(t *testing.T) {
	options := createTestOptions()
	testTicket := createOptionsTestTicket()

	result := options.Project(testTicket)

	expectedBody := "body"
	expected := &github.ProjectOptions{
		Name: &testTicket.Title,
		Body: &expectedBody,
	}
	assert.Equal(t, expected, result)
}

func TestOptions_ProjectColumn(t *testing.T) {
	options := createTestOptions()
	name := "column"

	result := options.ProjectColumn(name)

	expected := &github.ProjectColumnOptions{Name: name}
	assert.Equal(t, expected, result)
}

func createTestOptions() *Options {
	renderer := render.NewGithubBodyRenderer()
	return NewOptions(renderer)
}

func createOptionsTestTicket() *ticket.Ticket {
	testTicket := ticket.NewTicket()
	testTicket.Title = "title"
	testTicket.AddBodySegment(body.NewTextSegment("body"))

	return testTicket
}

package jira

import (
	"github.com/andygrunwald/go-jira"
	"tix/logger"
)

type Api interface {
	CreateIssue(issue *jira.Issue) (*jira.Issue, error)
	GetIssueFieldList() ([]jira.Field, error)
}

type jiraApi struct {
	client *jira.Client
}

func NewApi(userName string, apiToken string, baseUrl string) Api {
	tp := jira.BasicAuthTransport{
		Username: userName,
		Password: apiToken,
	}
	client, err := jira.NewClient(tp.Client(), baseUrl)
	if err != nil {
		logger.Error("Unable to create Jira client :scream: \n Error: %v", err)
		panic(err)
	}

	return &jiraApi{client: client}
}

func (j *jiraApi) CreateIssue(issue *jira.Issue) (*jira.Issue, error) {
	issue, response, err := j.client.Issue.Create(issue)
	if err != nil {
		return nil, j.generateError(err, response)
	}
	return issue, nil
}

func (j *jiraApi) GetIssueFieldList() ([]jira.Field, error) {
	panic("implement me")
}

func (j *jiraApi) generateError(original error, responose *jira.Response) error {
	panic("implement me")
}

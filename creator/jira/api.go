package jira

import (
	"errors"
	"fmt"
	"github.com/andygrunwald/go-jira"
	"io/ioutil"
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
		logger.Error(":scream: unable to create Jira client\n Error: %v", err)
		panic(err)
	}

	return &jiraApi{client: client}
}

func (j *jiraApi) CreateIssue(issue *jira.Issue) (*jira.Issue, error) {
	issue, response, err := j.client.Issue.Create(issue)
	if err != nil {
		return nil, j.generateError(":scream: unable to create jira issue", err, response)
	}
	return issue, nil
}

func (j *jiraApi) GetIssueFieldList() ([]jira.Field, error) {
	fields, response, err := j.client.Field.GetList()
	if err != nil {
		return nil, j.generateError(":scream: unable to fetch jira fields", err, response)
	}
	return fields, nil
}

func (j *jiraApi) generateError(preamble string, original error, response *jira.Response) error {
	body, _ := ioutil.ReadAll(response.Body)
	if body != nil {
		bodyStr := string(body)
		errStr := fmt.Sprintf("%v\nError: %v\nAPI Message: %v", preamble, original, bodyStr)
		return errors.New(errStr)
	} else {
		errStr := fmt.Sprintf("%v\nError: %v", preamble, original)
		return errors.New(errStr)
	}
}

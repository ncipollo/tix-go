package main

import (
	"github.com/andygrunwald/go-jira"
	"io/ioutil"
	"os"
	"strings"
	"tix/logger"
)

const md = `
# Epic Name

- List item 1
   - Nested Item
- List item 2
`

func main() {
	logger.Message("Test %s :scream:", "test")
}

func dontRun()  {
	envMap := make(map[string]string)
	for _, env := range os.Environ() {
		pair := strings.SplitN(env, "=", 2)
		envMap[pair[0]] = pair[1]
	}

	userName := envMap["AGENCY_BOT_USERNAME"]
	password := envMap["AGENCY_BOT_API_TOKEN"]

	tp := jira.BasicAuthTransport{
		Username: userName,
		Password: password,
	}

	client, err := jira.NewClient(tp.Client(), "https://levelup.atlassian.net/")
	if err != nil {
		panic(err)
	}

	issue, _, err := client.Issue.Get("SDK-3047", nil)

	fields, _, err := client.Field.GetList()
	epicNameKey := ""
	epicLinkKey := ""

	for _, field := range fields {
		if field.Name == "Epic Name" {
			epicNameKey = field.ID
		}
		if field.Name == "Epic Link" {
			epicLinkKey = field.ID
		}
		if field.ID == "customfield_13615" {
			println("here")
		}
	}

	i := jira.Issue{
		Fields: &jira.IssueFields{
			Components: []*jira.Component{
				{Name: "Android SDK"},
			},
			Description: "Test Issue",
			Labels:      []string{"higgs-pod"},
			Type: jira.IssueType{
				Name: "Epic",
			},
			Project: jira.Project{
				Key: "SDK",
			},
			Summary: "Android:: Tix Generated Test Epic",
			Unknowns: map[string]interface{}{
				epicNameKey: "Nick Test",
			},
		},
	}
	epicIssue, response, err := client.Issue.Create(&i)

	i = jira.Issue{
		Fields: &jira.IssueFields{
			Components: []*jira.Component{
				{Name: "Android SDK"},
			},
			Description: "Test Issue",
			Labels:      []string{"higgs-pod"},
			Type: jira.IssueType{
				Name: "Story",
			},
			Project: jira.Project{
				Key: "SDK",
			},
			Summary: "Android:: Tix Generated Test Issue",
			Unknowns: map[string]interface{}{
				epicLinkKey: epicIssue.Key,
				"customfield_13615": map[string]interface{}{
					"value": "No design review required",
				},
			},
		},
	}

	issue, response, err = client.Issue.Create(&i)

	body, err := ioutil.ReadAll(response.Body)
	bodyStr := string(body)
	println(issue)
	println(bodyStr)
	println(fields)
}

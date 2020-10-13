package env

import (
	"errors"
	"fmt"
)

const JiraApiToken = "JIRA_API_TOKEN"
const JiraUsername = "JIRA_USERNAME"

func CheckJiraEnvironment(envMap map[string]string) error {
	username := envMap[JiraUsername]
	apiToken := envMap[JiraApiToken]

	if len(username) == 0 {
		message := fmt.Sprintf(
			":scream: missing jira user name, please set the %v environment variable",
			JiraUsername)
		return errors.New(message)
	}

	if len(apiToken) == 0 {
		message := fmt.Sprintf(
			":scream: missing jira api token, please set the %v environment variable",
			JiraUsername)
		return errors.New(message)
	}
	return nil
}

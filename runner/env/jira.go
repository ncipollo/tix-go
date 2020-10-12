package env

import (
	"errors"
	"fmt"
)

const EnvJiraApiToken = "JIRA_API_TOKEN"
const EnvJiraUsername = "JIRA_USERNAME"

func CheckJiraEnvironment(envMap map[string]string) error {
	username := envMap[EnvJiraUsername]
	apiToken := envMap[EnvJiraApiToken]

	if len(username) == 0 {
		message := fmt.Sprintf(
			":scream: missing jira user name, please set the %v environment variable",
			EnvJiraUsername)
		return errors.New(message)
	}

	if len(apiToken) == 0 {
		message := fmt.Sprintf(
			":scream: missing jira api token, please set the %v environment variable",
			EnvJiraUsername)
		return errors.New(message)
	}
	return nil
}

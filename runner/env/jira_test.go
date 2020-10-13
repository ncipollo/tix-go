package env

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_checkJiraEnvironment_MissingApiToken(t *testing.T) {
	envMap := map[string]string{
		JiraUsername: "jira",
	}
	err := CheckJiraEnvironment(envMap)
	assert.Error(t, err)
}

func Test_checkJiraEnvironment_MissingUserName(t *testing.T) {
	envMap := map[string]string{
		JiraApiToken: "token",
	}
	err := CheckJiraEnvironment(envMap)
	assert.Error(t, err)
}

func Test_checkJiraEnvironment_NoMissingVariables(t *testing.T) {
	envMap := map[string]string{
		JiraUsername: "jira",
		JiraApiToken: "token",
	}
	err := CheckJiraEnvironment(envMap)
	assert.NoError(t, err)
}

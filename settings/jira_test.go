package settings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJira_Configured_IsConfigured(t *testing.T) {
	jiraSettings := Jira{Url: "url"}
	assert.True(t, jiraSettings.Configured())
}

func TestJira_Configured_IsNotConfigured(t *testing.T) {
	jiraSettings := Jira{Url: ""}
	assert.False(t, jiraSettings.Configured())
}
package settings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFromData_Empty(t *testing.T) {
	yaml := ""
	settings, err := FromData([]byte(yaml))

	assert.NoError(t, err)
	assert.Equal(t, Settings{}, settings)
}

func TestFromData_NoError(t *testing.T) {
	yaml := `
jira:
  no_epics: true
  url: https://api.example.com
  tickets:
    default: 
      default: default
    epic:
      epic: epic
    issue:
      issue: issue
    task:
      task: task
variables:
  key: value
`
	settings, err := FromData([]byte(yaml))

	expected := Settings{
		Github: Github{},
		Jira: Jira{
			NoEpics: true,
			Url:     "https://api.example.com",
			Tickets: JiraTicketFields{
				Default: map[string]interface{}{"default": "default",},
				Epic:    map[string]interface{}{"epic": "epic",},
				Issue:   map[string]interface{}{"issue": "issue",},
				Task:    map[string]interface{}{"task": "task",},
			},
		},
		Variables: map[string]string{"key": "value"},
	}
	assert.NoError(t, err)
	assert.Equal(t, expected, settings)
}

func TestFromData_ParsingError(t *testing.T) {
	yaml := `
	error
`
	_, err := FromData([]byte(yaml))

	assert.Error(t, err)
}

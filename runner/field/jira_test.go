package field

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"tix/settings"
)

func Test_jiraFieldState_noEpics_allFieldLevels(t *testing.T) {
	tixSettings := jiraWithAllFieldLevels(true)

	fieldState := JiraFieldState(tixSettings)

	assert.Equal(t, map[string]interface{}{"default": "default", "issue": "issue"}, fieldState.FieldsForLevel(0))
	assert.Equal(t, map[string]interface{}{"default": "default", "task": "task"}, fieldState.FieldsForLevel(1))
}

func Test_jiraFieldState_withEpics_allFieldLevels(t *testing.T) {
	tixSettings := jiraWithAllFieldLevels(false)

	fieldState := JiraFieldState(tixSettings)

	assert.Equal(t, map[string]interface{}{"default": "default", "epic": "epic"}, fieldState.FieldsForLevel(0))
	assert.Equal(t, map[string]interface{}{"default": "default", "issue": "issue"}, fieldState.FieldsForLevel(1))
	assert.Equal(t, map[string]interface{}{"default": "default", "task": "task"}, fieldState.FieldsForLevel(2))
}

func Test_jiraFieldState_noFieldLevels(t *testing.T) {
	tixSettings := jiraSettingsWithNoFields()

	fieldState := JiraFieldState(tixSettings)

	assert.Equal(t, map[string]interface{}{}, fieldState.FieldsForLevel(0))
	assert.Equal(t, map[string]interface{}{}, fieldState.FieldsForLevel(1))
	assert.Equal(t, map[string]interface{}{}, fieldState.FieldsForLevel(2))
}

func jiraSettingsWithNoFields() settings.Settings {
	return settings.Settings{}
}

func jiraWithAllFieldLevels(noEpics bool) settings.Settings {
	return settings.Settings{
		Jira: settings.Jira{
			NoEpics: noEpics,
			Tickets: settings.JiraTicketFields{
				Default: map[string]interface{}{
					"default": "default",
				},
				Epic: map[string]interface{}{
					"epic": "epic",
				},
				Issue: map[string]interface{}{
					"issue": "issue",
				},
				Task: map[string]interface{}{
					"task": "task",
				},
			},
		},
	}
}

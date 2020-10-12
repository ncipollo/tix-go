package field

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"tix/settings"
)

func Test_jiraFieldState_noEpics_allFieldLevels(t *testing.T) {
	tixSettings := settingsWithAllFieldLevels(true)

	fieldState := JiraFieldState(tixSettings)

	assert.Equal(t, map[string]interface{}{"default": "default", "issue": "issue",}, fieldState.FieldsForLevel(0))
	assert.Equal(t, map[string]interface{}{"default": "default", "task": "task",}, fieldState.FieldsForLevel(1))
}

func Test_jiraFieldState_withEpics_allFieldLevels(t *testing.T) {
	tixSettings := settingsWithAllFieldLevels(false)

	fieldState := JiraFieldState(tixSettings)

	assert.Equal(t, map[string]interface{}{"default": "default", "epic": "epic",}, fieldState.FieldsForLevel(0))
	assert.Equal(t, map[string]interface{}{"default": "default", "issue": "issue",}, fieldState.FieldsForLevel(1))
	assert.Equal(t, map[string]interface{}{"default": "default", "task": "task",}, fieldState.FieldsForLevel(2))
}

func Test_jiraFieldState_noFieldLevels(t *testing.T) {
	tixSettings := settingsWithNoFields()

	fieldState := JiraFieldState(tixSettings)

	assert.Equal(t, map[string]interface{}{}, fieldState.FieldsForLevel(0))
	assert.Equal(t, map[string]interface{}{}, fieldState.FieldsForLevel(1))
	assert.Equal(t, map[string]interface{}{}, fieldState.FieldsForLevel(2))
}

func settingsWithNoFields() settings.Settings {
	return settings.Settings{}
}

func settingsWithAllFieldLevels(noEpics bool) settings.Settings {
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

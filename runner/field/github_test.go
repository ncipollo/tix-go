package field

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"tix/settings"
)

func TestGithubFieldState_NoProjects(t *testing.T) {
	tixSettings := githubWithAllFieldLevels(true)
	fieldState := GithubFieldState(tixSettings)
	assert.Equal(t, map[string]interface{}{"default": "default", "issue": "issue"}, fieldState.FieldsForLevel(0))
}

func TestGithubFieldState_WithProjects(t *testing.T) {
	tixSettings := githubWithAllFieldLevels(false)

	fieldState := GithubFieldState(tixSettings)

	assert.Equal(
		t,
		map[string]interface{}{"default": "default", "project": "project"},
		fieldState.FieldsForLevel(0),
	)
	assert.Equal(t, map[string]interface{}{"default": "default", "issue": "issue"}, fieldState.FieldsForLevel(1))
}

func TestGithubFieldState_noFieldLevels(t *testing.T) {
	tixSettings := githubSettingsWithNoFields()

	fieldState := GithubFieldState(tixSettings)

	assert.Equal(t, map[string]interface{}{}, fieldState.FieldsForLevel(0))
	assert.Equal(t, map[string]interface{}{}, fieldState.FieldsForLevel(1))
}

func githubSettingsWithNoFields() settings.Settings {
	return settings.Settings{}
}

func githubWithAllFieldLevels(noProjects bool) settings.Settings {
	return settings.Settings{
		Github: settings.Github{
			NoProjects: noProjects,
			Tickets: settings.GithubTicketFields{
				Default: map[string]interface{}{
					"default": "default",
				},
				Project: map[string]interface{}{
					"project": "project",
				},
				Issue: map[string]interface{}{
					"issue": "issue",
				},
			},
		},
	}
}

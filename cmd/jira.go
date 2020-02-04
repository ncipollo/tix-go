package cmd

import (
	"errors"
	"fmt"
	"tix/creator/jira"
	"tix/md"
	"tix/settings"
)

const EnvJiraApiToken = "JIRA_API_TOKEN"
const EnvJiraUsername = "JIRA_USERNAME"

func createJiraApi(envMap map[string]string, settings settings.Settings) jira.Api {
	return jira.NewApi(envMap[EnvJiraUsername], envMap[EnvJiraApiToken], settings.Jira.Url)
}

func jiraCreator(api jira.Api, settings settings.Settings) *jira.Creator {
	if settings.Jira.NoEpics {
		return jira.NewCreatorWithoutEpics(api)
	} else {
		return jira.NewCreatorWithEpics(api)
	}
}

func checkJiraEnvironment(envMap map[string]string) error {
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

func jiraFieldState(settings settings.Settings) *md.FieldState {
	if settings.Jira.NoEpics {
		return jiraFieldStateWithoutEpics(settings)
	} else {
		return jiraFieldStateWithEpics(settings)
	}
}

func jiraFieldStateWithEpics(settings settings.Settings) *md.FieldState {
	fieldState := md.NewFieldState()
	ticketSettings := settings.Jira.Tickets
	if ticketSettings.Default != nil {
		fieldState.SetDefaultFields(ticketSettings.Default)
	}
	if ticketSettings.Epic != nil {
		fieldState.SetFieldsForLevel(ticketSettings.Epic, 0)
	}
	if ticketSettings.Issue != nil {
		fieldState.SetFieldsForLevel(ticketSettings.Issue, 1)
	}
	if ticketSettings.Task != nil {
		fieldState.SetFieldsForLevel(ticketSettings.Task, 2)
	}
	return fieldState
}

func jiraFieldStateWithoutEpics(settings settings.Settings) *md.FieldState {
	fieldState := md.NewFieldState()
	ticketSettings := settings.Jira.Tickets
	if ticketSettings.Default != nil {
		fieldState.SetDefaultFields(ticketSettings.Default)
	}
	if ticketSettings.Issue != nil {
		fieldState.SetFieldsForLevel(ticketSettings.Issue, 0)
	}
	if ticketSettings.Task != nil {
		fieldState.SetFieldsForLevel(ticketSettings.Task, 1)
	}
	return fieldState
}

package field

import (
	"tix/md"
	"tix/settings"
)

func JiraFieldState(settings settings.Settings) *md.FieldState {
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
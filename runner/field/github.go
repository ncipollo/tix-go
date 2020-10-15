package field

import (
	"tix/md"
	"tix/settings"
)

func GithubFieldState(settings settings.Settings) *md.FieldState {
	if settings.Github.NoProjects {
		return githubFieldStateWithoutProjects(settings)
	} else {
		return githubFieldStateWithProjects(settings)
	}
}

func githubFieldStateWithoutProjects(settings settings.Settings) *md.FieldState {
	fieldState := md.NewFieldState()
	ticketSettings := settings.Github.Tickets
	if ticketSettings.Default != nil {
		fieldState.SetDefaultFields(ticketSettings.Default)
	}
	if ticketSettings.Issue != nil {
		fieldState.SetFieldsForLevel(ticketSettings.Issue, 0)
	}
	return fieldState
}

func githubFieldStateWithProjects(settings settings.Settings) *md.FieldState {
	fieldState := md.NewFieldState()
	ticketSettings := settings.Github.Tickets
	if ticketSettings.Default != nil {
		fieldState.SetDefaultFields(ticketSettings.Default)
	}
	if ticketSettings.Project != nil {
		fieldState.SetFieldsForLevel(ticketSettings.Project, 0)
	}
	if ticketSettings.Issue != nil {
		fieldState.SetFieldsForLevel(ticketSettings.Issue, 1)
	}
	return fieldState
}
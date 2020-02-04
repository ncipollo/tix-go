package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"tix/creator/jira"
	"tix/md"
	"tix/settings"
)

const EnvJiraApiToken = "JIRA_API_TOKEN"
const EnvJiraUsername = "JIRA_USERNAME"

type TixCommand struct {
	envMap       map[string]string
	markdownPath string
	settingsPath string
}

func NewTixCommand(envMap map[string]string, markdownPath string) *TixCommand {
	directory := filepath.Dir(markdownPath)
	settingsPath := filepath.Join(directory, "tix.yml")

	return &TixCommand{
		envMap:       envMap,
		markdownPath: markdownPath,
		settingsPath: settingsPath,
	}
}

func (t TixCommand) Run() error {
	tixSettings, err := t.loadSettings()
	if err != nil {
		return err
	}
	markdownData, err := t.loadMarkDownData()

	return t.generateJiraTickets(markdownData, tixSettings)
}

func (t TixCommand) loadSettings() (settings.Settings, error) {
	data, err := ioutil.ReadFile(t.settingsPath)
	if err != nil {
		message := fmt.Sprintf(":scream: failed to open settings\n%v", err)
		return settings.Settings{}, errors.New(message)
	}

	return settings.FromData(data)
}

func (t TixCommand) loadMarkDownData() ([]byte, error) {
	return ioutil.ReadFile(t.markdownPath)
}

func (t TixCommand) generateJiraTickets(markdownData []byte, settings settings.Settings) error {
	if len(settings.Jira.Url) == 0 {
		return nil
	}

	err := t.checkJiraEnvironment()
	if err != nil {
		return err
	}

	fieldState := t.jiraFieldState(settings)
	markdownParser := md.NewParser(fieldState)
	tickets, err := markdownParser.Parse(markdownData)

	if err != nil {
		return err
	}

	creator := t.jiraCreator(settings)
	creator.CreateTickets(tickets)

	return nil
}

func (t TixCommand) checkJiraEnvironment() error {
	username := t.envMap[EnvJiraUsername]
	apiToken := t.envMap[EnvJiraApiToken]

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

func (t TixCommand) jiraFieldState(settings settings.Settings) *md.FieldState {
	if settings.Jira.NoEpics {
		return t.jiraFieldStateWithoutEpics(settings)
	} else {
		return t.jiraFieldStateWithEpics(settings)
	}
}

func (t TixCommand) jiraFieldStateWithEpics(settings settings.Settings) *md.FieldState {
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

func (t TixCommand) jiraFieldStateWithoutEpics(settings settings.Settings) *md.FieldState {
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

func (t TixCommand) jiraCreator(settings settings.Settings) *jira.Creator {
	api := jira.NewApi(t.envMap[EnvJiraUsername], t.envMap[EnvJiraApiToken], settings.Jira.Url)
	var startingLevel int
	if settings.Jira.NoEpics {
		startingLevel = 1
	} else {
		startingLevel = 0
	}
	return jira.NewCreator(api, startingLevel)
}

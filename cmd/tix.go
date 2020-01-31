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
	return ioutil.ReadFile(t.settingsPath)
}

func (t TixCommand) generateJiraTickets(markdownData []byte, settings settings.Settings) error {
	if len(settings.Jira.Url) == 0 {
		return nil
	}

	err := t.checkJiraEnvironment()
	if err != nil {
		return err
	}

	fieldState := md.NewFieldState()
	markdownParser := md.NewParser(fieldState)
	tickets, err := markdownParser.Parse(markdownData)

	if err != nil {
		return err
	}

	api := jira.NewApi(t.envMap[EnvJiraUsername], t.envMap[EnvJiraApiToken], settings.Jira.Url)
	creator := jira.NewCreator(api)
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

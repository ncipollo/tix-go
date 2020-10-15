package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"tix/runner"
	"tix/settings"
	"tix/transform"
)

type TixCommand struct {
	dryRun       bool
	envMap       map[string]string
	markdownPath string
	settingsPath string
}

func NewTixCommand(dryRun bool, envMap map[string]string, markdownPath string) *TixCommand {
	directory := filepath.Dir(markdownPath)
	settingsPath := filepath.Join(directory, "tix.yml")

	return &TixCommand{
		dryRun:       dryRun,
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
	if err != nil {
		return err
	}

	markdownData = t.transformMarkDownData(markdownData, tixSettings)

	runners := t.createRunners(tixSettings)
	if t.dryRun {
		for _, tixRunner := range runners {
			err := tixRunner.DryRun(markdownData)
			if err != nil {
				return err
			}
		}
	} else {
		for _, tixRunner := range runners {
			err := tixRunner.Run(markdownData)
			if err != nil {
				return err
			}
		}
	}
	return nil
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

func (t TixCommand) transformMarkDownData(markdownData []byte, settings settings.Settings) []byte {
	return transform.ApplyVariableTransform(markdownData, t.envMap, settings.Variables)
}

func (t TixCommand) createRunners(tixSettings settings.Settings) []runner.TixRunner {
	return []runner.TixRunner{
		runner.NewGithubRunner(t.envMap, &tixSettings),
		runner.NewJiraRunner(t.envMap, &tixSettings),
	}
}

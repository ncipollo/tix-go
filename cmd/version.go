package cmd

import "tix/logger"

type VersionCommand struct {
	version string
}

func NewVersionCommand(version string) *VersionCommand {
	return &VersionCommand{version}
}

func (v VersionCommand) Run() error {
	logger.Message(":rocket: tix version %v", v.version)
	return nil
}

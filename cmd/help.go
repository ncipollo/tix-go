package cmd

import "tix/logger"

const helpMessage =
`Tix - Ticket Generator
`

type HelpCommand struct {

}

func NewHelpCommand() *HelpCommand {
	return &HelpCommand{}
}

func (h HelpCommand) Run() error {
	logger.Output(helpMessage)
	return nil
}


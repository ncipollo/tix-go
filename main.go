package main

import (
	"os"
	"tix/cmd"
	"tix/logger"
)

var version = "0.0"

func main() {
	parser := cmd.NewParser(os.Environ(), version)
	command := parser.Parse()
	if command != nil {
		err := command.Run()
		if err == nil {
			os.Exit(0)
		} else {
			logger.Error("%v", err)
			os.Exit(1)
		}
	}
}
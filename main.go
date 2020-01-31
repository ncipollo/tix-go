package main

import (
	"os"
	"tix/cmd"
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
			os.Exit(1)
		}
	}
}
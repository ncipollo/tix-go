package cmd

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"tix/env"
	"tix/logger"
)

const helpUsage = "prints help for tix"
const quietUsage = "suppresses all log output except errors"
const verboseUsage = "enables verbose output"
const versionUsage = "prints tix version"
const shortHandSuffix = " (shorthand)"

type Parser struct {
	env     []string
	flag    *flag.FlagSet
	version string
}

func NewParser(env []string, version string) *Parser {
	name := os.Args[0]
	flagSet := flag.NewFlagSet(name, flag.ExitOnError)
	flagSet.Usage = func() {
		_, _ = fmt.Fprintf(flagSet.Output(),
			"Usage: %s [OPTIONS] <markdown file> \n",
			name)
		flagSet.PrintDefaults()
	}
	return &Parser{env: env, flag: flagSet, version: version}
}

func (parser *Parser) Parse() Command {
	quiet := parser.setupQuiet()
	help := parser.setupHelp()
	verbose := parser.setupVerbose()
	version := parser.setupVersion()

	_ = parser.flag.Parse(os.Args[1:])

	parser.adjustLogLevel(*quiet, *verbose)

	if *version {
		return NewVersionCommand(parser.version)
	} else if *help {
		return NewHelpCommand()
	} else {
		if parser.flag.NArg() < 1 {
			parser.printUsageAndExit()
		}
		path, _ := parser.localPath()
		return NewTixCommand(env.Map(), path)
	}
}

func (parser *Parser) setupHelp() *bool {
	var verbose bool
	parser.flag.BoolVar(&verbose, "help", false, helpUsage)
	parser.flag.BoolVar(&verbose, "h", false, helpUsage+shortHandSuffix)

	return &verbose
}

func (parser *Parser) setupQuiet() *bool {
	var verbose bool
	parser.flag.BoolVar(&verbose, "quiet", false, quietUsage)
	parser.flag.BoolVar(&verbose, "q", false, quietUsage+shortHandSuffix)

	return &verbose
}

func (parser *Parser) setupVerbose() *bool {
	var verbose bool
	parser.flag.BoolVar(&verbose, "verbose", false, verboseUsage)
	parser.flag.BoolVar(&verbose, "v", false, verboseUsage+shortHandSuffix)

	return &verbose
}

func (parser *Parser) setupVersion() *bool {
	var verbose bool
	parser.flag.BoolVar(&verbose, "version", false, versionUsage)

	return &verbose
}

func (parser *Parser) printUsageAndExit() {
	parser.flag.Usage()
	os.Exit(2)
}

func (parser Parser) adjustLogLevel(quiet bool, verbose bool) {
	if quiet {
		logger.SetLogLevel(logger.LogLevelQuiet)
	} else if verbose {
		logger.SetLogLevel(logger.LogLevelVerbose)
	} else {
		logger.SetLogLevel(logger.LogLevelNormal)
	}
}

func (parser *Parser) localPath() (string, error) {
	relativePath := parser.flag.Arg(0)
	return filepath.Abs(relativePath)
}

package logger

import (
	"fmt"
	"github.com/kyokomi/emoji"
)

type LogLevel int

const (
	LogLevelQuiet LogLevel = iota
	LogLevelNormal
	LogLevelVerbose
)

var currentLogLevel LogLevel = LogLevelNormal

func SetLogLevel(level LogLevel) {
	currentLogLevel = level
}

func Error(message string, args ...interface{})  {
	printLog(message, args...)
}

func Message(message string, args ...interface{})  {
	if currentLogLevel >= LogLevelNormal {
		printLog(message, args...)
	}
}

func Verbose(message string, args ...interface{})  {
	if currentLogLevel == LogLevelVerbose {
		printLog(message, args...)
	}
}

func printLog(message string, args ...interface{})  {
	formatted := fmt.Sprintf(message, args...)
	_, _ = emoji.Println(formatted)
}

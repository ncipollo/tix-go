package reporter

import (
	"fmt"
	"strings"
	"tix/logger"
)

func ReportFailedTicket(err error, startingLevel int, level int) {
	var builder strings.Builder
	for ii := startingLevel; ii < level; ii++ {
		builder.WriteString("\t")
	}
	builder.WriteString("- ")
	builder.WriteString(err.Error())

	logger.Error(builder.String())
}

func ReportSuccessfulTicket(issueKey string, startingLevel int, level int, title string) {
	var builder strings.Builder
	for ii := startingLevel; ii < level; ii++ {
		builder.WriteString("\t")
	}
	message := fmt.Sprintf("- :tada: %v: %v created", issueKey, title)
	builder.WriteString(message)

	logger.Message(builder.String())
}

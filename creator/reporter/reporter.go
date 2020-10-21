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

func ReportSuccessfulTicket(issueKey string, startingLevel int, level int, title string, updateKey string) {
	var builder strings.Builder
	for ii := startingLevel; ii < level; ii++ {
		builder.WriteString("\t")
	}
	
	var verb string
	if len(updateKey) > 0 {
		verb = "updated"
	} else {
		verb = "created"
	}

	message := fmt.Sprintf("- :tada: %v: %v %s", issueKey, title, verb)
	builder.WriteString(message)

	logger.Message(builder.String())
}

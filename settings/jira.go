package settings

type Jira struct {
	NoEpics bool  `yaml:"no_epics"`
	Url string
	Tickets JiraTicketFields
}

func (j Jira) Configured() bool {
	return len(j.Url) > 0
}

type JiraTicketFields struct {
	Default map[string]interface{}
	Epic map[string]interface{}
	Issue map[string]interface{}
	Task map[string]interface{}
}

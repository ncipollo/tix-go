package settings

type Jira struct {
	NoEpics bool  `yaml:"no_epics"`
	Url string
	Tickets JiraTicketFields
}

type JiraTicketFields struct {
	Default map[string]interface{}
	Epic map[string]interface{}
	Issue map[string]interface{}
	Task map[string]interface{}
}

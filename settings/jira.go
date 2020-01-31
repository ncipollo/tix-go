package settings

type Jira struct {
	Url string
	Tickets JiraTicketFields
}

type JiraTicketFields struct {
	Default map[string]interface{}
	Epic map[string]interface{}
	Issue map[string]interface{}
	Task map[string]interface{}
}

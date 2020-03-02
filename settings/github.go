package settings

type Github struct {
	NoProjects bool  `yaml:"no_projects"`
	Owner string
	Repo string
	Tickets GithubTicketFields
}

type GithubTicketFields struct {
	Default map[string]interface{}
	Project map[string]interface{}
	Issue map[string]interface{}
}

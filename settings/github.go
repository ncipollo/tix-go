package settings

type Github struct {
	NoProjects bool  `yaml:"no_projects"`
	Owner string
	Repo string
	Tickets GithubTicketFields
}

func (g Github) Configured() bool {
	return len(g.Owner) > 0 && len(g.Repo) > 0
}

type GithubTicketFields struct {
	Default map[string]interface{}
	Project map[string]interface{}
	Issue map[string]interface{}
}

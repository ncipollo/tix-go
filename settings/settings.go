package settings

import "gopkg.in/yaml.v2"

type Settings struct {
	Github    Github
	Jira      Jira
	Variables map[string]string
}

func FromData(data []byte) (Settings, error) {
	settings := Settings{}
	err := yaml.Unmarshal(data, &settings)
	return settings, err
}

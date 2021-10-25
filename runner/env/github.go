package env

import (
	"errors"
	"fmt"
)

const GithubApiToken = "GITHUB_API_TOKEN"

func CheckGithubEnvironment(envMap map[string]string) error {
	apiToken := envMap[GithubApiToken]
	if len(apiToken) == 0 {
		message := fmt.Sprintf(
			":scream: missing github api token, please set the %v environment variable",
			GithubApiToken)
		return errors.New(message)
	}
	return nil
}

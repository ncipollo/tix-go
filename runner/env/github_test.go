package env

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckGithubEnvironment_MissingApiToken(t *testing.T) {
	envMap := map[string]string{}
	err := CheckGithubEnvironment(envMap)
	assert.Error(t, err)
}

func TestCheckGithubEnvironment_NoErrors(t *testing.T) {
	envMap := map[string]string{
		GithubApiToken: "token",
	}
	err := CheckGithubEnvironment(envMap)
	assert.NoError(t, err)
}
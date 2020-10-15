package settings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGithub_Configured_IsConfigured(t *testing.T) {
	githubSettings := Github{Owner: "owner", Repo: "repo"}
	assert.True(t, githubSettings.Configured())
}

func TestGithub_Configured_IsNotConfigured_NoOwner(t *testing.T) {
	githubSettings := Github{Owner: "", Repo: "repo"}
	assert.False(t, githubSettings.Configured())
}

func TestGithub_Configured_IsNotConfigured_NoRepo(t *testing.T) {
	githubSettings := Github{Owner: "owner", Repo: ""}
	assert.False(t, githubSettings.Configured())
}

package transform

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApplyVariableTransform_EmptyStringIfEnvironmentVariableIsMissing(t *testing.T) {
	inputStr := "hello$replace"
	variables := map[string]string{
		"$replace": "$env1",
	}
	env := make(map[string]string)

	transformed := ApplyVariableTransform([]byte(inputStr), env, variables)

	assert.Equal(t, "hello", string(transformed))
}

func TestApplyVariableTransform_WithEnvironmentVariables(t *testing.T) {
	inputStr := "$replace1 world$replace2"
	variables := map[string]string{
		"$replace1": "hello",
		"$replace2": "$env",
	}
	env := map[string]string{
		"env": "!",
	}

	transformed := ApplyVariableTransform([]byte(inputStr), env, variables)

	assert.Equal(t, "hello world!", string(transformed))
}

func TestApplyVariableTransform_WithStaticVariables(t *testing.T) {
	inputStr := "$replace1 world$replace2"
	variables := map[string]string{
		"$replace1": "hello",
		"$replace2": ".",
	}
	env := make(map[string]string)

	transformed := ApplyVariableTransform([]byte(inputStr), env, variables)

	assert.Equal(t, "hello world.", string(transformed))
}

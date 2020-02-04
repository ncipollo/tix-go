package transform

import "strings"

func ApplyVariableTransform(inputData []byte, envMap map[string]string, variables map[string]string) []byte {
	transformed := string(inputData)

	for key, value := range variables {
		variable := variableFromValue(value, envMap)
		transformed = strings.ReplaceAll(transformed, key, variable)
	}

	return []byte(transformed)
}

func variableFromValue(value string, envMap map[string]string) string {
	if strings.HasPrefix(value, "$") {
		envKey := strings.TrimPrefix(value, "$")
		return envMap[envKey]
	}
	return value
}

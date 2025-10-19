package utils

import "fmt"

func ParsePath(path string, variableNames ...string) string {
	var variables []any
	for _, v := range variableNames {
		variables = append(variables, fmt.Sprintf("{%s:[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}}", v))
	}

	return fmt.Sprintf(path, variables...)
}

func ParsePathWithAll(path string, variableNames ...string) string {
	var variables []any
	for _, v := range variableNames {
		variables = append(variables, fmt.Sprintf("{%s:(?:all|[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12})}", v))
	}

	return fmt.Sprintf(path, variables...)
}

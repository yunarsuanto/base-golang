package user_repository

import (
	"fmt"
	"strings"
)

func mapQueryFilterListUser(search string, params *[]any) (string, error) {
	var result string
	var filterArray []string

	if search != "" {
		searchParams := fmt.Sprintf("%%%s%%", search)
		filterArray = append(filterArray, "u.username LIKE $1")
		(*params) = append((*params), searchParams)
	}

	result = strings.Join(filterArray, " AND ")
	if result != "" {
		result = fmt.Sprintf("WHERE %s AND %s", result, "u.username != 'superadmin'")
	} else {
		result = fmt.Sprintf("WHERE %s", "u.username != 'superadmin'")
	}

	return result, nil
}

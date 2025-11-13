package category_lesson_repository

import (
	"fmt"
	"strings"
)

func mapQueryFilterListCategoryLesson(search string, params *[]any) (string, error) {
	var result string
	var filterArray []string

	if search != "" {
		searchParams := fmt.Sprintf("%%%s%%", search)
		filterArray = append(filterArray, "u.title LIKE $1")
		(*params) = append((*params), searchParams)
	}

	result = strings.Join(filterArray, " AND ")
	if result != "" {
		result = fmt.Sprintf("WHERE %s", result)
	}
	return result, nil
}

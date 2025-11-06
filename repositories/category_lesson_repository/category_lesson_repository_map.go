package category_lesson_repository

import (
	"fmt"
	"strings"
)

func mapQueryFilterListCategoryLesson(search string, hasParent bool, params *[]any) (string, error) {
	var result string
	var filterArray []string

	if search != "" {
		searchParams := fmt.Sprintf("%%%s%%", search)
		filterArray = append(filterArray, "u.title LIKE $1")
		(*params) = append((*params), searchParams)
	}

	if hasParent {
		filterArray = append(filterArray, "u.category_lesson_id IS NOT NULL")
	} else {
		filterArray = append(filterArray, "u.category_lesson_id IS NULL")
	}

	result = strings.Join(filterArray, " AND ")
	if result != "" {
		result = fmt.Sprintf("WHERE %s", result)
	}
	return result, nil
}

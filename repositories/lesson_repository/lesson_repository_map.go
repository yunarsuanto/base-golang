package lesson_repository

import (
	"fmt"
	"strings"
)

func mapQueryFilterListLesson(search string, params *[]any, categoryLessonId string) (string, error) {
	var result string
	var filterArray []string

	if search != "" {
		searchParams := fmt.Sprintf("%%%s%%", search)
		filterArray = append(filterArray, "u.title LIKE $2")
		(*params) = append((*params), searchParams)
	}

	if categoryLessonId != "" {
		filterArray = append(filterArray, "u.category_lesson_id = $1")
	}

	result = strings.Join(filterArray, " AND ")
	if result != "" {
		result = fmt.Sprintf("WHERE %s", result)
	}
	return result, nil
}

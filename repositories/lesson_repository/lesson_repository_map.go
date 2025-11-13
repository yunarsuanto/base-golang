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
		filterArray = append(filterArray, fmt.Sprintf("u.title ILIKE $%d", len(*params)+1))
		(*params) = append((*params), searchParams)
	}

	if categoryLessonId != "" {
		filterArray = append(filterArray, fmt.Sprintf(`u.category_lesson_id = $%d`, len(*params)+1))
		(*params) = append((*params), categoryLessonId)
	}

	result = strings.Join(filterArray, " AND ")
	if result != "" {
		result = fmt.Sprintf("WHERE %s", result)
	}
	return result, nil
}

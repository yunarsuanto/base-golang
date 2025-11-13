package lesson_item_repository

import (
	"fmt"
	"strings"
)

func mapQueryFilterListLessonItem(search string, params *[]any, lessonId string) (string, error) {
	var result string
	var filterArray []string

	if search != "" {
		searchParams := fmt.Sprintf("%%%s%%", search)
		filterArray = append(filterArray, fmt.Sprintf(`u.content ILIKE $%d`, len(*params)+1))
		(*params) = append((*params), searchParams)
	}

	if lessonId != "" {
		filterArray = append(filterArray, fmt.Sprintf(`u.lesson_id = $%d`, len(*params)+1))
		(*params) = append((*params), lessonId)
	}

	result = strings.Join(filterArray, " AND ")
	if result != "" {
		result = fmt.Sprintf("WHERE %s", result)
	}
	return result, nil
}

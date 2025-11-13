package models

const LessonDatatitle = "lesson"

type ListLesson struct {
	Id                  string `db:"id"`
	Title               string `db:"title"`
	Description         string `db:"description"`
	CategoryLessonId    string `db:"category_lesson_id"`
	LessonType          string `db:"lesson_type"`
	Media               string `db:"media"`
	CategoryLessonTitle string `db:"category_lesson_title"`
	Level               uint32 `db:"level"`
}

func (ListLesson) ColumnQuery() string {
	return `
		u.id,
		u.title,
		u.description,
		u.category_lesson_id,
		u.lesson_type,
		u.media,
		cl.title as category_lesson_title,
		u.level
	`
}

func (ListLesson) TableQuery() string {
	return `
		FROM lessons u
		JOIN category_lessons cl ON cl.id = u.category_lesson_id
	`
}

type DetailLesson struct {
	Id               string `db:"id"`
	Title            string `db:"title"`
	Description      string `db:"description"`
	CategoryLessonId string `db:"category_lesson_id"`
	LessonType       string `db:"lesson_type"`
	Media            string `db:"media"`
	Level            uint32 `db:"level"`
}

func (DetailLesson) ColumnQuery() string {
	return `
		u.id,
		u.title,
		u.description,
		u.category_lesson_id,
		u.lesson_type,
		u.media,
		u.level
	`
}

func (DetailLesson) TableQuery() string {
	return `
		FROM lessons u
	`
}

func (DetailLesson) FilterQuery() string {
	return `
		WHERE u.id = $1
	`
}

type CreateLesson struct {
	Title            string `db:"title"`
	Description      string `db:"description"`
	CategoryLessonId string `db:"category_lesson_id"`
	LessonType       string `db:"lesson_type"`
	Media            string `db:"media"`
	Level            uint32 `db:"level"`
}

func (CreateLesson) InsertQuery() string {
	return `
		INSERT INTO
		lessons (
			title,
			description,
			category_lesson_id,
			lesson_type,
			media,
			level
		) VALUES (
			:title,
			:description,
			:category_lesson_id,
			:lesson_type,
			:media,
			:level
		)
		RETURNING id;
	`
}

type UpdateLesson struct {
	Id               string `db:"id"`
	Title            string `db:"title"`
	Description      string `db:"description"`
	CategoryLessonId string `db:"category_lesson_id"`
	LessonType       string `db:"lesson_type"`
	Media            string `db:"media"`
	Level            uint32 `db:"level"`
}

func (UpdateLesson) InsertQuery() string {
	return `
		UPDATE lessons SET
			title = :title,
			description = :description,
			category_lesson_id = :category_lesson_id,
			lesson_type = :lesson_type,
			media = :media,
			level = :level
		WHERE id = :id
	`
}

type DeleteLesson struct {
	Id string `db:"id"`
}

func (DeleteLesson) InsertQuery() string {
	return `
		DELETE FROM lessons WHERE id = :id
	`
}

package models

const CategoryLessonDataName = "category_lesson"

type ListCategoryLesson struct {
	Id               string  `db:"id"`
	Title            string  `db:"title"`
	Description      string  `db:"description"`
	CategoryLessonId *string `db:"category_lesson_id"`
	Media            string  `db:"media"`
}

func (ListCategoryLesson) ColumnQuery() string {
	return `
		u.id,
		u.title,
		u.description,
		u.category_lesson_id,
		u.media
	`
}

func (ListCategoryLesson) TableQuery() string {
	return `
		FROM category_lessons u
	`
}

type DetailCategoryLesson struct {
	Id               string  `db:"id"`
	Title            string  `db:"title"`
	Description      string  `db:"description"`
	CategoryLessonId *string `db:"category_lesson_id"`
	Media            string  `db:"media"`
}

func (DetailCategoryLesson) ColumnQuery() string {
	return `
		u.id,
		u.title,
		u.description,
		u.category_lesson_id,
		u.media
	`
}

func (DetailCategoryLesson) TableQuery() string {
	return `
		FROM category_lessons u
	`
}

func (DetailCategoryLesson) FilterQuery() string {
	return `
		WHERE u.id = $1 OR u.category_lesson_id = $1 ORDER BY (u.id = $1) DESC;
	`
}

type CreateCategoryLesson struct {
	Title            string  `db:"title"`
	Description      string  `db:"description"`
	CategoryLessonId *string `db:"category_lesson_id"`
	Media            string  `db:"media"`
}

func (CreateCategoryLesson) InsertQuery() string {
	return `
		INSERT INTO
		category_lessons (
			title,
			description,
			category_lesson_id,
			media
		) VALUES (
			:title,
			:description,
			:category_lesson_id,
			:media
		)
	`
}

type UpdateCategoryLesson struct {
	Id               string  `db:"id"`
	Title            string  `db:"title"`
	Description      string  `db:"description"`
	CategoryLessonId *string `db:"category_lesson_id"`
	Media            string  `db:"media"`
}

func (UpdateCategoryLesson) InsertQuery() string {
	return `
		UPDATE category_lessons SET
			title = :title,
			description = :description,
			category_lesson_id = :category_lesson_id,
			media = :media
		WHERE id = :id
	`
}

type DeleteCategoryLesson struct {
	Id string `db:"id"`
}

func (DeleteCategoryLesson) InsertQuery() string {
	return `
		DELETE FROM category_lessons WHERE id = :id
	`
}

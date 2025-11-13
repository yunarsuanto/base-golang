package models

const CategoryLessonDataName = "category_lesson"

type ListCategoryLesson struct {
	Id                 string `db:"id"`
	Title              string `db:"title"`
	Description        string `db:"description"`
	CategoryLessonType string `db:"category_lesson_type"`
	Media              string `db:"media"`
}

func (ListCategoryLesson) ColumnQuery() string {
	return `
		u.id,
		u.title,
		u.description,
		u.category_lesson_type,
		u.media
	`
}

func (ListCategoryLesson) TableQuery() string {
	return `
		FROM category_lessons u
	`
}

type DetailCategoryLesson struct {
	Id                 string `db:"id"`
	Title              string `db:"title"`
	Description        string `db:"description"`
	CategoryLessonType string `db:"category_lesson_type"`
	Media              string `db:"media"`
}

func (DetailCategoryLesson) ColumnQuery() string {
	return `
		u.id,
		u.title,
		u.description,
		u.category_lesson_type,
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
		WHERE u.id = $1;
	`
}

type CreateCategoryLesson struct {
	Title              string `db:"title"`
	Description        string `db:"description"`
	CategoryLessonType string `db:"category_lesson_type"`
	Media              string `db:"media"`
}

func (CreateCategoryLesson) InsertQuery() string {
	return `
		INSERT INTO
		category_lessons (
			title,
			description,
			category_lesson_type,
			media
		) VALUES (
			:title,
			:description,
			:category_lesson_type,
			:media
		)
	`
}

type UpdateCategoryLesson struct {
	Id                 string `db:"id"`
	Title              string `db:"title"`
	Description        string `db:"description"`
	CategoryLessonType string `db:"category_lesson_type"`
	Media              string `db:"media"`
}

func (UpdateCategoryLesson) InsertQuery() string {
	return `
		UPDATE category_lessons SET
			title = :title,
			description = :description,
			category_lesson_type = :category_lesson_type,
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

package models

const LessonItemDatacontent = "lesson_item"

type ListLessonItem struct {
	Id       string `db:"id"`
	LessonId string `db:"lesson_id"`
	Content  string `db:"content"`
	Order    uint32 `db:"order"`
}

func (ListLessonItem) ColumnQuery() string {
	return `
		u.id,
		u.lesson_id,
		u.content,
		u.order
	`
}

func (ListLessonItem) TableQuery() string {
	return `
		FROM lesson_items u
	`
}

type DetailLessonItem struct {
	Id       string `db:"id"`
	LessonId string `db:"lesson_id"`
	Content  string `db:"content"`
	Order    uint32 `db:"order"`
}

func (DetailLessonItem) ColumnQuery() string {
	return `
		u.id,
		u.lesson_id,
		u.content,
		u.order
	`
}

func (DetailLessonItem) TableQuery() string {
	return `
		FROM lesson_items u
	`
}

func (DetailLessonItem) FilterQuery() string {
	return `
		WHERE u.id = $1
	`
}

type CreateLessonItem struct {
	LessonId string `db:"lesson_id"`
	Content  string `db:"content"`
	Order    uint32 `db:"order"`
}

func (CreateLessonItem) InsertQuery() string {
	return `
		INSERT INTO
		lesson_items (
			lesson_id,
			content,
			"order"
		) VALUES (
		 	:lesson_id,
			:content,
			:order
		)
	`
}

type UpdateLessonItem struct {
	Id       string `db:"id"`
	LessonId string `db:"lesson_id"`
	Content  string `db:"content"`
	Order    uint32 `db:"order"`
}

func (UpdateLessonItem) InsertQuery() string {
	return `
		UPDATE lesson_items SET
			lesson_id = :lesson_id,
			content = :content,
			"order" = :order
		WHERE id = :id
	`
}

type DeleteLessonItem struct {
	Id string `db:"id"`
}

func (DeleteLessonItem) InsertQuery() string {
	return `
		DELETE FROM lesson_items WHERE id = :id
	`
}

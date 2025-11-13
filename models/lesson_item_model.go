package models

const LessonItemDatacontent = "lesson_item"

type ListLessonItem struct {
	Id       string `db:"id"`
	LessonId string `db:"lesson_id"`
	Content  string `db:"content"`
	Order    uint32 `db:"order"`
	Media    string `db:"media"`
	Group    uint32 `db:"group"`
	IsDone   bool   `db:"is_done"`
}

func (ListLessonItem) ColumnQuery() string {
	return `
		u.id,
		u.lesson_id,
		u.content,
		u.order,
		u.media,
		u.group,
		u.is_done
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
	Media    string `db:"media"`
	Group    uint32 `db:"group"`
	IsDone   bool   `db:"is_done"`
}

func (DetailLessonItem) ColumnQuery() string {
	return `
		u.id,
		u.lesson_id,
		u.content,
		u.order,
		u.media,
		u.group,
		u.is_done
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
	Media    string `db:"media"`
	Group    uint32 `db:"group"`
	IsDone   bool   `db:"is_done"`
}

func (CreateLessonItem) InsertQuery() string {
	return `
		INSERT INTO
		lesson_items (
			lesson_id,
			content,
			media,
			"group",
			is_done,
			"order"
		) VALUES (
		 	:lesson_id,
			:content,
			:media,
			:group,
			:is_done,
			:order
		)
	`
}

type UpdateLessonItem struct {
	Id       string `db:"id"`
	LessonId string `db:"lesson_id"`
	Content  string `db:"content"`
	Order    uint32 `db:"order"`
	Media    string `db:"media"`
	Group    uint32 `db:"group"`
	IsDone   bool   `db:"is_done"`
}

func (UpdateLessonItem) InsertQuery() string {
	return `
		UPDATE lesson_items SET
			lesson_id = :lesson_id,
			content = :content,
			media = :media,
			"group" = :group,
			is_done = :is_done,
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

type CopyLessonItem struct {
	Id        string `db:"id"`
	Level     uint32 `db:"level"`
	LevelFrom uint32 `db:"level_from"`
}

func (CopyLessonItem) InsertQuery() string {
	return `
		INSERT INTO lesson_items (lesson_id, content, order, media, "group", is_done)
		SELECT lesson_id, content, order, media, "group", is_done
		FROM lesson_items
		JOIN lessons ON lessons.id = lesson_items.lesson_id
		WHERE lessons.level = :level_from;
	`
}

type BulkCreateLessonItem struct {
	LessonId string `db:"lesson_id"`
	Content  string `db:"content"`
	Order    uint32 `db:"order"`
	Media    string `db:"media"`
	Group    uint32 `db:"group"`
	IsDone   bool   `db:"is_done"`
}

func (BulkCreateLessonItem) InsertQuery() string {
	return `
		INSERT INTO
		lesson_items (
			lesson_id,
			content,
			media,
			"group",
			is_done,
			"order"
		) VALUES (
		 	:lesson_id,
			:content,
			:media,
			:group,
			:is_done,
			:order
		)
	`
}

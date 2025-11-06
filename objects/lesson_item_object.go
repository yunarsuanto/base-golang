package objects

type DetailLessonItemRequest struct {
	Id string
}

type CreateLessonItemRequest struct {
	LessonId string
	Content  string
	Order    uint32
}

type UpdateLessonItemRequest struct {
	Id       string
	LessonId string
	Content  string
	Order    uint32
}

type DeleteLessonItemRequest struct {
	Id string
}

type ListLessonItemResponse struct {
	Id       string
	LessonId string
	Content  string
	Order    uint32
}

type DetailLessonItemResponse struct {
	Id       string
	LessonId string
	Content  string
	Order    uint32
}

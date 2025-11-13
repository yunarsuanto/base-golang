package objects

type DetailLessonItemRequest struct {
	Id string
}

type CreateLessonItemRequest struct {
	LessonId string
	Content  string
	Order    uint32
	Media    string
	Group    uint32
	IsDone   bool
}

type UpdateLessonItemRequest struct {
	Id       string
	LessonId string
	Content  string
	Order    uint32
	Media    string
	Group    uint32
	IsDone   bool
}

type DeleteLessonItemRequest struct {
	Id string
}

type ListLessonItemResponse struct {
	Id       string
	LessonId string
	Content  string
	Order    uint32
	Media    string
	Group    uint32
	IsDone   bool
}

type DetailLessonItemResponse struct {
	Id       string
	LessonId string
	Content  string
	Order    uint32
	Media    string
	Group    uint32
	IsDone   bool
}

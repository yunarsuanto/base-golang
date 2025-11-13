package objects

type DetailLessonRequest struct {
	Id string
}

type CreateLessonRequest struct {
	Title            string
	Description      string
	CategoryLessonId string
	LessonType       string
	Media            string
	Level            uint32
}

type UpdateLessonRequest struct {
	Id               string
	Title            string
	Description      string
	CategoryLessonId string
	LessonType       string
	Media            string
	Level            uint32
}

type DeleteLessonRequest struct {
	Id string
}

type ListLessonResponse struct {
	Id                  string
	Title               string
	Description         string
	CategoryLessonId    string
	LessonType          string
	Media               string
	CategoryLessonTitle string
	Level               uint32
}

type DetailLessonResponse struct {
	Id               string
	Title            string
	Description      string
	CategoryLessonId string
	LessonType       string
	Media            string
	Level            uint32
}

type CopyLessonRequest struct {
	LessonId  string
	Level     uint32
	LevelFrom uint32
}

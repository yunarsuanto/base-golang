package objects

type DetailCategoryLessonRequest struct {
	Id string
}

type CreateCategoryLessonRequest struct {
	Title              string
	Description        string
	CategoryLessonType string
	Media              string
}

type UpdateCategoryLessonRequest struct {
	Id                 string
	Title              string
	Description        string
	CategoryLessonType string
	Media              string
}

type DeleteCategoryLessonRequest struct {
	Id string
}

type ListCategoryLessonResponse struct {
	Id                 string
	Title              string
	Description        string
	CategoryLessonType string
	Media              string
}

type DetailCategoryLessonResponse struct {
	Id                 string
	Title              string
	Description        string
	CategoryLessonType string
	Media              string
}

type ListCategoryLessonPublicResponse struct {
	Id                 string
	Title              string
	Description        string
	CategoryLessonType string
	Media              string
}

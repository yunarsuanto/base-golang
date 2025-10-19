package common_input_handler

type Meta struct {
	Message        string `json:"message"`
	Status         int    `json:"status"`
	PermissionCode string `json:"permissionCode"`
}

type PaginationRequest struct {
	Page   int    `schema:"page" validate:"omitempty,min=1"`
	Limit  int    `schema:"limit" validate:"omitempty,min=1"`
	Search string `schema:"search"`
}

type Pagination struct {
	Page         int `json:"page"`
	Limit        int `json:"limit"`
	Prev         int `json:"prev"`
	Next         int `json:"next"`
	TotalPages   int `json:"totalPages"`
	TotalRecords int `json:"totalRecords"`
}

type ErrorResponse struct {
	Meta Meta `json:"meta"`
}

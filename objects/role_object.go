package objects

type CreateRoleRequest struct {
	Name string
}

type UpdateRoleRequest struct {
	Id   string
	Name string
}

type DeleteRoleRequest struct {
	Id string
}

type DetailRoleRequest struct {
	Id string
}

type ListRoleResponse struct {
	Id   string
	Name string
}

type DetailRoleResponse struct {
	Id             string
	Name           string
	PermissionId   *string
	PermissionName *string
}

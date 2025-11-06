package objects

type UpsertRolePermissionRequest struct {
	RoleId       string
	PermissionId string
}

type DeleteRolePermissionRequest struct {
	RoleId       string
	PermissionId string
}

type ListRolePermissionResponse struct {
	Id   string
	Name string
}

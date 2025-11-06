package objects

type UpsertUserRoleRequest struct {
	UserId   string
	RoleId   string
	IsActive bool
}

type DeleteUserRoleRequest struct {
	UserId string
	RoleId string
}

type ListUserRoleRequest struct {
	UserId string
}

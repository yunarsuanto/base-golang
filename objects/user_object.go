package objects

type CreateUserRequest struct {
	Username string
	Password string
}

type UpdateUserRequest struct {
	Id       string
	Username string
}

type DeleteUserRequest struct {
	Id string
}

type DetailUserRequest struct {
	Id string
}

type ListUserResponse struct {
	Id                string
	Username          string
	Password          string
	IsActive          bool
	ProviderId        *string
	Provider          *string
	TokenVerification *string
}

type DetailUserResponse struct {
	Id             string
	Username       string
	IsActive       bool
	RoleId         *string
	RoleName       *string
	RoleIsActive   *bool
	PermissionId   *string
	PermissionName *string
}

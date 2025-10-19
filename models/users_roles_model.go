package models

const UsersRolesDataName = "peran pengguna"

type GetUsersRoles struct {
	Id                string `db:"id"`
	UserId            string `db:"user_id"`
	UserName          string `db:"user_name"`
	RoleId            string `db:"role_id"`
	RoleName          string `db:"role_name"`
	RoleIsReporter    bool   `db:"role_is_reporter"`
	RoleIsVerificator bool   `db:"role_is_verificator"`
}

type CreateUsersRoles struct {
	UserId string `db:"user_id"`
	RoleId string `db:"role_id"`
}

type GetUsersPermissions struct {
	UserId            string `db:"user_id"`
	UserName          string `db:"user_name"`
	UserEmail         string `db:"user_email"`
	RoleId            string `db:"role_id"`
	RoleName          string `db:"role_name"`
	RoleIsReporter    bool   `db:"role_is_reporter"`
	RoleIsVerificator bool   `db:"role_is_verificator"`
	PermissionId      string `db:"permission_id"`
	PermissionName    string `db:"permission_name"`
	PermissionCode    string `db:"permission_code"`
}

func (GetUsersRoles) ColumnQuery() string {
	return `
			ur.id,
			u.id AS user_id,
			u.name AS user_name,
			r.id AS role_id,
			r.name AS role_name,
			r.is_reporter AS role_is_reporter,
			r.is_verificator AS role_is_verificator
	`
}

func (GetUsersRoles) TableQuery() string {
	return `
		FROM users_roles ur
		JOIN users u ON u.id = ur.user_id
		JOIN roles r ON r.id = ur.role_id
	`
}

func (GetUsersPermissions) ColumnQuery() string {
	return `
		u.id AS user_id,
		u.name AS user_name,
		u.email AS user_email,
		r.id AS role_id,
		r.name AS role_name,
		r.is_reporter AS role_is_reporter,
		r.is_verificator AS role_is_verificator,
		p.id AS permission_id,
		p.name AS permission_name,
		p.code AS permission_code
	`
}

func (GetUsersPermissions) TableQuery() string {
	return `
		FROM roles_permissions rp
		JOIN roles r ON r.id = rp.role_id
		JOIN permissions p ON p.id = rp.permission_id
		JOIN users_roles ur ON ur.role_id = r.id
		JOIN users u ON u.id = ur.user_id
	`
}

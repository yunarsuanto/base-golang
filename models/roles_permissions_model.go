package models

const RolesPermissionsDataName = "izin pengguna"

type GetRolesPermissions struct {
	Id             string `db:"id"`
	RoleId         string `db:"role_id"`
	RoleName       string `db:"role_name"`
	PermissionId   string `db:"permission_id"`
	PermissionName string `db:"permission_name"`
	PermissionCode string `db:"permission_code"`
}

type CreateRolesPermissions struct {
	RoleId       string `db:"role_id"`
	PermissionId string `db:"permission_id"`
}

type GetPermissionCode struct {
	PermissionCode string `db:"permission_code"`
}

func (GetRolesPermissions) ColumnQuery() string {
	return `
		rp.id,
		r.id AS role_id,
		r.name AS role_name,
		p.id AS permission_id,
		p.name AS permission_name,
		p.code AS permission_code
	`
}

func (GetRolesPermissions) TableQuery() string {
	return `
		FROM roles_permissions rp
		JOIN roles r ON r.id = rp.role_id
		JOIN permissions p ON p.id = rp.permission_id
	`
}

func (GetPermissionCode) ColumnQuery() string {
	return `
		DISTINCT(p.code) AS permission_code
	`
}

func (GetPermissionCode) TableQuery() string {
	return `
		FROM roles_permissions rp
		JOIN permissions p ON p.id = rp.permission_id
	`
}

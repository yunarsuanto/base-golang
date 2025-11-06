package models

const RolesPermissionsDataName = "izin pengguna"

type GetPermissionName struct {
	PermissionName string `db:"permission_name"`
}

func (GetPermissionName) ColumnQuery() string {
	return `
		DISTINCT(p.name) AS permission_name
	`
}

func (GetPermissionName) TableQuery() string {
	return `
		FROM role_permission rp
		JOIN permissions p ON p.id = rp.permission_id
		JOIN roles r ON r.id = rp.role_id
		
	`
}

func (GetPermissionName) FilterQuery() string {
	return `
		WHERE r.id = $1 ORDER BY permission_name
	`
}

type UpsertRolePermissionRequest struct {
	RoleId       string `db:"role_id"`
	PermissionId string `db:"permission_id"`
}

func (UpsertRolePermissionRequest) InsertQuery() string {
	return `
		INSERT INTO role_permission (role_id, permission_id)
		VALUES (:role_id, :permission_id)
		ON CONFLICT (role_id, permission_id)
		DO NOTHING;
	`
}

type DeleteRolePermissionRequest struct {
	RoleId       string `db:"role_id"`
	PermissionId string `db:"permission_id"`
}

func (DeleteRolePermissionRequest) InsertQuery() string {
	return `
		DELETE from role_permission WHERE role_id = :role_id AND permission_id = :permission_id;
	`
}

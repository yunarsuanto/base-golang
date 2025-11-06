package models

const UsersRolesDataName = "peran pengguna"

type GetUserRole struct {
	Id       string `db:"id"`
	UserId   string `db:"user_id"`
	UserName string `db:"username"`
	RoleId   string `db:"role_id"`
	RoleName string `db:"role_name"`
}

func (GetUserRole) ColumnQuery() string {
	return `
			ur.id,
			u.id AS user_id,
			u.username,
			r.id AS role_id,
			r.name AS role_name
	`
}

func (GetUserRole) TableQuery() string {
	return `
		FROM user_role ur
		JOIN users u ON u.id = ur.user_id
		LEFT JOIN roles r ON r.id = ur.role_id
	`
}

func (GetUserRole) FilterQuery() string {
	return `
		WHERE u.id = $1 AND ur.is_active = true
	`
}

type UpsertUserRoleRequest struct {
	UserId   string `db:"user_id"`
	RoleId   string `db:"role_id"`
	IsActive bool   `db:"is_active"`
}

func (UpsertUserRoleRequest) InsertQuery() string {
	return `
		INSERT INTO user_role (user_id, role_id, is_active)
		VALUES (:user_id, :role_id, :is_active)
		ON CONFLICT (user_id, role_id)
		DO NOTHING;
	`
}

type DeleteUserRoleRequest struct {
	UserId string `db:"user_id"`
	RoleId string `db:"role_id"`
}

func (DeleteUserRoleRequest) InsertQuery() string {
	return `
		DELETE from user_role WHERE user_id = :user_id AND role_id = :role_id
	`
}

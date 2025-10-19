package user_role_repository

const (
	deleteQuery = `
		DELETE FROM users_roles WHERE user_id = ? %s
	`

	upsertQuery = `
		INSERT IGNORE INTO users_roles (user_id, role_id) VALUES (:user_id, :role_id)
	`
)

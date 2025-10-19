package role_permission_repository

const (
	deleteQuery = `
		DELETE FROM roles_permissions WHERE role_id = ? %s
	`

	upsertQuery = `
		INSERT IGNORE INTO roles_permissions (role_id, permission_id) VALUES (:role_id, :permission_id)
	`

	seedSuperUserQuery = `
		INSERT IGNORE roles_permissions (role_id, permission_id)
		SELECT r.id, p.id
		FROM roles r
		CROSS JOIN permissions p
		WHERE r.name = ?
	`
)

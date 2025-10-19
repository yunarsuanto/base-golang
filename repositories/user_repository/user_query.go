package user_repository

const (
	createQuery = `
		INSERT INTO users (
			id,
			name,
			email,
			username,
			password,
			id_number,
			phone_number,
			profile_picture_path,
			employee_id,
			is_active
		) VALUES (
			:id,
			:name,
			:email,
			:username,
			:password,
			:id_number,
			:phone_number,
			:profile_picture_path,
			:employee_id,
			:is_active
		)
	`

	updateQuery = `
		UPDATE users u SET
			name = :name,
			email = :email,
			username = COALESCE(:username, u.username),
			id_number = COALESCE(:id_number, u.id_number),
			phone_number = :phone_number,
			profile_picture_path = COALESCE(:profile_picture_path, u.profile_picture_path),
			is_active = COALESCE(:is_active, u.is_active)
		WHERE id = :id
	`

	deleteQuery = `
		DELETE FROM users WHERE id = ?
	`

	changePasswordQuery = `
		UPDATE users SET password = :password WHERE id = :id
	`

	updateProfileQuery = `
		UPDATE users SET
			username = :username,
			phone_number = :phone_number,
			profile_picture_path = :profile_picture_path
		WHERE id = :id
	`

	updateActivationQuery = `
		UPDATE users SET is_active = ? WHERE id = ?
	`
)

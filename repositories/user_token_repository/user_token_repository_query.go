package user_token_repository

const (
	upsertQuery = `
		INSERT INTO user_tokens (
			user_id,
			platform,
			fcm_token,
			expiry_time
		) VALUES (
			:user_id,
			:platform,
			:fcm_token,
			:expiry_time
		) AS NEW
		ON DUPLICATE KEY UPDATE
			fcm_token = NEW.fcm_token,
			expiry_time = NEW.expiry_time,
			fcm_is_valid = true
	`

	deleteQuery = `
		DELETE FROM user_tokens
		WHERE user_id = ? AND platform = ?
	`
)

package activity_log_repository

const (
	createQuery = `
		INSERT INTO activity_logs (
			user_id,
			host,
			path,
			body,
			status_code,
			error_message,
			ip_address,
			user_agent,
			memory_usage
		) VALUES (
			:user_id,
			:host,
			:path,
			:body,
			:status_code,
			:error_message,
			:ip_address,
			:user_agent,
			:memory_usage
		)
	`
)

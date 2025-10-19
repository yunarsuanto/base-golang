package models

import (
	"database/sql"
	"time"
)

const ActivityLogDataName = "log aktivitas"

type GetActivityLog struct {
	Id           string    `db:"id"`
	CreatedAt    time.Time `db:"created_at"`
	UserId       *string   `db:"user_id"`
	UserName     *string   `db:"user_name"`
	UserEmail    *string   `db:"user_email"`
	Host         string    `db:"host"`
	Path         string    `db:"path"`
	Body         *string   `db:"body"`
	StatusCode   int       `db:"status_code"`
	ErrorMessage *string   `db:"error_message"`
	IpAddress    string    `db:"ip_address"`
	UserAgent    string    `db:"user_agent"`
	MemoryUsage  float64   `db:"memory_usage"`
}

type CreateActivityLog struct {
	UserId       sql.NullString `db:"user_id"`
	Host         string         `db:"host"`
	Path         string         `db:"path"`
	Body         sql.NullString `db:"body"`
	StatusCode   int            `db:"status_code"`
	ErrorMessage sql.NullString `db:"error_message"`
	IpAddress    string         `db:"ip_address"`
	UserAgent    string         `db:"user_agent"`
	MemoryUsage  float64        `db:"memory_usage"`
}

func (GetActivityLog) ColumnQuery() string {
	return `
		h.id,
		h.name
	`
}

func (GetActivityLog) TableQuery() string {
	return `FROM activity_logs h`
}

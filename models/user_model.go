package models

import "database/sql"

const UserDataName = "pengguna"

type GetUser struct {
	Id                 string  `db:"id"`
	Name               string  `db:"name"`
	Email              string  `db:"email"`
	Username           string  `db:"username"`
	Password           string  `db:"password"`
	IdNumber           *string `db:"id_number"`
	PhoneNumber        *string `db:"phone_number"`
	ProfilePicturePath *string `db:"profile_picture_path"`
	IsActive           bool    `db:"is_active"`
	EmployeeId         *string `db:"employee_id"`
	EmployeeNumber     *string `db:"employee_number"`
	IsReporter         bool    `db:"is_reporter"`
	IsVerificator      bool    `db:"is_verificator"`
}

type CreateUser struct {
	Id                 string         `db:"id"`
	Name               string         `db:"name"`
	Email              string         `db:"email"`
	Username           string         `db:"username"`
	Password           string         `db:"password"`
	IdNumber           sql.NullString `db:"id_number"`
	PhoneNumber        sql.NullString `db:"phone_number"`
	ProfilePicturePath sql.NullString `db:"profile_picture_path"`
	EmployeeId         sql.NullString `db:"employee_id"`
	IsActive           bool           `db:"is_active"`
}

type UpdateUser struct {
	Id                 string         `db:"id"`
	Name               string         `db:"name"`
	Email              string         `db:"email"`
	Username           sql.NullString `db:"username"`
	IdNumber           sql.NullString `db:"id_number"`
	PhoneNumber        sql.NullString `db:"phone_number"`
	ProfilePicturePath sql.NullString `db:"profile_picture_path"`
	IsActive           bool           `db:"is_active"`
}

type ChangeUserPassword struct {
	Id       string `db:"id"`
	Password string `db:"password"`
}

func (GetUser) ColumnQuery() string {
	return `
		u.id,
		u.name,
		u.email,
		u.username,
		u.password,
		u.id_number,
		u.phone_number,
		u.profile_picture_path,
		u.is_active,
		e.id AS employee_id,
		e.employee_number,
		COALESCE(ur.is_reporter, false) AS is_reporter,
		COALESCE(ur.is_verificator, false) AS is_verificator
	`
}

func (GetUser) TableQuery() string {
	return `
		FROM users u
		LEFT JOIN employees e ON e.id = u.employee_id
		LEFT JOIN (
			SELECT 
				u.id AS user_id,
				MAX(r.is_reporter) AS is_reporter,
				MAX(r.is_verificator) AS is_verificator
			FROM users_roles ur
			JOIN users u ON u.id = ur.user_id
			JOIN roles r ON r.id = ur.role_id
			GROUP BY u.id
		) ur ON ur.user_id = u.id
	`
}

package models

const UsersHospitalsDataName = "modul pengguna"

type GetUsersHospitals struct {
	Id           string `db:"id"`
	UserId       string `db:"user_id"`
	UserName     string `db:"user_name"`
	HospitalId   string `db:"hospital_id"`
	HospitalName string `db:"hospital_name"`
}

type CreateUsersHospitals struct {
	UserId     string `db:"user_id"`
	HospitalId string `db:"hospital_id"`
}

func (GetUsersHospitals) ColumnQuery() string {
	return `
		uh.id,
		u.id AS user_id,
		u.name AS user_name,
		h.id AS hospital_id,
		h.name AS hospital_name
	`
}

func (GetUsersHospitals) TableQuery() string {
	return `
		FROM users_hospitals uh
		JOIN users u ON u.id = uh.user_id
		JOIN hospitals h ON h.id = uh.hospital_id
	`
}

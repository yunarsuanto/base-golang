package models

const CompanyDataName = "company"

type ListCompany struct {
	Id       string `db:"id"`
	Code     string `db:"code"`
	Name     string `db:"name"`
	Image    string `db:"image"`
	Email    string `db:"email"`
	IsActive bool   `db:"is_active"`
}

func (ListCompany) ColumnQuery() string {
	return `
		u.id,
		u.code,
		u.name,
		u.image,
		u.email,
		u.is_active
	`
}

func (ListCompany) TableQuery() string {
	return `
		FROM companies u
	`
}

type CreateCompany struct {
	Code            string  `db:"code"`
	Name            string  `db:"name"`
	Image           string  `db:"image"`
	Email           string  `db:"email"`
	Address         string  `db:"address"`
	PostalCode      string  `db:"postal_code"`
	NationalId      string  `db:"national_id"`
	NationalIdImage string  `db:"national_id_image"`
	Npwp            string  `db:"npwp"`
	NpwpImage       string  `db:"npwp_image"`
	DirectorName    string  `db:"director_name"`
	Longitude       float32 `db:"longitude"`
	Latitude        float32 `db:"latitude"`
	IsActive        bool    `db:"is_active"`
}

func (CreateCompany) InsertQuery() string {
	return `
		INSERT INTO
		companies (
			code,
			name,
			image,
			email,
			address,
			postal_code,
			national_id,
			national_id_image,
			npwp,
			npwp_image,
			director_name,
			longitude,
			latitude,
			is_active
		) VALUES (
			:code,
			:name,
			:image,
			:email,
			:address,
			:postal_code,
			:national_id,
			:national_id_image,
			:npwp,
			:npwp_image,
			:director_name,
			:longitude,
			:latitude,
			:is_active
		)
	`
}

type UpdateCompany struct {
	Id              string  `db:"id"`
	Code            string  `db:"code"`
	Name            string  `db:"name"`
	Image           string  `db:"image"`
	Email           string  `db:"email"`
	Address         string  `db:"address"`
	PostalCode      string  `db:"postal_code"`
	NationalId      string  `db:"national_id"`
	NationalIdImage string  `db:"national_id_image"`
	Npwp            string  `db:"npwp"`
	NpwpImage       string  `db:"npwp_image"`
	DirectorName    string  `db:"director_name"`
	Longitude       float32 `db:"longitude"`
	Latitude        float32 `db:"latitude"`
	IsActive        bool    `db:"is_active"`
}

func (UpdateCompany) InsertQuery() string {
	return `
		UPDATE companies SET
			code = :code,
			name = :name,
			image = :image,
			email = :email,
			address = :address,
			postal_code = :postal_code,
			national_id = :national_id,
			national_id_image = :national_id_image,
			npwp = :npwp,
			npwp_image = :npwp_image,
			director_name = :director_name,
			longitude = :longitude,
			latitude = :latitude,
			is_active = :is_active
		WHERE id = :id
	`
}

type DeleteCompany struct {
	Id string `db:"id"`
}

func (DeleteCompany) InsertQuery() string {
	return `
		DELETE FROM companies WHERE id = :id
	`
}

type DetailCompany struct {
	Id              string  `db:"id"`
	Code            string  `db:"code"`
	Name            string  `db:"name"`
	Image           string  `db:"image"`
	Email           string  `db:"email"`
	Address         string  `db:"address"`
	PostalCode      string  `db:"postal_code"`
	NationalId      string  `db:"national_id"`
	NationalIdImage string  `db:"national_id_image"`
	Npwp            string  `db:"npwp"`
	NpwpImage       string  `db:"npwp_image"`
	DirectorName    string  `db:"director_name"`
	Longitude       float32 `db:"longitude"`
	Latitude        float32 `db:"latitude"`
	IsActive        bool    `db:"is_active"`
}

func (DetailCompany) ColumnQuery() string {
	return `
		u.id,
		u.code,
		u.name,
		u.image,
		u.email,
		u.address,
		u.postal_code,
		u.national_id,
		u.national_id_image,
		u.npwp,
		u.npwp_image,
		u.director_name,
		u.longitude,
		u.latitude,
		u.is_active
	`
}

func (DetailCompany) TableQuery() string {
	return `
		FROM companies u
	`
}

func (DetailCompany) FilterQuery() string {
	return `
		WHERE u.id = $1
	`
}

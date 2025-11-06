package models

const UserDataName = "pengguna"

type ListUser struct {
	Id                string  `db:"id"`
	Username          string  `db:"username"`
	Password          string  `db:"password"`
	IsActive          bool    `db:"is_active"`
	ProviderId        *string `db:"provider_id"`
	Provider          *string `db:"provider"`
	TokenVerification *string `db:"token_verification"`
}

func (ListUser) ColumnQuery() string {
	return `
		u.id,
		u.username,
		u.password,
		u.is_active,
		u.provider_id,
		u.provider,
		u.token_verification
	`
}

func (ListUser) TableQuery() string {
	return `
		FROM users u
	`
}

type CreateUser struct {
	Username   string `db:"username"`
	Password   string `db:"password"`
	IsActive   bool   `db:"is_active"`
	ProviderId string `db:"provider_id"`
	Provider   string `db:"provider"`
}

func (CreateUser) InsertQuery() string {
	return `
		INSERT INTO
		users (
			username,
			password,
			is_active,
			provider_id,
			provider
		) VALUES (
			:username,
			:password,
			:is_active ,
			:provider_id,
			:provider
		)
	`
}

type UpdateUser struct {
	Id         string  `db:"id"`
	Username   string  `db:"username"`
	ProviderId *string `db:"provider_id"`
}

func (UpdateUser) InsertQuery() string {
	return `
		UPDATE users SET
			username = :username,
			provider_id = :provider_id
		WHERE id = :id
	`
}

type DeleteUser struct {
	Id string `db:"id"`
}

func (DeleteUser) InsertQuery() string {
	return `
		DELETE FROM users WHERE id = :id
	`
}

type DetailUser struct {
	Id             string  `db:"id"`
	Username       string  `db:"username"`
	IsActive       bool    `db:"is_active"`
	RoleId         *string `db:"role_id"`
	RoleName       *string `db:"role_name"`
	RoleIsActive   *bool   `db:"role_is_active"`
	PermissionId   *string `db:"permission_id"`
	PermissionName *string `db:"permission_name"`
}

func (DetailUser) ColumnQuery() string {
	return `
		u.id,
		u.username,
		u.is_active,
		r.id as role_id,
		r.name as role_name,
		ur.is_active as role_is_active,
		p.id as permission_id,
		p.name as permission_name
	`
}

func (DetailUser) TableQuery() string {
	return `
		FROM users u
		LEFT JOIN user_role ur ON ur.user_id = u.id
		LEFT JOIN roles r ON r.id = ur.role_id
		LEFT JOIN role_permission rp ON rp.role_id = r.id
		LEFT JOIN permissions p ON p.id = rp.permission_id
	`
}

func (DetailUser) FilterQuery() string {
	return `
		WHERE u.id = $1
	`
}

type CreateUserProfile struct {
	UserId          string  `db:"user_id"`
	NationalId      string  `db:"national_id"`
	Fullname        string  `db:"fullname"`
	Email           string  `db:"email"`
	Phone           string  `db:"phone"`
	Address         string  `db:"address"`
	PostalCode      string  `db:"postal_code"`
	Age             uint16  `db:"age"`
	Latitude        float64 `db:"latitude"`
	Longitude       float64 `db:"longitude"`
	ProfileImage    string  `db:"profile_image"`
	NationalIdImage string  `db:"national_id_image"`
	GuardianName    string  `db:"guardian_name"`
}

func (CreateUserProfile) InsertQuery() string {
	return `
		INSERT INTO
		user_profiles (
			user_id,
			national_id,
			fullname,
			email,
			phone,
			address,
			postal_code,
			age,
			latitude,
			longitude,
			profile_image,
			national_id_image,
			guardian_name
		) VALUES (
			:user_id,
			:national_id,
			:fullname,
			:email,
			:phone,
			:address,
			:postal_code,
			:age,
			:latitude,
			:longitude,
			:profile_image,
			:national_id_image,
			:guardian_name
		)
	`
}

type UpdateUserTokenVerification struct {
	Username string `db:"username"`
	Token    string `db:"token_verification"`
}

func (UpdateUserTokenVerification) InsertQuery() string {
	return `
		UPDATE users SET token_verification = :token_verification WHERE username = :username
	`
}

type UpdateUserIsActiveTokenVerification struct {
	Username string `db:"username"`
}

func (UpdateUserIsActiveTokenVerification) InsertQuery() string {
	return `
		UPDATE users SET token_verification = null, is_active = true WHERE username = :username
	`
}

package models

const RoleDataName = "role"

type ListRole struct {
	Id   string `db:"id"`
	Name string `db:"name"`
}

func (ListRole) ColumnQuery() string {
	return `
		u.id,
		u.name
	`
}

func (ListRole) TableQuery() string {
	return `
		FROM roles u
	`
}

type CreateRole struct {
	Name string `db:"name"`
}

func (CreateRole) InsertQuery() string {
	return `
		INSERT INTO
		roles (
			name
		) VALUES (
			:name
		)
	`
}

type UpdateRole struct {
	Id   string `db:"id"`
	Name string `db:"name"`
}

func (UpdateRole) InsertQuery() string {
	return `
		UPDATE roles SET
			name = :name
		WHERE id = :id
	`
}

type DeleteRole struct {
	Id string `db:"id"`
}

func (DeleteRole) InsertQuery() string {
	return `
		DELETE FROM roles WHERE id = :id
	`
}

type DetailRole struct {
	Id             string  `db:"id"`
	Name           string  `db:"name"`
	PermissionId   *string `db:"permission_id"`
	PermissionName *string `db:"permission_name"`
}

func (DetailRole) ColumnQuery() string {
	return `
		u.id,
		u.name,
		p.id as permission_id,
		p.name as permission_name
	`
}

func (DetailRole) TableQuery() string {
	return `
		FROM roles u
		LEFT JOIN role_permission rp ON rp.role_id = u.id
		LEFT JOIN permissions p ON p.id = rp.permission_id
	`
}

func (DetailRole) FilterQuery() string {
	return `
		WHERE u.id = $1
	`
}

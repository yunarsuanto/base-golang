package models

const PermissionDataName = "permission"

type ListPermission struct {
	Id   string `db:"id"`
	Name string `db:"name"`
}

func (ListPermission) ColumnQuery() string {
	return `
				u.id,
				u.name
			`
}

func (ListPermission) TableQuery() string {
	return `
				FROM permissions u
			`
}

type CreatePermission struct {
	Name string `db:"name"`
}

func (CreatePermission) InsertQuery() string {
	return `
				INSERT INTO
				permissions (
					name
				) VALUES (
					:name
				)
			`
}

type UpdatePermission struct {
	Id   string `db:"id"`
	Name string `db:"name"`
}

func (UpdatePermission) InsertQuery() string {
	return `
				UPDATE permissions SET
					name = :name
				WHERE id = :id
			`
}

type DeletePermission struct {
	Id string `db:"id"`
}

func (DeletePermission) InsertQuery() string {
	return `
				DELETE FROM permissions WHERE id = :id
			`
}

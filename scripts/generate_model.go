package scripts

import (
	"fmt"
	"strings"

	"github.com/yunarsuanto/base-go/utils"
)

func GenerateModel(name string) {
	capitalName := utils.ToPascalCase(name)
	template := fmt.Sprintf(`
		package models

		const %sDataName = "%s"

		type List%s struct {
			Id       string __BACKTICK__db:"id"__BACKTICK__
			Name string __BACKTICK__db:"name"__BACKTICK__
		}

		func (List%s) ColumnQuery() string {
			return __BACKTICK__
				u.id,
				u.name
			__BACKTICK__
		}

		func (List%s) TableQuery() string {
			return __BACKTICK__
				FROM %ss u
			__BACKTICK__
		}

		type Create%s struct {
			Name string __BACKTICK__db:"name"__BACKTICK__
		}

		func (Create%s) InsertQuery() string {
			return __BACKTICK__
				INSERT INTO
				%ss (
					name
				) VALUES (
					:name,
				)
			__BACKTICK__
		}

		type Update%s struct {
			Id       string __BACKTICK__db:"id"__BACKTICK__
			Name string __BACKTICK__db:"name"__BACKTICK__
		}

		func (Update%s) InsertQuery() string {
			return __BACKTICK__
				UPDATE %ss SET
					name = :name
				WHERE id = :id
			__BACKTICK__
		}

		type Delete%s struct {
			Id string __BACKTICK__db:"id"__BACKTICK__
		}

		func (Delete%s) InsertQuery() string {
			return __BACKTICK__
				DELETE FROM %ss WHERE id = :id
			__BACKTICK__
		}


	`,
		capitalName,
		name,
		capitalName,
		capitalName,
		capitalName,
		name,
		capitalName,
		capitalName,
		name,
		capitalName,
		capitalName,
		name,
		capitalName,
		capitalName,
		name,
	)

	code := strings.ReplaceAll(template, "__BACKTICK__", "`")
	filePath := fmt.Sprintf("models/%s_model.go", name)

	save(filePath, code)
}

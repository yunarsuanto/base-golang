package scripts

import (
	"fmt"
	"strings"

	"github.com/yunarsuanto/base-go/utils"
)

func GenerateRepositoryInterface(name string) {
	capitalName := utils.ToPascalCase(name)
	template := fmt.Sprintf(`
		package %s_repository

		import (
			"context"

			"github.com/jmoiron/sqlx"
			"github.com/yunarsuanto/base-go/constants"
			"github.com/yunarsuanto/base-go/models"
			"github.com/yunarsuanto/base-go/objects"
		)

		type %sRepositoryInterface interface {
			List%s(ctx context.Context, tx *sqlx.Tx, pagination *objects.Pagination) ([]models.List%s, *constants.ErrorResponse)
			Create%s(ctx context.Context, tx *sqlx.Tx, data models.Create%s) *constants.ErrorResponse
			Update%s(ctx context.Context, tx *sqlx.Tx, data models.Update%s) *constants.ErrorResponse
			Delete%s(ctx context.Context, tx *sqlx.Tx, data models.Delete%s) *constants.ErrorResponse
		}

		func New%sRepository() %sRepositoryInterface {
			return &repository{}
		}


	`,
		name,
		capitalName,
		capitalName,
		capitalName,
		capitalName,
		capitalName,
		capitalName,
		capitalName,
		capitalName,
		capitalName,
		capitalName,
		capitalName,
	)

	code := strings.ReplaceAll(template, "__BACKTICK__", "`")
	filePath := fmt.Sprintf("repositories/%s_repository/%s_repository_interface.go", name, name)

	save(filePath, code)
}

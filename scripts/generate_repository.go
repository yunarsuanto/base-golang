package scripts

import (
	"fmt"
	"strings"

	"github.com/yunarsuanto/base-go/utils"
)

func GenerateRepository(name string) {
	capitalName := utils.ToPascalCase(name)
	template := fmt.Sprintf(`
		package %s_repository

		import (
			"context"
			"fmt"

			"github.com/jmoiron/sqlx"
			"github.com/yunarsuanto/base-go/constants"
			"github.com/yunarsuanto/base-go/models"
			"github.com/yunarsuanto/base-go/objects"
			"github.com/yunarsuanto/base-go/utils"
		)

		type repository struct{}

		func (repository) List%s(ctx context.Context, tx *sqlx.Tx, pagination *objects.Pagination) ([]models.List%s, *constants.ErrorResponse) {
			var result []models.List%s
			var query models.List%s

			params := []any{}

			filterQuery, err := mapQueryFilterList%s(pagination.Search, &params)
			if err != nil {
				return result, utils.ErrDatabase(err, models.%sDataName)
			}

			getQuery := fmt.Sprintf("SELECT %%s %%s %%s", query.ColumnQuery(), query.TableQuery(), filterQuery)
			countQuery := fmt.Sprintf("SELECT COUNT(1) %%s %%s", query.TableQuery(), filterQuery)
			if errs := utils.QueryOperation(&getQuery, [][2]string{{"u.name", constants.Ascending}}, pagination.Limit, pagination.Page); errs != nil {
				return result, errs
			}

			err = tx.SelectContext(
				ctx,
				&result,
				getQuery,
				params...,
			)
			if err != nil {
				return result, utils.ErrDatabase(err, models.%sDataName)
			}

			var count int
			err = tx.QueryRowContext(
				ctx,
				countQuery,
				params...,
			).Scan(&count)
			if err != nil {
				return result, utils.ErrDatabase(err, models.%sDataName)
			}

			pagination.GetPagination(count)
			return result, nil
		}

		func (repository) Create%s(ctx context.Context, tx *sqlx.Tx, data models.Create%s) *constants.ErrorResponse {
			_, err := tx.NamedExecContext(
				ctx,
				data.InsertQuery(),
				data,
			)
			if err != nil {
				return utils.ErrDatabase(err, models.%sDataName)
			}

			return nil
		}
		func (repository) Update%s(ctx context.Context, tx *sqlx.Tx, data models.Update%s) *constants.ErrorResponse {
			_, err := tx.NamedExecContext(
				ctx,
				data.InsertQuery(),
				data,
			)
			if err != nil {
				return utils.ErrDatabase(err, models.%sDataName)
			}

			return nil
		}
		func (repository) Delete%s(ctx context.Context, tx *sqlx.Tx, data models.Delete%s) *constants.ErrorResponse {
			_, err := tx.NamedExecContext(
				ctx,
				data.InsertQuery(),
				data,
			)
			if err != nil {
				return utils.ErrDatabase(err, models.%sDataName)
			}

			return nil
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
		capitalName,
		capitalName,
		capitalName,
		capitalName,
		capitalName,
		capitalName,
	)

	code := strings.ReplaceAll(template, "__BACKTICK__", "`")
	filePath := fmt.Sprintf("repositories/%s_repository/%s_repository.go", name, name)

	save(filePath, code)
}

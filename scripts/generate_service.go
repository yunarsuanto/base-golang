package scripts

import (
	"fmt"
	"strings"

	"github.com/yunarsuanto/base-go/utils"
)

func GenerateService(name string) {
	capitalName := utils.ToPascalCase(name)
	template := fmt.Sprintf(`
		package %s_service

		import (
			"context"

			"github.com/yunarsuanto/base-go/constants"
			"github.com/yunarsuanto/base-go/infra/initiator/infra"
			"github.com/yunarsuanto/base-go/infra/initiator/repository"
			"github.com/yunarsuanto/base-go/models"
			"github.com/yunarsuanto/base-go/objects"
			"github.com/yunarsuanto/base-go/utils"
		)

		type service struct {
			*repository.RepoCtx
			*infra.InfraCtx
		}

		func (a service) List%s(ctx context.Context, pagination *objects.Pagination) ([]objects.List%sResponse, *constants.ErrorResponse) {
			var result []objects.List%sResponse

			tx, err := a.Db.Begin(ctx)
			if err != nil {
				return result, utils.ErrorInternalServer(err.Error())
			}

			resultData, errs := a.%sRepo.List%s(ctx, tx, pagination)
			if errs != nil {
				_ = tx.Rollback()
				return result, errs
			}

			for _, v := range resultData {
				result = append(result, objects.List%sResponse(v))
			}

			err = tx.Commit()
			if err != nil {
				_ = tx.Rollback()
				return result, utils.ErrorInternalServer(err.Error())
			}

			return result, nil
		}

		func (a service) Create%s(ctx context.Context, req objects.Create%sRequest) *constants.ErrorResponse {
			tx, err := a.Db.Begin(ctx)
			if err != nil {
				return utils.ErrorInternalServer(err.Error())
			}

			createData := models.Create%s{
				Name: req.Name,
			}

			errs := a.%sRepo.Create%s(ctx, tx, createData)
			if errs != nil {
				_ = tx.Rollback()
				return errs
			}

			err = tx.Commit()
			if err != nil {
				_ = tx.Rollback()
				return utils.ErrorInternalServer(err.Error())
			}

			return nil
		}
		func (a service) Update%s(ctx context.Context, req objects.Update%sRequest) *constants.ErrorResponse {
			tx, err := a.Db.Begin(ctx)
			if err != nil {
				return utils.ErrorInternalServer(err.Error())
			}

			updateData := models.Update%s{
				Id:       req.Id,
				Name: 	req.Name,
			}

			errs := a.%sRepo.Update%s(ctx, tx, updateData)
			if errs != nil {
				_ = tx.Rollback()
				return errs
			}

			err = tx.Commit()
			if err != nil {
				_ = tx.Rollback()
				return utils.ErrorInternalServer(err.Error())
			}

			return nil
		}
		func (a service) Delete%s(ctx context.Context, req objects.Delete%sRequest) *constants.ErrorResponse {
			tx, err := a.Db.Begin(ctx)
			if err != nil {
				return utils.ErrorInternalServer(err.Error())
			}

			deleteData := models.Delete%s{
				Id: req.Id,
			}

			errs := a.%sRepo.Delete%s(ctx, tx, deleteData)
			if errs != nil {
				_ = tx.Rollback()
				return errs
			}

			err = tx.Commit()
			if err != nil {
				_ = tx.Rollback()
				return utils.ErrorInternalServer(err.Error())
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
		capitalName,
		capitalName,
		capitalName,
		capitalName,
	)

	code := strings.ReplaceAll(template, "__BACKTICK__", "`")
	filePath := fmt.Sprintf("services/%s_service/%s_service.go", name, name)

	save(filePath, code)
}

package scripts

import (
	"fmt"
	"strings"

	"github.com/yunarsuanto/base-go/utils"
)

func GenerateServiceInterface(name string) {
	capitalName := utils.ToPascalCase(name)
	template := fmt.Sprintf(`
		package %s_service

		import (
			"context"

			"github.com/yunarsuanto/base-go/constants"
			"github.com/yunarsuanto/base-go/infra/initiator/infra"
			"github.com/yunarsuanto/base-go/infra/initiator/repository"
			"github.com/yunarsuanto/base-go/objects"
		)

		type %sServiceInterface interface {
			List%s(ctx context.Context, pagination *objects.Pagination) ([]objects.List%sResponse, *constants.ErrorResponse)
			Create%s(ctx context.Context, req objects.Create%sRequest) *constants.ErrorResponse
			Update%s(ctx context.Context, req objects.Update%sRequest) *constants.ErrorResponse
			Delete%s(ctx context.Context, req objects.Delete%sRequest) *constants.ErrorResponse
		}

		func New%sService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) %sServiceInterface {
			return &service{
				repoCtx,
				infraCtx,
			}
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
	filePath := fmt.Sprintf("services/%s_service/%s_service_interface.go", name, name)

	save(filePath, code)
}

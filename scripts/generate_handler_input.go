package scripts

import (
	"fmt"
	"strings"

	"github.com/yunarsuanto/base-go/utils"
)

func GenerateHandlerInput(name string) {
	capitalName := utils.ToPascalCase(name)
	template := fmt.Sprintf(`
		package %s_handler

		import (
			"github.com/yunarsuanto/base-go/constants"
			common_input_handler "github.com/yunarsuanto/base-go/handler"
			"github.com/yunarsuanto/base-go/objects"
			"github.com/yunarsuanto/base-go/utils"
			"golang.org/x/net/context"
		)

		type list%sRequest struct {
			common_input_handler.PaginationRequest
		}

		type list%sResponse struct {
			Meta       common_input_handler.Meta        __BACKTICK__json:"meta"__BACKTICK__
			Pagination *common_input_handler.Pagination __BACKTICK__json:"pagination"__BACKTICK__
			Data       []*list%sResponseData          __BACKTICK__json:"data"__BACKTICK__
		}

		type list%sResponseData struct {
			Id       string __BACKTICK__json:"id"__BACKTICK__
			Name string __BACKTICK__json:"name"__BACKTICK__
		}

		type create%sRequest struct {
			Name string __BACKTICK__json:"name" schema:"name" validate:"required"__BACKTICK__
		}

		type create%sResponse struct {
			Meta common_input_handler.Meta __BACKTICK__json:"meta"__BACKTICK__
			Data *create%sResponseData   __BACKTICK__json:"data"__BACKTICK__
		}

		type create%sResponseData struct {
		}

		type update%sRequest struct {
			Id       string __BACKTICK__json:"id" schema:"id" validate:"required,uuid"__BACKTICK__
			Name string __BACKTICK__json:"name" schema:"name" validate:"required"__BACKTICK__
		}

		type update%sResponse struct {
			Meta common_input_handler.Meta __BACKTICK__json:"meta"__BACKTICK__
			Data *update%sResponseData   __BACKTICK__json:"data"__BACKTICK__
		}

		type update%sResponseData struct {
		}

		type delete%sRequest struct {
			Id string __BACKTICK__json:"id" schema:"id" validate:"required,uuid"__BACKTICK__
		}

		type delete%sResponse struct {
			Meta common_input_handler.Meta __BACKTICK__json:"meta"__BACKTICK__
			Data *delete%sResponseData   __BACKTICK__json:"data"__BACKTICK__
		}

		type delete%sResponseData struct {
		}

		func (a handler) checkPermission(ctx context.Context, permission string) *constants.ErrorResponse {
			claims, ok := ctx.Value(constants.ClaimsContextKey).(*objects.JWTClaims)
			if !ok || claims == nil {
				return constants.ErrTokenInvalid
			}

			if !claims.IsSuperAdmin {
				if !utils.InArrayExist(permission, claims.Permissions) {
					return constants.ErrIneligibleAccess
				}
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
	)

	code := strings.ReplaceAll(template, "__BACKTICK__", "`")
	filePath := fmt.Sprintf("handler/%s_handler/%s_handler_input.go", name, name)

	save(filePath, code)
}

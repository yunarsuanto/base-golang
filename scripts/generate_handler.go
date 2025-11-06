package scripts

import (
	"fmt"
	"strings"

	"github.com/yunarsuanto/base-go/utils"
)

func GenerateHandler(name string) {
	capitalName := utils.ToPascalCase(name)
	template := fmt.Sprintf(`
		package %s_handler

		import (
			"net/http"

			"github.com/yunarsuanto/base-go/constants"
			"github.com/yunarsuanto/base-go/infra/initiator/service"
			"github.com/yunarsuanto/base-go/objects"
			"github.com/yunarsuanto/base-go/utils"
		)

		type handler struct {
			*service.ServiceCtx
		}

		func (a handler) List%s(w http.ResponseWriter, r *http.Request) {
			var result list%sResponse
			permission := constants.Permission%sList

			ctx := r.Context()
			errs := a.checkPermission(ctx, permission)
			if errs != nil {
				result = list%sResponse{Meta: utils.SetErrorMeta(errs, permission)}
				utils.JSONResponse(w, errs.HttpCode, &result)
				return
			}

			var in list%sRequest
			defer func() {
				a.ActivityLogService.Create(ctx, objects.CreateActivityLog{
					Request:      r,
					Body:         utils.MaskBody(&in),
					ResponseMeta: result.Meta,
				})
			}()

			errs = utils.DecodeUrlQueryParams(&in, r.URL.Query())
			if errs != nil {
				result = list%sResponse{Meta: utils.SetErrorMeta(errs, permission)}
				utils.JSONResponse(w, errs.HttpCode, &result)
				return
			}

			pagination := objects.NewPagination()
			pagination.MapFromRequest(in.PaginationRequest)
			data, errs := a.%sService.List%s(ctx, pagination)
			if errs != nil {
				result = list%sResponse{Meta: utils.SetErrorMeta(errs, permission)}
				utils.JSONResponse(w, errs.HttpCode, &result)
				return
			}

			resultData := make([]*list%sResponseData, len(data))
			for i, v := range data {
				resultData[i] = &list%sResponseData{
					Id:       v.Id,
					Name: v.Name,
				}
			}

			result = list%sResponse{
				Meta:       utils.SetSuccessMeta("List %s", permission),
				Pagination: pagination.MapToResponse(),
				Data:       resultData,
			}

			utils.JSONResponse(w, result.Meta.Status, &result)
		}

		func (a handler) Create%s(w http.ResponseWriter, r *http.Request) {
			var result create%sResponse
			permission := constants.Permission%sCreate

			ctx := r.Context()
			errs := a.checkPermission(ctx, permission)
			if errs != nil {
				result = create%sResponse{Meta: utils.SetErrorMeta(errs, permission)}
				utils.JSONResponse(w, errs.HttpCode, &result)
				return
			}
			var in create%sRequest
			defer func() {
				a.ActivityLogService.Create(ctx, objects.CreateActivityLog{
					Request:      r,
					Body:         utils.MaskBody(&in),
					ResponseMeta: result.Meta,
				})
			}()

			errs = utils.DecodeJson(&in, r.Body)
			if errs != nil {
				result = create%sResponse{Meta: utils.SetErrorMeta(errs, permission)}
				utils.JSONResponse(w, errs.HttpCode, &result)
				return
			}

			req := objects.Create%sRequest(in)
			errs = a.%sService.Create%s(ctx, req)
			if errs != nil {
				result = create%sResponse{Meta: utils.SetErrorMeta(errs, permission)}
				utils.JSONResponse(w, errs.HttpCode, &result)
				return
			}

			result = create%sResponse{
				Meta: utils.SetSuccessMeta("Create %s", permission),
			}
			utils.JSONResponse(w, result.Meta.Status, &result)
		}
		func (a handler) Update%s(w http.ResponseWriter, r *http.Request) {
			var result update%sResponse
			permission := constants.Permission%sUpdate

			ctx := r.Context()
			errs := a.checkPermission(ctx, permission)
			if errs != nil {
				result = update%sResponse{Meta: utils.SetErrorMeta(errs, permission)}
				utils.JSONResponse(w, errs.HttpCode, &result)
				return
			}
			var in update%sRequest
			defer func() {
				a.ActivityLogService.Create(ctx, objects.CreateActivityLog{
					Request:      r,
					Body:         utils.MaskBody(&in),
					ResponseMeta: result.Meta,
				})
			}()

			errs = utils.DecodeJson(&in, r.Body)
			if errs != nil {
				result = update%sResponse{Meta: utils.SetErrorMeta(errs, permission)}
				utils.JSONResponse(w, errs.HttpCode, &result)
				return
			}

			req := objects.Update%sRequest(in)
			errs = a.%sService.Update%s(ctx, req)
			if errs != nil {
				result = update%sResponse{Meta: utils.SetErrorMeta(errs, permission)}
				utils.JSONResponse(w, errs.HttpCode, &result)
				return
			}

			result = update%sResponse{
				Meta: utils.SetSuccessMeta("Update %s", permission),
			}
			utils.JSONResponse(w, result.Meta.Status, &result)
		}
		func (a handler) Delete%s(w http.ResponseWriter, r *http.Request) {
			var result delete%sResponse
			permission := constants.Permission%sDelete

			ctx := r.Context()
			errs := a.checkPermission(ctx, permission)
			if errs != nil {
				result = delete%sResponse{Meta: utils.SetErrorMeta(errs, permission)}
				utils.JSONResponse(w, errs.HttpCode, &result)
				return
			}
			var in delete%sRequest
			defer func() {
				a.ActivityLogService.Create(ctx, objects.CreateActivityLog{
					Request:      r,
					Body:         utils.MaskBody(&in),
					ResponseMeta: result.Meta,
				})
			}()

			errs = utils.DecodeJson(&in, r.Body)
			if errs != nil {
				result = delete%sResponse{Meta: utils.SetErrorMeta(errs, permission)}
				utils.JSONResponse(w, errs.HttpCode, &result)
				return
			}

			req := objects.Delete%sRequest(in)
			errs = a.%sService.Delete%s(ctx, req)
			if errs != nil {
				result = delete%sResponse{Meta: utils.SetErrorMeta(errs, permission)}
				utils.JSONResponse(w, errs.HttpCode, &result)
				return
			}

			result = delete%sResponse{
				Meta: utils.SetSuccessMeta("Delete %s", permission),
			}
			utils.JSONResponse(w, result.Meta.Status, &result)
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
		capitalName,
		capitalName,
		capitalName,
		capitalName,
		capitalName,
		capitalName,
		capitalName,
	)

	code := strings.ReplaceAll(template, "__BACKTICK__", "`")
	filePath := fmt.Sprintf("handler/%s_handler/%s_handler.go", name, name)

	save(filePath, code)
}

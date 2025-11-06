package file_handler

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

func (a handler) UploadMultipart(w http.ResponseWriter, r *http.Request) {
	var result uploadFileResponse

	ctx := r.Context()
	file, header, err := r.FormFile(constants.MultipartFormFileName)
	if err != nil {
		errs := utils.ErrorInternalServer(err.Error())
		result = uploadFileResponse{Meta: utils.SetErrorMeta(errs)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}
	defer file.Close()

	data := objects.FileUploadMultipartRequest{
		File:       file,
		FileHeader: header,
		Context:    r.FormValue(constants.MultipartFormContextName),
	}
	resultData, errs := a.FileService.UploadMultipart(ctx, data)
	if errs != nil {
		result = uploadFileResponse{Meta: utils.SetErrorMeta(errs)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = uploadFileResponse{
		Meta: utils.SetSuccessMeta("Upload File Multipart"),
		Data: &uploadFileResponseData{
			Url:      resultData.Url,
			Path:     resultData.Path,
			MimeType: resultData.MimeType,
			Size:     resultData.Size,
		},
	}
	utils.JSONResponse(w, result.Meta.Status, &result)
}

func (a handler) UploadBase64(w http.ResponseWriter, r *http.Request) {
	var result uploadFileResponse
	permission := constants.PermissionUserCreate

	ctx := r.Context()
	errs := a.checkPermission(ctx, permission)
	if errs != nil {
		result = uploadFileResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}
	var in uploadBase64Request
	errs = utils.DecodeJson(&in, r.Body)
	if errs != nil {
		result = uploadFileResponse{Meta: utils.SetErrorMeta(errs)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	data := objects.FileUploadBase64Request{
		File:    in.File,
		Context: in.Context,
	}
	resultData, errs := a.FileService.UploadBase64(ctx, data)
	if errs != nil {
		result = uploadFileResponse{Meta: utils.SetErrorMeta(errs)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = uploadFileResponse{
		Meta: utils.SetSuccessMeta("Upload File Base64"),
		Data: &uploadFileResponseData{
			Url:      resultData.Url,
			Path:     resultData.Path,
			MimeType: resultData.MimeType,
			Size:     resultData.Size,
		},
	}
	utils.JSONResponse(w, result.Meta.Status, &result)
}

func (a handler) GetUrl(w http.ResponseWriter, r *http.Request) {
	var result getUrlResponse

	ctx := r.Context()
	var in getUrlRequest
	errs := utils.DecodeUrlQueryParams(&in, r.URL.Query())
	if errs != nil {
		result = getUrlResponse{Meta: utils.SetErrorMeta(errs)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	data := objects.FileGetUrlRequest{
		Path: in.Path,
	}
	resultData, errs := a.FileService.GetUrl(ctx, data)
	if errs != nil {
		result = getUrlResponse{Meta: utils.SetErrorMeta(errs)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = getUrlResponse{
		Meta: utils.SetSuccessMeta("Get File URL"),
		Data: &getUrlResponseData{
			Url: resultData.Url,
		},
	}
	utils.JSONResponse(w, result.Meta.Status, &result)
}

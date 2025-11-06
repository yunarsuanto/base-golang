package file_handler

import (
	"github.com/yunarsuanto/base-go/constants"
	common_input_handler "github.com/yunarsuanto/base-go/handler"
	"github.com/yunarsuanto/base-go/objects"
	"github.com/yunarsuanto/base-go/utils"
	"golang.org/x/net/context"
)

type uploadBase64Request struct {
	File    string `json:"file"`
	Context string `json:"context"`
}

type uploadFileResponseData struct {
	Url      string `json:"url"`
	Path     string `json:"path"`
	MimeType string `json:"mimeType"`
	Size     int64  `json:"size"`
}

type uploadFileResponse struct {
	Meta common_input_handler.Meta `json:"meta"`
	Data *uploadFileResponseData   `json:"data"`
}

type getUrlRequest struct {
	Path string `schema:"path" validate:"required"`
}

type getUrlResponseData struct {
	Url string `json:"url"`
}

type getUrlResponse struct {
	Meta common_input_handler.Meta `json:"meta"`
	Data *getUrlResponseData       `json:"data"`
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

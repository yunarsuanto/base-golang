package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/schema"
	"github.com/yunarsuanto/base-go/constants"
)

func notEmail(fl validator.FieldLevel) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return !emailRegex.MatchString(fl.Field().String())
}

func DecodeUrlQueryParams(in any, req url.Values) *constants.ErrorResponse {
	err := schema.NewDecoder().Decode(in, req)
	if err != nil {
		return ErrorInternalServer(err.Error())
	}
	err = validator.New().Struct(in)
	if err != nil {
		return constants.Error(http.StatusBadRequest, err.Error())
	}

	return nil
}

func DecodeJson(in any, req io.ReadCloser) *constants.ErrorResponse {
	err := json.NewDecoder(req).Decode(&in)
	if err != nil {
		return ErrorInternalServer(err.Error())
	}
	validate := validator.New()
	validate.RegisterValidation("notemail", notEmail)
	err = validate.Struct(in)
	if err != nil {
		return constants.Error(http.StatusBadRequest, err.Error())
	}

	return nil
}

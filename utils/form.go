package utils

import (
	"net/http"
	"slices"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/yunarsuanto/base-go/constants"
)

func ValidateFormDefaultValue(inputType, defaultValue string, enums []string) *constants.ErrorResponse {
	if defaultValue == "" {
		return nil
	}

	switch inputType {
	case constants.FormInputTypeDate:
		if defaultValue != constants.FormDateDefaultValueCurrentDate {
			return constants.Error(http.StatusBadRequest, constants.ErrInvalidInputType, defaultValue, inputType)
		}
	case constants.FormInputTypeTime:
		if defaultValue != constants.FormDateDefaultValueCurrentTime {
			return constants.Error(http.StatusBadRequest, constants.ErrInvalidInputType, defaultValue, inputType)
		}
	case constants.FormInputTypeDatetime:
		if defaultValue != constants.FormDateDefaultValueCurrentDateTime {
			return constants.Error(http.StatusBadRequest, constants.ErrInvalidInputType, defaultValue, inputType)
		}
	case constants.FormInputTypeVarchar:
		break
	case constants.FormInputTypeText:
		break
	case constants.FormInputTypeInteger:
		_, err := strconv.Atoi(defaultValue)
		if err != nil {
			return constants.Error(http.StatusBadRequest, constants.ErrInvalidInputType, defaultValue, inputType)
		}
	case constants.FormInputTypeDecimal:
		_, err := strconv.ParseFloat(defaultValue, 64)
		if err != nil {
			return constants.Error(http.StatusBadRequest, constants.ErrInvalidInputType, defaultValue, inputType)
		}
	case constants.FormInputTypeEmail:
		err := validator.New().Struct(struct {
			value string `validate:"email"`
		}{
			value: defaultValue,
		})
		if err != nil {
			return constants.Error(http.StatusBadRequest, constants.ErrInvalidInputType, defaultValue, inputType)
		}
	case constants.FormInputTypePhone:
		err := validator.New().Struct(struct {
			value string `validate:"e164"`
		}{
			value: defaultValue,
		})
		if err != nil {
			return constants.Error(http.StatusBadRequest, constants.ErrInvalidInputType, defaultValue, inputType)
		}
	case constants.FormInputTypeIdNumber:
		err := validator.New().Struct(struct {
			value string `validate:"len=16,number"`
		}{
			value: defaultValue,
		})
		if err != nil {
			return constants.Error(http.StatusBadRequest, constants.ErrInvalidInputType, defaultValue, inputType)
		}
	case constants.FormInputTypePostalCode:
		err := validator.New().Struct(struct {
			value string `validate:"len=5,number"`
		}{
			value: defaultValue,
		})
		if err != nil {
			return constants.Error(http.StatusBadRequest, constants.ErrInvalidInputType, defaultValue, inputType)
		}
	case constants.FormInputTypeVehicleNumber:
		_, errs := FormatVehicleNumber(defaultValue)
		if errs != nil {
			return constants.Error(http.StatusBadRequest, constants.ErrInvalidInputType, defaultValue, inputType)
		}
	case constants.FormInputTypeRadio:
		if !slices.Contains(enums, defaultValue) {
			return constants.Error(http.StatusBadRequest, constants.ErrInvalidInputType, defaultValue, inputType)
		}
	case constants.FormInputTypeCheckbox:
		if !slices.Contains(enums, defaultValue) {
			return constants.Error(http.StatusBadRequest, constants.ErrInvalidInputType, defaultValue, inputType)
		}
	case constants.FormInputTypeSelect:
		if !slices.Contains(enums, defaultValue) {
			return constants.Error(http.StatusBadRequest, constants.ErrInvalidInputType, defaultValue, inputType)
		}
	case constants.FormInputTypeAutoComplete:
		if !slices.Contains(enums, defaultValue) {
			return constants.Error(http.StatusBadRequest, constants.ErrInvalidInputType, defaultValue, inputType)
		}
	default:
		return constants.Error(http.StatusBadRequest, constants.ErrFilledDefaultInputType, inputType)
	}

	return nil
}

func ErrApprovedEntry(label string) *constants.ErrorResponse {
	return constants.Error(http.StatusBadRequest, constants.ErrApprovedEntry, label)
}

func IndonesianFormReportType(reportType string) string {
	var result string

	switch reportType {
	case constants.FormReportTypeDaily:
		result = "Harian"
	case constants.FormReportTypeWeekly:
		result = "Mingguna"
	case constants.FormReportTypeMonthly:
		result = "Bulanan"
	case constants.FormReportTypeEvery2Months:
		result = "2 Bulan Sekali"
	case constants.FormReportTypeEvery6Months:
		result = "6 Bulan Sekali"
	case constants.FormReportTypeAnnual:
		result = "Tahunan"
	case constants.FormReportTypeEvent:
		result = "Event"
	}

	return result
}

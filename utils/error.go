package utils

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/yunarsuanto/base-go/constants"
)

func ErrDatabase(err error, table string) *constants.ErrorResponse {
	if err == sql.ErrNoRows {
		return constants.Error(http.StatusNotFound, constants.ErrDataNotFound, table)
	}

	if strings.Contains(err.Error(), constants.DbDuplicateConstraint) {
		errs := strings.Split(err.Error(), "'")
		s := trimQuote(errs[1])
		if constants.CustomDatabaseUniqueKeyErrorResponse[s] != nil {
			return constants.CustomDatabaseUniqueKeyErrorResponse[s]
		}
		return ErrorInternalServer(constants.ErrDuplicate, s)
	}

	if strings.Contains(err.Error(), constants.DbForeignKeyConstraint) {
		errs := strings.Split(err.Error(), "`")
		s := trimQuote(errs[7])
		return ErrorInternalServer(constants.ErrForeignKey, strcase.ToLowerCamel(s))
	}

	if strings.Contains(err.Error(), constants.DbDeleteRestriction) {
		return constants.ErrDeleteRestriction
	}

	return ErrorInternalServer(err.Error())
}

func ErrHttpClient(m string) *constants.ErrorResponse {
	parsedUrl, err := url.Parse(m)
	if err != nil {
		return ErrorInternalServer(err.Error())
	}

	return ErrorInternalServer(constants.ErrHttpClient, parsedUrl.Host)
}

func ErrJwt(err error) *constants.ErrorResponse {
	if err.Error() == constants.JWTInvalidType {
		return constants.ErrKeyIsNotInvalidType
	}
	if strings.Contains(err.Error(), constants.JWTExpired) {
		return constants.ErrTokenReplaced
	}

	return ErrorInternalServer(err.Error())
}

func ErrRedis(m string, key string) *constants.ErrorResponse {
	httpStatusCode := http.StatusInternalServerError

	if m == constants.RedisNilValue {
		return constants.ErrTokenReplaced
	}
	return constants.Error(httpStatusCode, constants.ErrRedis, m, key)
}

func ErrInvalidDateFormat(m string) *constants.ErrorResponse {
	return constants.Error(http.StatusBadRequest, constants.ErrInvalidDateFormat, m)
}

func ErrInvalidOrderKey(m []string) *constants.ErrorResponse {
	return constants.Error(http.StatusBadRequest, constants.ErrInvalidOrderKey, m...)
}

func ErrInvalidOrder(m []string) *constants.ErrorResponse {
	return constants.Error(http.StatusBadRequest, constants.ErrInvalidOrder, m...)
}

func ErrMailTooFrequent(t TimeCountdown) *constants.ErrorResponse {
	var remainingTime []string
	if t.Days > 0 {
		remainingTime = append(remainingTime, fmt.Sprintf("%d %s", t.Days, constants.Days))
	}
	if t.Hours > 0 {
		remainingTime = append(remainingTime, fmt.Sprintf("%d %s", t.Hours, constants.Hours))
	}
	if t.Minutes > 0 {
		remainingTime = append(remainingTime, fmt.Sprintf("%d %s", t.Minutes, constants.Minutes))
	}
	if t.Seconds > 0 {
		remainingTime = append(remainingTime, fmt.Sprintf("%d %s", t.Seconds, constants.Seconds))
	}

	return constants.Error(http.StatusBadRequest, constants.ErrTooFrequentMail, strings.Join(remainingTime, " "))
}

func ErrorInternalServer(message string, args ...string) *constants.ErrorResponse {
	return constants.Error(http.StatusInternalServerError, message, args...)
}

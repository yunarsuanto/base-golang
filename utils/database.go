package utils

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/yunarsuanto/base-go/constants"
)

func checkOrderKeyRegex(order string) bool {
	snakeCaseRegex := regexp.MustCompile(`^(?:[a-z]+\.)?[a-z]+(?:_[a-z]+)*$`)
	match := snakeCaseRegex.MatchString(order)

	return match
}

func QueryOperation(query *string, order [][2]string, limit uint32, page int) *constants.ErrorResponse {
	orderVal := []string{}

	if len(order) != 0 {
		var invalidKeys []string
		var invalidValues []string
		for _, val := range order {
			validOrder := checkOrderKeyRegex(val[0])
			if !validOrder {
				invalidKeys = append(invalidKeys, val[0])
			}
			for i, v := range constants.ValidOrderValue() {
				if v == val[1] {
					break
				}
				if i == len(constants.ValidOrderValue())-1 {
					invalidValues = append(invalidValues, val[0])
				}
			}
		}
		if len(invalidKeys) != 0 {
			return ErrInvalidOrderKey(invalidKeys)
		}

		if len(invalidValues) != 0 {
			return ErrInvalidOrder(invalidValues)
		}

		for _, val := range order {
			orderVal = append(orderVal, fmt.Sprintf("%s %s", val[0], val[1]))
		}

		orderQueryResult := fmt.Sprintf("%s ORDER BY %s", *query, strings.Join(orderVal, ","))
		*query = orderQueryResult
	}

	offsetVal := (uint32(page) - 1) * limit
	limitQueryResult := fmt.Sprintf("%s LIMIT %d OFFSET %d", *query, limit, offsetVal)
	*query = limitQueryResult

	return nil
}

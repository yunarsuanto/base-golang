package utils

import (
	"database/sql"
	"math"
)

type Number interface {
	int | int32 | int64 | float32 | float64
}

func AbsoluteNumber[T Number](num T) T {
	if num < 0 {
		return -num
	}
	return num
}

func NewNullInt64[T Number](s T) sql.NullInt64 {
	if s == 0 {
		return sql.NullInt64{}
	}
	return sql.NullInt64{
		Int64: int64(s),
		Valid: true,
	}
}

func NewNullFloat64[T Number](s T) sql.NullFloat64 {
	if s == 0 {
		return sql.NullFloat64{}
	}
	return sql.NullFloat64{
		Float64: float64(s),
		Valid:   true,
	}
}

func ToFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	rounded := int(num*output + math.Copysign(0.5, num*output))

	return float64(rounded) / output
}

func FirstDigit(n int) int {
	if n < 0 {
		n = -n
	}
	for n >= 10 {
		n /= 10
	}
	return n
}

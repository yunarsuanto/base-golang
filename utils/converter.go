package utils

func NullScan[T any](data *T) T {
	var result T
	if data == nil {
		return result
	}

	return *data
}

func ToNullScan[T comparable](data T, zero T) *T {
	if data == zero {
		return nil
	}
	return &data
}

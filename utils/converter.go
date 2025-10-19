package utils

func NullScan[T any](data *T) T {
	var result T
	if data == nil {
		return result
	}

	return *data
}

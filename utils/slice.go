package utils

func RemoveDuplicate[T comparable](slice []T) []T {
	keys := make(map[T]bool)
	list := []T{}
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func InArrayExist[T comparable](val T, array []T) (exists bool) {
	exists = false
	for i := 0; i < len(array); i++ {
		if val == array[i] {
			exists = true
			break
		}
	}
	return
}

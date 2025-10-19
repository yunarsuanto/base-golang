package utils

import (
	"reflect"
)

func MaskBody(data any) any {
	val := reflect.ValueOf(data)

	if val.Kind() == reflect.Ptr && val.Elem().Kind() == reflect.Interface {
		val = val.Elem()
	}

	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		return data
	}

	original := val.Elem()
	copyVal := reflect.New(original.Type()).Elem()
	copyVal.Set(original)

	for i := 0; i < original.NumField(); i++ {
		field := original.Type().Field(i)
		tag := field.Tag.Get("log")

		if tag == "masked" {
			copyField := copyVal.Field(i)
			if copyField.CanSet() {
				copyField.Set(reflect.Zero(copyField.Type()))
			}
		}
	}

	return copyVal.Interface()
}

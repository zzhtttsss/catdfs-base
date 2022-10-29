package util

import (
	"github.com/google/uuid"
)

func GenerateUUIDString() string {
	return uuid.NewString()
}

func Interfaces2TypeArr[T interface{}](arr []interface{}) (res []T) {
	for _, a := range arr {
		res = append(res, a.(T))
	}
	return
}

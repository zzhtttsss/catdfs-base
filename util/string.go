package util

import "github.com/google/uuid"

func GenerateUUIDString() string {
	return uuid.NewString()
}

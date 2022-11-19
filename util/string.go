package util

import (
	"github.com/google/uuid"
	"hash/crc32"
	"strconv"
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

func CRC32String(bytes []byte) string {
	return strconv.Itoa(int(crc32.ChecksumIEEE(bytes)))
}

func CRC32(bytes []byte) uint32 {
	return crc32.ChecksumIEEE(bytes)
}

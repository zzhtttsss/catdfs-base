package util

import (
	"github.com/google/uuid"
	"hash/crc32"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
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

// String2Bytes is faster method to convert string to byte array. It should be
// noted that byte array cannot be modified.
func String2Bytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// CombineString uses builder to combine multiple strings into one string.
func CombineString(strs ...string) string {
	var builder strings.Builder
	for _, s := range strs {
		builder.WriteString(s)
	}
	return builder.String()
}

// CombineStringWithLength uses []byte to combine multiple strings into one string.
// It is faster than CombineString, but it needs the total length of all strings.
func CombineStringWithLength(len int, strs ...string) string {
	buf := make([]byte, 0, len)
	for _, s := range strs {
		buf = append(buf, s...)
	}
	return string(buf)
}

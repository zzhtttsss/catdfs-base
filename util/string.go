package util

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"tinydfs-base/common"
)

func GenerateUUIDString() string {
	return uuid.NewString()
}

func Marshal(a any) string {
	bytes, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("%v marshal failed.Error code %v\n", a, common.JsonMarshalFailed)
	}
	return string(bytes)
}

func Unmarshal(s string, a *any) {
	err := json.Unmarshal([]byte(s), a)
	if err != nil {
		fmt.Printf("%s unmarshal failed.Error code %v\n", s, common.JsonUnmarshalFailed)
	}
}

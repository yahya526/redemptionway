package util

import "encoding/json"

func MustMarshal(obj interface{}) string {
	bytes, _ := json.Marshal(obj)
	return string(bytes)
}

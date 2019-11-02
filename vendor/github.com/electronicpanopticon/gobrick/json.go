package gobrick

import (
	"encoding/json"
	"fmt"
	"strings"
)

func ToJsonBytes(v interface{}) []byte {
	b, err := json.MarshalIndent(v, "", "  ")
	if string(b) == "null" {
		return nil
	}
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return b
}

func ToJson(v interface{}) string {
	b := strings.TrimSpace(string(ToJsonBytes(v)))
	if b == "" || b == "\"\"" {
		return "{}"
	}
	return string(b)
}
package interfaces

import (
	"encoding/json"
)

var maskKeys = []string{
	"authorize_key", "password",
}

func Log(s string) (string, error) {
	var r map[string]interface{}
	if err := json.Unmarshal([]byte(s), &r); err != nil {
		return "", err
	}
	for _, key := range maskKeys {
		if _, ok := r[key]; ok {
			r[key] = "***"
		}
	}
	output, err := json.Marshal(r)
	if err != nil {
		return "", err
	}
	return string(output), nil
}

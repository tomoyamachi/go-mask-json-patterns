package interfaces

import (
	"encoding/json"
	"errors"
)

var maskKeys = []string{
	"authorize_key", "password",
}

func Log(s string) (string, error) {
	var r interface{}
	if err := json.Unmarshal([]byte(s), &r); err != nil {
		return "", err
	}
	mask, ok := r.(map[string]interface{})
	if !ok {
		return "", errors.New("invalid json struct")
	}
	for _, key := range maskKeys {
		if _, ok := mask[key]; ok {
			mask[key] = "***"
		}
	}
	output, err := json.Marshal(mask)
	if err != nil {
		return "", err
	}
	return string(output), nil
}

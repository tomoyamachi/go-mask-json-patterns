package interfaces

import (
	"encoding/json"
	"strings"
)

func Log(s string, maskKeys []string) (string, error) {
	var r map[string]interface{}
	if err := json.Unmarshal([]byte(s), &r); err != nil {
		return "", err
	}
	for _, key := range maskKeys {
		mask(r, strings.Split(key, "/"))
	}
	output, err := json.Marshal(r)
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func mask(target map[string]interface{}, path []string) bool {
	if val, ok := target[path[0]]; ok {
		if len(path) > 1 {
			if v, ok := val.(map[string]interface{}); ok {
				return mask(v, path[1:])
			}
		}
		target[path[0]] = "***"
		return true
	}
	return false
}

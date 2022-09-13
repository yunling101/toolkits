package core

import "encoding/json"

// ObjectsToString
func ObjectsToString(obj interface{}) (string, error) {
	if b, err := json.Marshal(obj); err != nil {
		return "", err
	} else {
		return string(b), nil
	}
}

package utils

import "encoding/json"

// convert json in byte array to target type
func DecodeJSON(source []byte, target interface{}) error {
	err := json.Unmarshal(source, &target)
	return err
}

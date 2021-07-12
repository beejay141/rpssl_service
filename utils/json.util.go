package utils

import "encoding/json"

// convert json in byte array to target type
func DecodeJSON(source []byte, target interface{}) error {
	err := json.Unmarshal(source, &target)
	return err
}

func EncodeJSON(source interface{}) ([]byte, error) {
	target, err := json.Marshal(&source)
	if err != nil {
		return nil, err
	}
	return target, nil
}

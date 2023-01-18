package helpers

import (
	"bytes"
	"encoding/json"
)

func TrimJson(jsonBytes []byte) ([]byte, error) {
	buffer := new(bytes.Buffer)
	if err := json.Compact(buffer, jsonBytes); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

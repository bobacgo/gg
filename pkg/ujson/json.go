package ujson

import (
	"encoding/json"
)

func MarshalIndent(bytes []byte) []byte {
	var result map[string]any
	if err := json.Unmarshal(bytes, &result); err != nil {
		println("[ujson] json.Unmarshal err", err.Error())
		return bytes
	}
	respBody, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		println("[ujson] json.MarshalIndent err", err)
		return bytes
	}
	return respBody
}

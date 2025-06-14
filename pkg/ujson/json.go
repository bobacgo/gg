package ujson

import (
	"encoding/json"
	"fmt"
)

func MarshalIndent(bytes []byte) string {
	var result map[string]any
	if err := json.Unmarshal(bytes, &result); err != nil {
		return string(bytes)
	}
	respBody, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		fmt.Println("[ujson] json.MarshalIndent err", err)
		return string(bytes)
	}
	return string(respBody)
}

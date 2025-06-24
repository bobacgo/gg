package cmd

import (
	"encoding/json"
	"fmt"
	"os"
)

var cfg = struct {
	Http HttpConfig `json:"http"`
}{}

var cfgPath = ".gg.json"

func initConfig() {
	bytes, err := os.ReadFile(cfgPath)
	if err != nil {
		if !os.IsNotExist(err) {
			fmt.Printf("[read] os.ReadFile config path [%s] err: %s", cfgPath, err)
			return
		}
	}
	if len(bytes) == 0 {
		return
	}
	if err := json.Unmarshal(bytes, &cfg); err != nil {
		fmt.Println("parse cfg err: ", err)
	}
}

func saveConfig() {
	bytes, err := json.Marshal(cfg)
	if err != nil {
		fmt.Println("marshal cfg err: ", err)
		return
	}
	if err := os.WriteFile(cfgPath, bytes, 0666); err != nil {
		fmt.Println("write cfg err: ", err)
	}
}

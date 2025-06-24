package ufile

import (
	"fmt"
	"os"
)

func Overwrite(filepath string, process func([]byte) []byte) error {
	var data []byte
	if _, err := os.Stat(filepath); err == nil { // 如果是文件，就把数据读取出来
		data, err = os.ReadFile(filepath)
		if err != nil {
			return fmt.Errorf("os.ReadFile: %w", err)
		}
		data = process(data)
	} else if os.IsNotExist(err) {
		data = process([]byte(filepath))
	} else if err != nil {
		return fmt.Errorf("os.Stat: %w", err)
	}
	println(string(data))
	if len(data) == 0 {
		return fmt.Errorf("no data to write to file: %s", filepath)
	}
	if err := os.WriteFile(filepath, data, 0666); err != nil {
		return fmt.Errorf("os.WriteFile: %w", err)
	}
	return nil
}

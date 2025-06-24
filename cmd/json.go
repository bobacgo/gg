/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"strconv"

	"github.com/bobacgo/gg/pkg/ufile"
	"github.com/bobacgo/gg/pkg/ujson"
	"github.com/spf13/cobra"
)

var jsonFlag = struct {
	format    string
	marshal   string
	unmarshal string
}{}

// jsonCmd represents the json command
var jsonCmd = &cobra.Command{
	Use:   "json",
	Short: "格式化、序列化 JSON 数据",
	Long: `支持对 JSON 数据进行格式化（美化）、序列化操作。

用法示例:

1. 格式化 JSON 文件内容并覆盖原文件:
   gg json -f data.json

2. 格式化 JSON 字符串:
   gg json -f '{"name":"Alice","age":30}'

3. 将文件内容序列化为 JSON 字符串:
   gg json -m data.txt

参数说明:
  -f, --format   对 JSON 进行格式化（默认开启）
  -m, --marshal  将文件内容序列化为 JSON 字符串
`,
	Run: func(cmd *cobra.Command, args []string) {
		if jsonFlag.format != "" {
			if err := ufile.Overwrite(jsonFlag.format, func(data []byte) []byte {
				return ujson.MarshalIndent(data)
			}); err != nil {
				println("[format] ufile.Overwrite err: ", err)
			}
		}
		if jsonFlag.marshal != "" {
			if err := ufile.Overwrite(jsonFlag.marshal, func(b []byte) []byte {
				data, err := json.Marshal(string(b))
				if err != nil {
					println("[marshal] json.Marshal err: ", err)
				}
				return data
			}); err != nil {
				println("[encode] ufile.Overwrite err: ", err)
			}
		}
		if jsonFlag.unmarshal != "" {
			if err := ufile.Overwrite(jsonFlag.unmarshal, func(b []byte) []byte {
				res, err := strconv.Unquote(string(b))
				if err != nil {
					println("[unmarshal] strconv.Unquote err: ", err.Error())
					return nil
				}
				return ujson.MarshalIndent([]byte(res))
			}); err != nil {
				println("[decode] ufile.Overwrite err: ", err.Error())
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(jsonCmd)
	jsonCmd.Flags().StringVarP(&jsonFlag.format, "format", "f", "", "json format")
	jsonCmd.Flags().StringVarP(&jsonFlag.marshal, "encode", "e", "", "json marshal")
	jsonCmd.Flags().StringVarP(&jsonFlag.unmarshal, "decode", "d", "", "json unmarshal")
}

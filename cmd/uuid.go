/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var uuidFmt string

// uuidCmd represents the uuid command
var uuidCmd = &cobra.Command{
	Use:   "uuid",
	Short: "生成 UUID",
	Long: `生成一个或多个 UUID，可自定义分隔符。

用法示例:
  gg uuid           # 生成一个标准 UUID
  gg uuid 5         # 生成 5 个 UUID
  gg uuid -f ""     # 生成无分隔符的 UUID
  gg uuid 3 -f ":"  # 生成 3 个以冒号分隔的 UUID
`,
	Run: func(cmd *cobra.Command, args []string) {
		n := 1
		if len(args) > 0 {
			arg1, _ := strconv.Atoi(args[0])
			if arg1 > 0 {
				n = arg1
			}
		}
		for range n {
			if uuidFmt == "" {
				fmt.Println(strings.ReplaceAll(uuid.NewString(), "-", uuidFmt))
			} else {
				fmt.Println(uuid.NewString())
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(uuidCmd)
	uuidCmd.Flags().StringVarP(&uuidFmt, "fmt", "f", "", "指定uuid分隔符")
}

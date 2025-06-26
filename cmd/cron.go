/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"
)

var cronYear int8

// cronCmd represents the cron command
var cronCmd = &cobra.Command{
	Use:   "cron",
	Short: "解析并展示 cron 表达式的未来执行时间",
	Long: `该命令用于解析 cron 表达式，并输出未来的执行时间点。

支持标准 cron 表达式格式，例如：

	* * * * * <command>
	- - - - -
	| | | | |
	| | | | +----- 星期 (0 - 7) (0 和 7 都代表周日)
	| | | +------- 月 (1 - 12)
	| | +--------- 日 (1 - 31)
	| +----------- 小时 (0 - 23)
	+------------- 分钟 (0 - 59)

	* * * * * * <command>
	- - - - - -
	| | | | | |
	| | | | | +---- 年份 (可选，1970 - 2099)
	| | | | +------ 星期 (0 - 7)
	| | | +-------- 月 (1 - 12)
	| | +---------- 日 (1 - 31)
	| +------------ 小时 (0 - 23)
	+-------------- 分钟 (0 - 59)

示例用法:
	gg cron "0 0 * * *"
	
	近5次执行时间（未来）:
	2025-06-26 00:00:00
	2025-06-27 00:00:00
	2025-06-28 00:00:00
	2025-06-29 00:00:00
	2025-06-30 00:00:00
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := args[0]
		// 解析 cron 表达式
		// 输出近5次执行时间

		if cmd.Flags().Changed("6") {
			// TODO 实现
		}

		expr, err := cron.ParseStandard(arg)
		if err != nil {
			fmt.Println("解析失败:", err)
			return
		}

		now := time.Now()
		next := expr.Next(now)
		fmt.Println("近5次执行时间（未来）:")
		for range 5 {
			fmt.Println(next.Format(time.DateTime))
			next = expr.Next(next)
		}
	},
}

func init() {
	rootCmd.AddCommand(cronCmd)
	cronCmd.Flags().Int8VarP(&cronYear, "6", "6", 6, "supper year")
}

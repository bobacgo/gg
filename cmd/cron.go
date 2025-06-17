/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"
	"time"
)

var cronYear int8

// cronCmd represents the cron command
var cronCmd = &cobra.Command{
	Use:   "cron",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

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
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}
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
		for i := 0; i < 5; i++ {
			fmt.Println(next.Format(time.DateTime))
			next = expr.Next(next)
		}
	},
}

func init() {
	rootCmd.AddCommand(cronCmd)
	cronCmd.Flags().Int8VarP(&cronYear, "6", "6", 6, "supper year")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cronCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cronCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

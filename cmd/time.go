/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

// timeCmd represents the time command
var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "时间戳与日期转换工具",
	Long: `time 命令用于在时间戳和日期字符串之间进行转换。

用法示例:

  time                # 输出当前时间、秒级时间戳、毫秒级时间戳
  time 2024-06-01     # 输出该日期的起止时间戳
  time "2024-06-01 12:00:00" # 输出该时间的时间戳
  time 1717219200     # 输出时间戳对应的时间

支持秒、毫秒、微秒级时间戳与日期字符串的相互转换。`,
	Aliases: []string{"t"},
	Args:    cobra.MaximumNArgs(1), // 允许最多一个参数
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(time.Now())
			fmt.Println(time.Now().Unix())
			fmt.Println(time.Now().UnixMilli())
			return
		}
		t := args[0]

		ts, err := strconv.ParseInt(t, 10, 64)
		if err != nil {
			if len(t) == len(time.DateOnly) {
				tm, err := time.Parse(time.DateOnly, t)
				if err != nil {
					fmt.Println("Invalid date format:", err)
					return
				}
				fmt.Println(tm.Unix())                      // 输出当天的开始时间戳
				fmt.Println(tm.AddDate(0, 0, 1).Unix() - 1) // 输出当天的结束时间戳
			} else if len(t) == len(time.DateTime) {
				tm, err := time.Parse(time.DateTime, t)
				if err != nil {
					fmt.Println("Invalid date format:", err)
					return
				}
				fmt.Println(tm.Unix())
			}
		} else if len(t) == 10 {
			fmt.Println(time.Unix(ts, 0))
		} else if len(t) == 13 {
			fmt.Println(time.UnixMilli(cast.ToInt64(t)))
		} else if len(t) == 16 {
			fmt.Println(time.UnixMicro(cast.ToInt64(t)))
		}
	},
}

func init() {
	rootCmd.AddCommand(timeCmd)
}

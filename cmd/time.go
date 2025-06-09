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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
				fmt.Println(tm.Unix())
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// timeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// timeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

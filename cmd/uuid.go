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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
	uuidCmd.Flags().StringVarP(&uuidFmt, "fmt", "F", "", "指定uuid分隔符")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uuidCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uuidCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/bobacgo/gg/pkg/ujson"
	"github.com/spf13/cobra"
	"os"
)

var (
	jsonFormat    bool
	jsonMarshal   bool
	jsonUnmarshal bool
)

// jsonCmd represents the json command
var jsonCmd = &cobra.Command{
	Use:   "json",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			return
		}
		arg1 := args[0]
		if cmd.Flags().Changed("format") {
			_, err := os.Stat(arg1)
			if err != nil {
				if os.IsNotExist(err) {
					indent := ujson.MarshalIndent([]byte(arg1))
					fmt.Println(indent)
				} else {
					fmt.Println("os.Stat err: ", err)
				}
			} else { // 指定的是有效文件路径
				bytes, err := os.ReadFile(arg1)
				if err != nil {
					fmt.Println("os.ReadFile err: ", err)
					return
				}
				indent := ujson.MarshalIndent(bytes)
				fmt.Println(indent)
				if err := os.WriteFile(arg1, []byte(indent), 0666); err != nil {
					fmt.Println("os.WriteFile err: ", err)
				}
			}
		} else if cmd.Flags().Changed("marshal") {
			bytes, err := os.ReadFile(arg1)
			res, err := json.Marshal(string(bytes))
			if err != nil {
				fmt.Println("json.Marshal", err)
			} else {
				fmt.Println(string(res))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(jsonCmd)
	jsonCmd.Flags().BoolVarP(&jsonFormat, "format", "f", true, "json format")
	jsonCmd.Flags().BoolVarP(&jsonMarshal, "marshal", "m", true, "json marshal")
}
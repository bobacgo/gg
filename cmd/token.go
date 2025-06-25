/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/cobra"
)

// tokenCmd represents the token command
var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Parse and display JWT token claims",
	Long: `Parse a JWT token and pretty-print its claims as JSON.

Example usage:
  gg token <your-jwt-token>
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Help()
			return
		}
		tokenStr := args[0]
		parse, _ := jwt.Parse(tokenStr, nil)
		bytes, err := json.MarshalIndent(parse.Claims, "", "  ")
		if err != nil {
			fmt.Println("json.MarshalIndent err: ", err)
			return
		}
		fmt.Println(string(bytes))
	},
}

func init() {
	rootCmd.AddCommand(tokenCmd)
}

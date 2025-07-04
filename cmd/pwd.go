/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var pwdDeletFlag string

// pwdCmd represents the pwd command
var pwdCmd = &cobra.Command{
	Use:   "pwd",
	Short: "Manage your password entries",
	Long: `The pwd command allows you to manage password entries.

You can list all passwords, add a new password, or delete an existing one.

Examples:
  gg pwd                # List all passwords
  gg pwd key value      # Add or update a password with key and value
  gg pwd -d key         # Delete the password with the specified key
`,
	Aliases: []string{"kv"},
	Run: func(cmd *cobra.Command, args []string) {
		if pwdDeletFlag != "" {
			delete(cfg.PwdMgr, pwdDeletFlag)
			saveConfig()
			return
		}
		if len(args) == 0 {
			for k, v := range cfg.PwdMgr {
				fmt.Println(k, v)
			}
			return
		}

		if len(args) < 2 {
			cmd.Help()
			return
		}
		key := args[0]
		value := strings.Join(args[1:], " ")
		cfg.PwdMgr[key] = value
		saveConfig()
	},
}

func init() {
	rootCmd.AddCommand(pwdCmd)
	pwdCmd.Flags().StringVarP(&pwdDeletFlag, "delete", "d", "", "delete a password by key")
	if len(cfg.PwdMgr) == 0 {
		cfg.PwdMgr = make(map[string]string)
	}
}

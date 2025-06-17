/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"os"
)

// md5Cmd represents the md5 command
var md5Cmd = &cobra.Command{
	Use:   "md5",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Help()
			return
		}
		arg := args[0]
		_, err := os.Stat(arg)
		if err != nil {
			if !os.IsNotExist(err) {
				fmt.Println("os.Stat err: ", err)
				return
			}
			fmt.Println(MD5(arg))
		} else {
			fmt.Println("file:", fileMD5(arg))
		}
	},
}

func init() {
	rootCmd.AddCommand(md5Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// md5Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// md5Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// --- MD5签名 ---

func MD5(data string) string {
	if data == "" {
		return ""
	}
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func fileMD5(filename string) string {
	f, err := os.Open(filename)
	if err != nil {
		return ""
	}
	defer f.Close()
	h := md5.New()
	buf := make([]byte, 1024*1024*10) // 10MB 缓冲区
	_, err = io.CopyBuffer(h, f, buf)
	sum := h.Sum(nil)
	return hex.EncodeToString(sum)
}

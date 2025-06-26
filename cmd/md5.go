/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

// md5Cmd represents the md5 command
var md5Cmd = &cobra.Command{
	Use:   "md5",
	Short: "Calculate the MD5 hash of a string or file",
	Long: `Calculate the MD5 hash for a given string or file.

Usage:
  md5 <string|filepath>

If the argument is a valid file path, the MD5 of the file's contents will be calculated.
Otherwise, the MD5 of the input string will be calculated.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := args[0]
		if _, err := os.Stat(arg); err != nil {
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
	if _, err = io.CopyBuffer(h, f, buf); err != nil {
		fmt.Println("io.CopyBuffer err: ", err)
		return ""
	}
	sum := h.Sum(nil)
	return hex.EncodeToString(sum)
}

/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/base64"
	"fmt"

	"github.com/spf13/cobra"
)

var base64Flags = struct {
	Encode string
	Decode string
}{}

// base64Cmd represents the base64 command
var base64Cmd = &cobra.Command{
	Use:   "base64",
	Short: "Base64 encode or decode input",
	Long: `Encode or decode data using base64.

Examples:
  gg base64 encode "hello"
  gg base64 decode "aGVsbG8="`,
	Aliases: []string{"b64"},
	Run: func(cmd *cobra.Command, args []string) {
		if base64Flags.Encode != "" {
			encoded := encodeBase64(base64Flags.Encode)
			fmt.Println(encoded)
		}
		if base64Flags.Decode != "" {
			decoded, err := decodeBase64(base64Flags.Decode)
			if err != nil {
				fmt.Println("decode error:", err)
				return
			}
			fmt.Println(decoded)
		}
	},
}

func encodeBase64(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func decodeBase64(s string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func init() {
	rootCmd.AddCommand(base64Cmd)
	base64Cmd.Flags().StringVarP(&base64Flags.Encode, "encode", "e", "", "String to encode in base64")
	base64Cmd.Flags().StringVarP(&base64Flags.Decode, "decode", "d", "", "Base64 string to decode")
}

/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/bobacgo/gg/pkg/ujson"
	"github.com/spf13/cobra"
	"resty.dev/v3"
)

var httpClt *resty.Client

var httpFlag = struct {
	Env          string
	BaseURL      string
	Headers      []string
	RequestQuery string
	Debug        bool
}{}

type HttpConfig struct {
	Env     string            `yaml:"env"`
	BaseURL map[string]string `json:"baseURL"` // dev, test, prod
	Headers map[string]string `json:"headers"`
}

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Send HTTP requests with flexible configuration",
	Long: `Send HTTP requests with configurable base URLs, headers, and request bodies.
Examples:
  gg http http://localhost:8080/api/v1/cfg --debug -e=test
  gg http /api/v1/cfg -b dev=http://localhost:8080
  gg http post http://localhost:8080/api/v1/user -r "{\"name\": \"bobacgo\"}"
  gg http post http://www.imooc.com/search/hotwords -H token=234 -H app=1
`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}

		if cmd.Flags().Changed("debug") {
			httpClt.SetDebug(httpFlag.Debug)
		}

		isChange := false

		if httpFlag.Env != "" {
			isChange = true
			cfg.Http.Env = httpFlag.Env
		}

		if httpFlag.BaseURL != "" {
			isChange = true

			parts := strings.Split(httpFlag.BaseURL, "=")
			if cfg.Http.BaseURL == nil {
				cfg.Http.BaseURL = make(map[string]string, len(httpFlag.BaseURL))
			}

			if len(parts) == 1 { // 没有等号，认为是默认值 "dev=http://localhost:8080"
				cfg.Http.Env = "dev"
				cfg.Http.BaseURL["dev"] = parts[0]
			} else if parts[1] == "" { // 删除某个环境配置 "dev="
				if cfg.Http.Env == parts[0] { // 重置为默认环境
					cfg.Http.Env = "dev"
				}
				delete(cfg.Http.BaseURL, parts[0])
			} else {
				cfg.Http.Env = parts[0]
				cfg.Http.BaseURL[parts[0]] = parts[1]
			}
		}
		for _, h := range httpFlag.Headers {
			if !strings.Contains(h, "=") {
				fmt.Printf("Invalid header format: %s\n", h)
				continue
			}
			parts := strings.Split(h, "=")
			if len(parts) != 2 {
				fmt.Printf("Invalid header format: %s\n", h)
				continue
			}
			isChange = true
			if cfg.Http.Headers == nil {
				cfg.Http.Headers = make(map[string]string)
			}
			if parts[1] == "" { // 删除某个headers "app="
				delete(cfg.Http.Headers, parts[0])
			} else {
				cfg.Http.Headers[parts[0]] = parts[1]
			}
		}

		if len(cfg.Http.BaseURL) == 1 { // 只有一个环境，则使用当前环境配置
			for k, v := range cfg.Http.BaseURL {
				cfg.Http.Env = k
				httpClt.SetBaseURL(v)
			}
		} else if len(cfg.Http.BaseURL) > 0 {
			baseURL := cfg.Http.BaseURL[cfg.Http.Env]
			httpClt.SetBaseURL(baseURL)
		}
		if len(cfg.Http.Headers) > 0 {
			httpClt.SetHeaders(cfg.Http.Headers)
		}

		fmt.Println("http config:", cfg.Http)

		if isChange {
			saveConfig()
		}

		if len(args) == 1 {
			resp, err := httpClt.R().Get(args[0])
			if err != nil {
				fmt.Printf("Error making GET request: %v\n", err)
				return
			}

			fmt.Println(string(ujson.MarshalIndent(resp.Bytes())))
		} else {
			method := strings.ToUpper(args[0])
			url := args[1]

			r := httpClt.R()
			if httpFlag.RequestQuery != "" {
				_, err := os.Stat(httpFlag.RequestQuery)
				if err != nil {
					if !os.IsNotExist(err) {
						fmt.Println("os.Stat err: ", err)
					}
					r = r.SetBody(httpFlag.RequestQuery)
				} else { // 指定的是有效文件路径
					bytes, err := os.ReadFile(httpFlag.RequestQuery)
					if err != nil {
						fmt.Println("os.ReadFile err: ", err)
						return
					}
					r = r.SetBody(string(bytes))
				}
			}
			resp, err := r.Execute(method, url)
			if err != nil {
				fmt.Printf("Error making %s request to %s: %v\n", method, url, err)
				return
			}
			fmt.Println(string(ujson.MarshalIndent(resp.Bytes())))
		}
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
	httpCmd.Flags().StringVarP(&httpFlag.Env, "env", "e", "", "set base URL env")
	httpCmd.Flags().StringVarP(&httpFlag.BaseURL, "baseURL", "b", "", "set base URL default dev dev=http://localhost:8080")
	httpCmd.Flags().StringSliceVarP(&httpFlag.Headers, "headers", "H", nil, "set HTTP header app=1")
	httpCmd.Flags().StringVarP(&httpFlag.RequestQuery, "request", "r", "", "request body")
	httpCmd.Flags().BoolVarP(&httpFlag.Debug, "debug", "d", true, "debug mode")
	httpClt = resty.New()
}

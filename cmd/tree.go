/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	treelevel int
)

// treeCmd represents the tree command
var treeCmd = &cobra.Command{
	Use:   "tree [directory]",
	Short: "Display directory structure as a tree",
	Long: `tree displays the directory structure of a given path in a tree-like format.

You can specify the maximum depth to display using the --level or -L flag.
Hidden files and directories (those starting with a dot) are ignored by default.

Examples:
  gg tree
  gg tree /path/to/dir -L 2
`,
	Run: func(cmd *cobra.Command, args []string) {
		dir := "./"
		if len(args) > 0 {
			dir = args[0]
		}
		printTree(dir, "", true, 0, treelevel)
	},
}

func init() {
	rootCmd.AddCommand(treeCmd)
	treeCmd.Flags().IntVarP(&treelevel, "level", "L", -1, "Max display depth of the directory tree")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// treeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// treeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Add currentDepth and maxDepth parameters
func printTree(path, prefix string, isRoot bool, currentDepth, maxDepth int) {
	info, err := os.Stat(path)
	if err != nil {
		fmt.Printf("%s[error: %v]\n", prefix, err)
		return
	}
	if !isRoot {
		fmt.Print(prefix)
		// 判断是否为最后一个子项，prefix 末尾由调用方控制
		if prefix != "" && len(prefix) >= 4 && prefix[len(prefix)-4:] == "    " {
			fmt.Print("└── ")
		} else {
			fmt.Print("├── ")
		}
		fmt.Println(info.Name())
	} else {
		fmt.Println(info.Name())
	}

	if !info.IsDir() {
		return
	}

	// Stop if we've reached the max depth (unless unlimited)
	if maxDepth >= 0 && currentDepth >= maxDepth {
		return
	}

	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Printf("%s[error: %v]\n", prefix, err)
		return
	}

	// 过滤掉以.开头的文件和目录
	var visibleFiles []os.DirEntry
	for _, file := range files {
		if file.Name() == "" || file.Name()[0] == '.' {
			continue
		}
		visibleFiles = append(visibleFiles, file)
	}

	for i, file := range visibleFiles {
		isLastChild := i == len(visibleFiles)-1
		newPrefix := prefix
		if !isRoot {
			if isLastChild {
				newPrefix += "    "
			} else {
				newPrefix += "│   "
			}
		}
		// 传递 isLastChild 信息，决定前缀符号
		printTree(filepath.Join(path, file.Name()), newPrefix, false, currentDepth+1, maxDepth)
	}
}

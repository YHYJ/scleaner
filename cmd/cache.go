/*
File: version.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-02-21 13:51:10

Description: 程序子命令'cache'时执行
*/

package cmd

import (
	"github.com/YHYJ/scleaner/function"
	"github.com/spf13/cobra"
)

// cacheCmd represents the cache command
var cacheCmd = &cobra.Command{
	Use:   "cache",
	Short: "清除软件包缓存",
	Long:  `清除软件包的缓存，包括pip/npm/yarn`,
	Run: func(cmd *cobra.Command, args []string) {
		function.CacheCleaner()
	},
}

func init() {
	cacheCmd.Flags().BoolP("help", "h", false, "Help for cache")
	rootCmd.AddCommand(cacheCmd)
}

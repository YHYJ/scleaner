/*
File: version.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-02-21 13:51:10

Description: 执行子命令 'cache'
*/

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yhyj/scleaner/cli"
)

// cacheCmd represents the cache command
var cacheCmd = &cobra.Command{
	Use:   "cache",
	Short: "Clear Package Cache",
	Long:  `Clear the cache of packages, including pip/npm/yarn.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.CacheCleaner()
	},
}

func init() {
	cacheCmd.Flags().BoolP("help", "h", false, "help for cache command")
	rootCmd.AddCommand(cacheCmd)
}

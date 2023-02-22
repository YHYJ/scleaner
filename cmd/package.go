/*
File: view.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-02-17 16:08:21

Description: 程序子命令'package'时执行
*/

package cmd

import (
	"github.com/YHYJ/scleaner/function"
	"github.com/spf13/cobra"
)

// packageCmd represents the package command
var packageCmd = &cobra.Command{
	Use:   "package",
	Short: "清除孤立依赖包",
	Long:  `清除作为依赖关系安装，但已不再被依赖的孤立包`,
	Run: func(cmd *cobra.Command, args []string) {
		function.PackageCleaner()
	},
}

func init() {
	packageCmd.Flags().BoolP("help", "h", false, "Help for package")
	rootCmd.AddCommand(packageCmd)
}

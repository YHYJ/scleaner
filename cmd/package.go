/*
File: view.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-02-17 16:08:21

Description: 程序子命令'package'时执行
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// packageCmd represents the package command
var packageCmd = &cobra.Command{
	Use:   "package",
	Short: "清除系统孤立包",
	Long:  `清除系统中没有依赖关系的孤立包`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("package called")
	},
}

func init() {
	packageCmd.Flags().BoolP("help", "h", false, "Help for package")
	rootCmd.AddCommand(packageCmd)
}

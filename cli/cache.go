/*
File: cache.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-02-22 13:35:18

Description: 子命令`cache`功能函数
*/

package cli

import (
	"fmt"

	"github.com/yhyj/scleaner/general"
)

func CacheCleaner() {
	// 清除pip缓存
	fmt.Printf("%v %v\n", "-->", "Cleaning pip cache")
	pipArgs := []string{"cache", "purge"}
	if err := general.RunCommand("pip", pipArgs); err != nil {
		fmt.Printf("\x1b[31m%s\x1b[0m\n", err)
	}
	fmt.Println()

	// 清除npm缓存
	fmt.Printf("%v %v\n", "-->", "Cleaning npm cache")
	npmArgs := []string{"cache", "clean", "--force"}
	if err := general.RunCommand("npm", npmArgs); err != nil {
		fmt.Printf("\x1b[31m%s\x1b[0m\n", err)
	}
	fmt.Println()

	// 清除yarn缓存
	fmt.Printf("%v %v\n", "-->", "Cleaning yarn cache")
	yarnArgs := []string{"cache", "clean"}
	if err := general.RunCommand("yarn", yarnArgs); err != nil {
		fmt.Printf("\x1b[31m%s\x1b[0m\n", err)
	}
	fmt.Println()
}

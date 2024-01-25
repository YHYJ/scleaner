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

// CacheCleaner 清理缓存
func CacheCleaner() {
	// 清除 pip 缓存
	fmt.Printf("%v %v\n", "-->", "Cleaning pip cache")
	pipArgs := []string{"cache", "purge"}
	if err := general.RunCommand("pip", pipArgs); err != nil {
		fmt.Printf(general.ErrorBaseFormat, err)
	}
	fmt.Println()

	// 验证 npm 缓存文件夹
	fmt.Printf("%v %v\n", "-->", "Verify the cache folder")
	npmArgs := []string{"cache", "clean", "verify"}
	if err := general.RunCommand("npm", npmArgs); err != nil {
		fmt.Printf(general.ErrorBaseFormat, err)
	}
	fmt.Println()

	// 清除 yarn 缓存
	fmt.Printf("%v %v\n", "-->", "Cleaning yarn cache")
	yarnArgs := []string{"cache", "clean"}
	if err := general.RunCommand("yarn", yarnArgs); err != nil {
		fmt.Printf(general.ErrorBaseFormat, err)
	}
	fmt.Println()
}

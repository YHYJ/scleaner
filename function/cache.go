/*
File: cache.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-02-22 13:35:18

Description: 子命令`cache`功能函数
*/

package function

import (
	"fmt"
)

func CacheCleaner() {
	var flag bool
	// 清除pip缓存
	pipArgs := []string{"cache", "purge"}
	flag = RunCommandGetFlag("pip", pipArgs)
	if flag == true {
		fmt.Printf("%7v %v\n\n", "-->", "Cleared pip cache")
	}
	// 清除npm缓存
	npmArgs := []string{"cache", "clean", "--force"}
	flag = RunCommandGetFlag("npm", npmArgs)
	if flag == true {
		fmt.Printf("%7v %v\n\n", "-->", "Cleared npm cache")
	}
	// 清除yarn缓存
	yarnArgs := []string{"cache", "clean"}
	flag = RunCommandGetFlag("yarn", yarnArgs)
	if flag == true {
		fmt.Printf("%7v %v\n\n", "-->", "Cleared yarn cache")
	}
}

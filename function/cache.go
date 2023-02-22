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
	"os/exec"
)

func CacheCleaner() {
	var (
		_   string
		err error
	)
	// 清除pip缓存
	_, err = exec.LookPath("pip")
	if err == nil {
		pipArgs := []string{"cache", "purge"}
		RunCommand("pip", pipArgs)
		fmt.Printf("%7v %v\n\n", "-->", "Cleared pip cache")
	}
	// 清除npm缓存
	_, err = exec.LookPath("npm")
	if err == nil {
		npmArgs := []string{"cache", "clean", "--force"}
		RunCommand("npm", npmArgs)
		fmt.Printf("%7v %v\n\n", "-->", "Cleared npm cache")
	}
	// 清除yarn缓存
	_, err = exec.LookPath("yarn")
	if err == nil {
		yarnArgs := []string{"cache", "clean"}
		RunCommand("yarn", yarnArgs)
		fmt.Printf("%7v %v\n\n", "-->", "Cleared yarn cache")
	}
}

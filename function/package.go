/*
File: package.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-02-21 14:01:12

Description: 子命令`package`功能函数
*/

package function

import (
	"fmt"
	"strings"
)

func PackageCleaner() {
	// 检查命令
	checkArgs := []string{"-Qtdq"}
	lonelyPackages := RunCommandGetResult("pacman", checkArgs)

	// Logo命令
	mascotArgs := []string{}
	mascot := RunCommandGetResult("repo-elephant", mascotArgs)

	// 检查命令结果解析
	if lonelyPackages == "" {
		fmt.Println("\033[35m[✔]\033[0m 没有孤立依赖包")
		fmt.Println(mascot)
	} else {
		// 卸载命令
		uninstallArgs := []string{"-Rn"}
		uninstallCmd := append(uninstallArgs, strings.Split(lonelyPackages, "\n")...)
		RunCommandGetFlag("pacman", uninstallCmd)
	}
}

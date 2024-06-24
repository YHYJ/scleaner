/*
File: package.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-02-21 14:01:12

Description: 子命令 'package' 的实现
*/

package cli

import (
	"os/exec"
	"strings"

	"github.com/gookit/color"
	"github.com/yhyj/scleaner/general"
)

// PackageCleaner 清理孤立软件包
//
// 参数：
//   - noLogoFlag: 是否不显示 Logo
func PackageCleaner(noLogoFlag bool) {
	// 检查命令
	checkArgs := []string{"-Qtdq"}
	lonelyPackages, err := general.RunCommandGetResult("pacman", checkArgs)
	if err != nil {
		if _, ok := err.(*exec.ExitError); !ok {
			fileName, lineNo := general.GetCallerInfo()
			color.Printf("%s %s -> Unable to find orphan packages: %s\n", general.DangerText("Error:"), general.SecondaryText("[", fileName, ":", lineNo+1, "]"), err)
		}
	}

	// 检查命令结果解析
	if lonelyPackages == "" {
		color.Printf("%s %s\n", general.SuccessText("[✔]"), general.LightText("没有孤立依赖包"))
		if !noLogoFlag {
			// 输出 Logo 的命令
			mascotArgs := []string{}
			mascot, err := general.RunCommandGetResult("repo-elephant", mascotArgs)
			if err != nil {
				fileName, lineNo := general.GetCallerInfo()
				color.Printf("%s %s -> Unable to get mascot: %s\n", general.DangerText("Error:"), general.SecondaryText("[", fileName, ":", lineNo+1, "]"), err)
			}
			color.Println(general.SuccessText(mascot))
		}
	} else {
		// 卸载命令
		uninstallArgs := []string{"-Rn"}
		uninstallCmd := append(uninstallArgs, strings.Split(lonelyPackages, "\n")...)
		if err := general.RunCommand("pacman", uninstallCmd); err != nil {
			fileName, lineNo := general.GetCallerInfo()
				color.Printf("%s %s -> Unable to clean orphan packages: %s\n", general.DangerText("Error:"), general.SecondaryText("[", fileName, ":", lineNo+1, "]"), err)
		}
	}
}

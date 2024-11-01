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
	// 查询命令
	queryArgs := []string{"-Qtdq"}
	lonelyPackages, _, err := general.RunCommandToBuffer("pacman", queryArgs)
	if err != nil {
		if _, ok := err.(*exec.ExitError); !ok {
			fileName, lineNo := general.GetCallerInfo()
			color.Printf("%s %s %s\n", general.DangerText(general.ErrorInfoFlag), general.SecondaryText("[", fileName, ":", lineNo+1, "]"), err)
		}
	}

	// 查询命令结果解析
	if lonelyPackages == "" {
		color.Printf("%s %s\n", general.SuccessText("[✔]"), general.LightText("没有孤立依赖包"))
		if !noLogoFlag {
			// 输出 Logo 的命令
			mascotArgs := []string{}
			mascot, _, err := general.RunCommandToBuffer("repo-elephant", mascotArgs)
			if err != nil {
				fileName, lineNo := general.GetCallerInfo()
				color.Printf("%s %s %s\n", general.DangerText(general.ErrorInfoFlag), general.SecondaryText("[", fileName, ":", lineNo+1, "]"), err)
			}
			color.Println(general.SuccessText(mascot))
		}
	} else {
		// 检查用户权限
		root := general.RootPermissions()
		// 执行命令
		var (
			command       string
			uninstallArgs []string
		)
		if root {
			command = "pacman"
			uninstallArgs = []string{"-Rn"}
		} else {
			command = "sudo"
			uninstallArgs = []string{"pacman", "-Rn"}

			// 检测是否安装了 sudo
			if _, errMessage := exec.LookPath(command); errMessage != nil {
				err := "This function requires superuser permissions. Please re-run as root"
				fileName, lineNo := general.GetCallerInfo()
				color.Printf("%s %s %s\n", general.DangerText(general.ErrorInfoFlag), general.SecondaryText("[", fileName, ":", lineNo+1, "]"), err)
				return
			}
		}
		// 卸载命令
		uninstallCmd := append(uninstallArgs, strings.Split(lonelyPackages, "\n")...)
		if err := general.RunCommandToOS(command, uninstallCmd); err != nil {
			fileName, lineNo := general.GetCallerInfo()
			color.Printf("%s %s %s\n", general.DangerText(general.ErrorInfoFlag), general.SecondaryText("[", fileName, ":", lineNo+1, "]"), err)
		}
	}
}

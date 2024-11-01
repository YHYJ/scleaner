/*
File: define_permissions.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2024-11-01 14:13:48

Description: 权限方面的功能
*/

package general

import (
	"os/exec"

	"github.com/gookit/color"
)

// RootPermissions 检查当前用户是否具有管理员权限
//
// 返回：
//   - 是否具有管理员权限
func RootPermissions() bool {
	currentUser, err := GetCurrentUserInfo()
	if err != nil {
		if _, ok := err.(*exec.ExitError); !ok {
			fileName, lineNo := GetCallerInfo()
			color.Printf("%s %s %s\n", DangerText(ErrorInfoFlag), SecondaryText("[", fileName, ":", lineNo+1, "]"), err)
		}
	}

	// 检查当前用户 ID 是否为 0 （root 用户）
	if currentUser.Uid == "0" {
		return true
	}
	return false
}

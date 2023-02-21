/*
File: package.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-02-21 14:01:12

Description:
*/

package function

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// 运行指定命令
func RunCommand(command string, args []string, display bool) string {
	// 定义命令
	cmd := exec.Command(command, args...)

	var stdin, stdout, stderr bytes.Buffer
	cmd.Stdin = &stdin   // 标准输入
	cmd.Stdout = &stdout // 标准输出
	cmd.Stderr = &stderr // 标准错误

	// 执行命令获取输出
	// err := cmd.Run()
	errS := cmd.Start()
	errW := cmd.Wait()
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())

	// 是否打印标准输出，必须放在判断命令执行情况语句的前面
	if display {
		fmt.Printf("%v", outStr)
	}

	// 判断命令执行情况
	if errS != nil || errW != nil {
		fmt.Println(errStr)
	}

	return strings.TrimRight(outStr, "\n")
}

func CheckPackages() {
	// 执行检查命令
	checkArgs := []string{"-Qtdq"}
	lonelyPackages := RunCommand("pacman", checkArgs, false)
	// 检查命令结果解析
	if lonelyPackages == "" {
		fmt.Println("没有需要清理的包")
	} else {
		// 执行卸载命令
		uninstallArgs := []string{"-Rn", "--noconfirm", lonelyPackages}
		RunCommand("pacman", uninstallArgs, true)
		// fmt.Printf("卸载: %v\n", lonelyPackages)
	}
}

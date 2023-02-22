/*
File: package.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-02-21 14:01:12

Description:
*/

package function

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// 运行指定命令并获取命令输出
func RunCommandGetReturn(command string, args []string, display bool) string {
	// 定义命令
	cmd := exec.Command(command, args...)
	// 执行命令并获取命令输出
	output, _ := cmd.Output()
	// 类型转换
	result := strings.TrimRight(string(output), "\n")

	// 是否打印标准输出，必须放在判断命令执行情况语句的前面
	if display {
		fmt.Printf("找到孤立依赖包：%v\n", result)
	}

	return result
}

// 运行指定命令
func RunCommand(command string, args []string) {
	// 定义命令
	cmd := exec.Command(command, args...)
	// 定义标准输入/输出/错误
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// 执行命令
	cmd.Run()
}

func CheckPackages() {
	// 检查命令
	checkArgs := []string{"-Qtdq"}
	lonelyPackages := RunCommandGetReturn("pacman", checkArgs, false)

	// Logo命令
	mascotArgs := []string{}
	mascot := RunCommandGetReturn("repo-elephant", mascotArgs, false)

	// 检查命令结果解析
	if lonelyPackages == "" {
		fmt.Println("\033[35m[✔]\033[0m 没有孤立的依赖包")
		fmt.Println(mascot)
	} else {
		// 卸载命令
		uninstallArgs := []string{"-Rn", lonelyPackages}
		RunCommand("pacman", uninstallArgs)
	}
}

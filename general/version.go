/*
File: version.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-02-21 13:58:00

Description: 子命令`version`功能函数
*/

package general

import "fmt"

// 程序信息
const (
	name    = "Scleaner"
	version = "v0.4.0"
	project = "github.com/yhyj/scleaner"
)

// 编译信息
var (
	gitCommitHash string = "unknown"
	buildTime     string = "unknown"
	buildBy       string = "unknown"
)

func ProgramInfo(only bool) string {
	programInfo := fmt.Sprintf("%s\n", version)
	if !only {
		programInfo = fmt.Sprintf("%s version: %s\nGit commit hash: %s\nBuilt on: %s\nBuilt by: %s\n", name, version, gitCommitHash, buildTime, buildBy)
	}
	return programInfo
}
# README

<!-- File: README.md -->
<!-- Author: YJ -->
<!-- Email: yj1516268@outlook.com -->
<!-- Created Time: 2023-02-20 16:29:54 -->

---

## Table of Contents

<!-- vim-markdown-toc GFM -->

* [Usage](#usage)
* [Compile](#compile)
  * [当前平台](#当前平台)
  * [交叉编译](#交叉编译)
    * [Linux](#linux)
    * [macOS](#macos)
    * [Windows](#windows)

<!-- vim-markdown-toc -->

---

<!--------------------------------------------->
<!--           _                             -->
<!--  ___  ___| | ___  __ _ _ __   ___ _ __  -->
<!-- / __|/ __| |/ _ \/ _` | '_ \ / _ \ '__| -->
<!-- \__ \ (__| |  __/ (_| | | | |  __/ |    -->
<!-- |___/\___|_|\___|\__,_|_| |_|\___|_|    -->
<!--------------------------------------------->

---

适用于Arch Linux的系统清理工具

## Usage

- `cache`子命令

    该子命令用于清除'pip/npm/yarn'的缓存文件

- `package`子命令

    该子命令用于清除作为依赖项安装但现在不再被任何包依赖的孤立包

- `version`子命令

    查看程序版本信息

- `help`

    查看程序帮助信息

## Compile

### 当前平台

```bash
go build -gcflags="-trimpath" -ldflags="-s -w -X github.com/yhyj/scleaner/function.buildTime=`date +%s` -X github.com/yhyj/scleaner/function.buildBy=$USER" -o scleaner main.go
```

### 交叉编译

使用命令`go tool dist list`查看支持的平台

#### Linux

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -gcflags="-trimpath" -ldflags="-s -w -X github.com/yhyj/scleaner/function.buildTime=`date +%s` -X github.com/yhyj/scleaner/function.buildBy=$USER" -o scleaner main.go
```

> 使用`uname -m`确定硬件架构
>
> - 结果是x86_64则GOARCH=amd64
> - 结果是aarch64则GOARCH=arm64

#### macOS

```bash
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -gcflags="-trimpath" -ldflags="-s -w -X github.com/yhyj/scleaner/function.buildTime=`date +%s` -X github.com/yhyj/scleaner/function.buildBy=$USER" -o scleaner main.go
```

> 使用`uname -m`确定硬件架构
>
> - 结果是x86_64则GOARCH=amd64
> - 结果是aarch64则GOARCH=arm64

#### Windows

```powershell
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -gcflags="-trimpath" -ldflags="-s -w -H windowsgui -X github.com/yhyj/scleaner/function.buildTime=`date +%s` -X github.com/yhyj/scleaner/function.buildBy=$USER" -o scleaner main.go
```

> 使用`echo %PROCESSOR_ARCHITECTURE%`确定硬件架构
>
> - 结果是x86_64则GOARCH=amd64
> - 结果是aarch64则GOARCH=arm64

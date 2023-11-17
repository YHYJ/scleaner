<h1 align="center">Scleaner</h1>

<!-- File: README.md -->
<!-- Author: YJ -->
<!-- Email: yj1516268@outlook.com -->
<!-- Created Time: 2023-02-20 16:29:54 -->

---

<p align="center">
  <a href="https://github.com/YHYJ/scleaner/actions/workflows/release.yml"><img src="https://github.com/YHYJ/scleaner/actions/workflows/release.yml/badge.svg" alt="Go build and release by GoReleaser"></a>
</p>

---

## Table of Contents

<!-- vim-markdown-toc GFM -->

* [Install](#install)
  * [一键安装](#一键安装)
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

适用于 Arch Linux 的系统清理工具

## Install

### 一键安装

```bash
curl -fsSL https://raw.githubusercontent.com/YHYJ/scleaner/main/install.sh | sudo bash -s
```

## Usage

- `cache`子命令

  该子命令用于清除'pip/npm/yarn'的缓存文件

- `package`子命令

  该子命令用于清除作为依赖项安装但现在不再被任何包依赖的孤立包

- `version`子命令

  查看程序版本信息

- `help`子命令

  查看程序帮助信息

## Compile

### 当前平台

```bash
go build -gcflags="-trimpath" -ldflags="-s -w -X github.com/yhyj/scleaner/general.GitCommitHash=`git rev-parse HEAD` -X github.com/yhyj/scleaner/general.BuildTime=`date +%s` -X github.com/yhyj/scleaner/general.BuildBy=$USER" -o build/scleaner main.go
```

### 交叉编译

使用命令`go tool dist list`查看支持的平台

#### Linux

```bash
CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -gcflags="-trimpath" -ldflags="-s -w -X github.com/yhyj/scleaner/general.GitCommitHash=`git rev-parse HEAD` -X github.com/yhyj/scleaner/general.BuildTime=`date +%s` -X github.com/yhyj/scleaner/general.BuildBy=$USER" -o build/scleaner main.go
```

> 使用`uname -m`确定硬件架构
>
> - 结果是 x86_64 则 GOARCH=amd64
> - 结果是 aarch64 则 GOARCH=arm64

#### macOS

```bash
CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -gcflags="-trimpath" -ldflags="-s -w -X github.com/yhyj/scleaner/general.GitCommitHash=`git rev-parse HEAD` -X github.com/yhyj/scleaner/general.BuildTime=`date +%s` -X github.com/yhyj/scleaner/general.BuildBy=$USER" -o build/scleaner main.go
```

> 使用`uname -m`确定硬件架构
>
> - 结果是 x86_64 则 GOARCH=amd64
> - 结果是 aarch64 则 GOARCH=arm64

#### Windows

```powershell
CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -gcflags="-trimpath" -ldflags="-s -w -H windowsgui -X github.com/yhyj/scleaner/general.GitCommitHash=`git rev-parse HEAD` -X github.com/yhyj/scleaner/general.BuildTime=`date +%s` -X github.com/yhyj/scleaner/general.BuildBy=$USER" -o build/scleaner.exe main.go
```

> 使用`echo %PROCESSOR_ARCHITECTURE%`确定硬件架构
>
> - 结果是 x86_64 则 GOARCH=amd64
> - 结果是 aarch64 则 GOARCH=arm64

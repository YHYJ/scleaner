<h1 align="center">Scleaner</h1>
<h3 align="center">用于 Arch Linux 的系统清理工具</h3>

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

* [适配](#适配)
* [安装](#安装)
  * [一键安装](#一键安装)
  * [编译安装](#编译安装)
    * [当前平台](#当前平台)
    * [交叉编译](#交叉编译)
* [用法](#用法)

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

## 适配

- Linux: 适配 Arch Linux 及基于 Arch Linux 的发行版
- macOS: 不适配
- Windows: 不适配

## 安装

### 一键安装

```bash
curl -fsSL https://raw.githubusercontent.com/YHYJ/scleaner/main/install.sh | sudo bash -s
```

也可以从 [GitHub Releases](https://github.com/YHYJ/scleaner/releases) 下载解压后使用

### 编译安装

#### 当前平台

如果要为当前平台编译，可以使用以下命令：

```bash
go build -gcflags="-trimpath" -ldflags="-s -w -X github.com/yhyj/scleaner/general.GitCommitHash=`git rev-parse HEAD` -X github.com/yhyj/scleaner/general.BuildTime=`date +%s` -X github.com/yhyj/scleaner/general.BuildBy=$USER" -o build/scleaner main.go
```

#### 交叉编译

> 使用命令`go tool dist list`查看支持的平台
>
> Linux 和 macOS 使用命令`uname -m`，Windows 使用命令`echo %PROCESSOR_ARCHITECTURE%` 确认系统架构
>
> - 例如 x86_64 则设 GOARCH=amd64
> - 例如 aarch64 则设 GOARCH=arm64
> - ...

设置如下系统变量后使用 [编译安装](#编译安装) 的命令即可进行交叉编译：

- CGO_ENABLED: 不使用 CGO，设为 0
- GOOS: 设为 linux
- GOARCH: 根据当前系统架构设置

## 用法

- `cache`子命令

  清除'pip/npm/yarn'的缓存文件

- `package`子命令

  清除作为依赖项安装但现在不再被任何包依赖的孤立包

- `version`子命令

  查看程序版本信息

- `help`子命令

  查看程序帮助信息

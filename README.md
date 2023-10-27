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

- 编译当前平台可执行文件：

```bash
go build main.go
```

- **交叉编译**指定平台可执行文件：

```bash
# 适用于Linux AArch64平台
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build main.go
```

```bash
# 适用于macOS amd64平台
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go
```

```bash
# 适用于Windows amd64平台
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
```

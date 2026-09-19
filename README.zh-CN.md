[English](README.md) · **简体中文**

> 英文版是规范版本。本页与 [README.md](README.md) 不一致时，以英文版为准。

<!-- translation-of: README.md sha256:13ea747437e5639f -->

<!-- Source: Best-README-Template BLANK_README (Unlicense) — https://github.com/othneildrew/Best-README-Template -->
<a id="readme-top"></a>

# ShiyanlouRePostpone

一个个人使用的 Go 命令行工具，通过实验楼平台自己的接口查询某个实验任务的剩余时间，并在剩余时间不足时提交延时请求，直到用完设定的延时次数为止。

[![License](https://img.shields.io/github/license/anyingiit/shiyanlouRePostpone)](LICENSE)

[报告问题](https://github.com/anyingiit/shiyanlouRePostpone/issues/new?template=bug_report.yml) · [提出需求](https://github.com/anyingiit/shiyanlouRePostpone/issues/new?template=feature_request.yml)

<details>
  <summary>目录</summary>
  <ol>
    <li><a href="#about-the-project">关于本项目</a></li>
    <li><a href="#getting-started">开始使用</a></li>
    <li><a href="#usage">用法</a></li>
    <li><a href="#contributing">参与贡献</a></li>
    <li><a href="#license">许可证</a></li>
    <li><a href="#contact">联系方式</a></li>
  </ol>
</details>

## 关于本项目

ShiyanlouRePostpone 是仓库根目录下的一个单文件 Go 程序，用来自动延长实验楼在线实验平台上一个实验任务的会话时间。它调用实验楼自己的 `labtask` 接口读取当前任务的剩余分钟数，一旦剩余时间不足八分钟，就调用对应的延时接口把截止时间往后推，如此循环，直到用完由 `-max` 参数设定的延时次数为止。`testProgram/` 目录下另有一个互不相关的文件，只是用来试验 Go 语言 `flag` 包用法的练习代码，并不会请求实验楼的任何接口。

计划中的功能与已知问题，见 [open issues](https://github.com/anyingiit/shiyanlouRePostpone/issues)。

## 开始使用

### 环境要求

- Go（较新的、支持模块（module）特性的版本；仓库没有提交 `go.mod`，因此没有锁定最低版本）
- 可以访问外网以拉取 `github.com/pkg/errors`，这是该程序唯一用到的第三方依赖
- 你自己的实验楼登录会话 Cookie：仓库中原本写死会话 Cookie 的那一行请求头已被脱敏移除，所以在你补上自己的 Cookie 之前，代码无法编译通过（见下方"安装"）

### 安装

```sh
git clone https://github.com/anyingiit/shiyanlouRePostpone.git
cd shiyanlouRePostpone
go mod init shiyanlourepostpone
go mod tidy
```

仓库根目录的 Go 文件中，有两个函数里被脱敏移除了写死的实验楼会话 Cookie 请求头；在能够编译之前，需要把这一行请求头用你自己的 Cookie 补回去：

```go
req.Header.Add("Cookie", "<your-own-shiyanlou-session-cookie>")
```

## 用法

```sh
go run 实验楼时间延长_2.go -max 5
```

上面的命令以 `-max` 参数为 5（默认值为 3）启动前面描述的循环：程序会检查任务剩余时间，等到剩余时间不足八分钟时发起延时，如此重复，直到用完这个次数后自动退出。

## 参与贡献

欢迎参与。[CONTRIBUTING.md](CONTRIBUTING.md) 说明如何提交 issue 或 pull request，[CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md) 说明对所有参与者的行为要求。

请不要在公开的 issue 或 pull request 中报告安全问题。[SECURITY.md](SECURITY.md) 说明了私下报告的方式。

## 许可证

以 MIT 许可证分发。详见 [LICENSE](LICENSE)。

## 联系方式

项目地址：[https://github.com/anyingiit/shiyanlouRePostpone](https://github.com/anyingiit/shiyanlouRePostpone)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

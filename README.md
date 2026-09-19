<!-- Source: Best-README-Template BLANK_README (Unlicense) — https://github.com/othneildrew/Best-README-Template -->
<a id="readme-top"></a>

# ShiyanlouRePostpone

A personal Go command-line tool that polls the Shiyanlou learning platform's own API for a lab task's remaining time and submits an extension request to push back the deadline until a configured number of extensions have been used.

**English** · [简体中文](README.zh-CN.md)

[![License](https://img.shields.io/github/license/anyingiit/shiyanlouRePostpone)](LICENSE)

[Report a bug](https://github.com/anyingiit/shiyanlouRePostpone/issues/new?template=bug_report.yml) · [Request a feature](https://github.com/anyingiit/shiyanlouRePostpone/issues/new?template=feature_request.yml)

<details>
  <summary>Table of Contents</summary>
  <ol>
    <li><a href="#about-the-project">About The Project</a></li>
    <li><a href="#getting-started">Getting Started</a></li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>

## About The Project

ShiyanlouRePostpone is a single Go program at the repository root that automates renewing a lab session on the Shiyanlou online-course platform. It calls Shiyanlou's own `labtask` API to read the minutes remaining on the current task, and once fewer than eight minutes are left it calls the matching extend endpoint to push the deadline back, repeating until a `-max`-flag-controlled number of extensions have been used. A second, unrelated file under `testProgram/` is a scratch program for experimenting with Go's `flag` package and does not talk to Shiyanlou at all.

See the [open issues](https://github.com/anyingiit/shiyanlouRePostpone/issues) for planned features and known issues.

## Getting Started

### Prerequisites

- Go (a recent release with module support; no `go.mod` is committed, so no minimum version is pinned)
- Network access to fetch `github.com/pkg/errors`, the program's one third-party import
- A Shiyanlou session cookie of your own: the request-header line that once carried a hardcoded one has been redacted from the repository's Go source file, so it will not compile until you add your own (see Installation)

### Installation

```sh
git clone https://github.com/anyingiit/shiyanlouRePostpone.git
cd shiyanlouRePostpone
go mod init shiyanlourepostpone
go mod tidy
```

The Go file at the repository root had a hardcoded Shiyanlou session-cookie header removed from two functions; add the missing header back with your own cookie before it will build:

```go
req.Header.Add("Cookie", "<your-own-shiyanlou-session-cookie>")
```

## Usage

```sh
go run 实验楼时间延长_2.go -max 5
```

This starts the loop described above with `-max` set to 5 (the default is 3): it checks the task's remaining time, waits until fewer than eight minutes are left, extends it, and repeats up to that many times before exiting on its own.

## Contributing

Contributions are welcome. Read [CONTRIBUTING.md](CONTRIBUTING.md) for how to open an issue or a pull request, and [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md) for the standards expected of everyone taking part.

Please do not report security issues in public issues or pull requests. [SECURITY.md](SECURITY.md) explains how to report them privately.

## License

Distributed under the MIT License. See [LICENSE](LICENSE) for details.

## Contact

Project link: [https://github.com/anyingiit/shiyanlouRePostpone](https://github.com/anyingiit/shiyanlouRePostpone)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

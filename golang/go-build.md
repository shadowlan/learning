# Go 语言交叉编译和构建标签

现代应用支持多平台运行是一件稀松平常的事情，在 Go 语言里面，为了支持应用的多平台部署，给用户提供了方便的配置方式
来轻松构建针对不同操作系统和平台的运行文件。

## Go 构建标签

Go 的构建约束，即构建标签，是以`// go:build`为开始的行注释，如果是 1.16 或之前的版本，格式是`// +build`。
跟此变更相关的 issue 可以参考[25348](https://github.com/golang/go/issues/25348), 以及相关的[proposal](https://go.googlesource.com/proposal/+/master/design/draft-gobuild.md)。

构建标签必须出现在 package 子句之前。为了区分构建标签和包文档的描述注释，构建标签后面应该有一个空行。

构建标签由||, &&, !运算符以及括号来组合表达。运算符与 Go 中的含义相同。
例如，以下构建标签在满足`linux`和`386`约束，或者满足`darwin`而`cgo`不满足时构建文件：
`//go:build (linux && 386) || (darwin && !cgo)`

又如：仅在使用 cgo 时，且仅在 Linux 和 OS X 上构建文件：`//go:build cgo && (linux || darwin)`

注意：1.17 及以后的表达格式里，一个文件有多个 `//go:build` 行是错误的。

在 1.16 及以前的版本，多行构建标签是允许的，并且组合方式是通过空格和逗号等来区分，空格符表示 OR，逗号表示 AND,感叹号表示 NOT。
而多行之间则表示 OR。gofmt 命令将在遇到旧语法时添加等效的 `//go:build` 约束。如下是示例：

| 标签                                      | 约束含义                                  |
| ----------------------------------------- | ----------------------------------------- |
| // +build linux,386 darwin,!cgo           | (linux AND 386) OR (darwin AND (NOT cgo)) |
| // +build linux darwin<br>// +build amd64 | (linux OR darwin) AND amd64               |

如果文件名在去除扩展名和可能的`_test`后缀后匹配以下任何模式, （例如：source_windows_amd64.go）其中 GOOS 和 GOARCH 分别
代表任何已知的操作系统和体系结构值，那么认为该文件除了文件中的任何显式约束之外，具有这些术语的所表达的隐式构建标签。

- `_GOOS`
- `_GOARCH`
- `_GOOS_GOARCH`

要使文件构建时被忽略，可以使用：`//go:build ignore`，其他任何没有被用来定义为标签的词也可以，但"ignore"是约定俗成的。）

Go 语言目前支持的系统和架构可以参考[官方文档](https://go.dev/doc/install/source#environment)。
除了官方提供的针对不同平台的内置标签，用户也可以使用自定义标签，例如`//go:build prod`, 只需要在执行`go build`时显式带上
标签名`go build --tags=prod`.

## 参考文章

- [官方介绍](https://pkg.go.dev/cmd/go#hdr-Build_constraints)
- [Go (Golang) GOOS and GOARCH](https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63)
- [Using Go's build tags](https://wawand.co/blog/posts/using-build-tags/)

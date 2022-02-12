# 使用pprof

## 什么是profile

golang官方runtime/pprof里有对于profile的[解释](https://golang.org/pkg/runtime/pprof/#Profile),暂且理解为应用的各类资源的画像，而我们通常关注的主要是CPU/Memory/Goroutine的profile.

## 如何收集

针对应用资源的运行情况，如何收集数据并产生profile？golang本身提供了两个库来实现，分别是runtime/pprof 和 net/http/pprof。

1. runtime/pprof

使用runtime/pprof，需要通过编写代码来实现，通常针对是非服务器类型的应用，应用运行结束退出时生成profile文件，供后续分析。以下是收集cpu和mem的资源运行情况的示例代码,一般把部分内容写在 main.go文件中：
```golang
   cpuf, err := os.Create("cpu_profile")
    if err != nil {
        log.Fatal(err)
    }
    pprof.StartCPUProfile(cpuf)
    defer pprof.StopCPUProfile()
    // the application code
    ...
    memf, err := os.Create("mem_profile")
    if err != nil {
        log.Fatal("could not create memory profile: ", err)
    }
    if err := pprof.WriteHeapProfile(memf); err != nil {
        log.Fatal("could not write memory profile: ", err)
    }
```

2. net/http/pprof

一般针对服务器类型应用，通常使用net/http/pprof更为简单。如果使用了默认的http.DefaultServeMux(通常是代码直接使用http.ListenAndServe("0.0.0.0:8000", nil)），只需要添加一行`import _ "net/http/pprof"`,如果使用自定义的Mux，则需要手动注册一些路由规则：
```
r.HandleFunc("/debug/pprof/", pprof.Index)
r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
r.HandleFunc("/debug/pprof/profile", pprof.Profile)
r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
r.HandleFunc("/debug/pprof/trace", pprof.Trace)
```

无论是默认还是自定义Mux，应用启动后，都可以通过浏览器访问http://127.0.0.1:8080/debug/pprof/ 来获得当前运行在服务器端profile信息。

## 如何分析数据

go tool pprof是分析profile数据的命令行工具。注意获取profiling数据是动态的，为了获取有效数据，需要保证应用处于有负载的情况下，特别是在高负载的情况下能够暴露更多问题。

pprof工具生成的调用关系图和火焰图都需要用到graphviz，这里给出ubuntu和mac下的安装命令：

```shell
apt-get install -y graphviz
brew install graphviz
```

基本命令为`go tool pprof [binary] [source]`，binary是应用的二进制文件，用来解析各种符号；source表示profile数据的来源，可以是本地的文件，也可以是http地址。比如：http://localhost:8080/debug/pprof/profile 默认采集时间为30s，如果需要调整采集数据的时间，可以在url后加上seconds参数，例如`?seconds=60`。

### CPU Profiling

命令示例：`go tool pprof ./app http://localhost:8080/debug/pprof/profile\?seconds\=60`， 在终端执行命令后需要等待一会儿，然后进入交互模式。可以输入`help`来查看支持那些命令。常用的有topN，找到最消耗CPU的代码函数。用`web`命令，能自动生成一个svg文件，并跳转到浏览器打开，生成了一个函数调用图。
`list`命令后面跟着某个函数名或一个正则表达式，就能查看匹配函数的代码以及每行代码的耗时。或者用`weblist`，类似list，但是会将结果从默认浏览器中打开，更便于查看。

### Memory Profiling

命令示例：`go tool pprof ./app http://localhost:8080/debug/pprof/heap`, 上述针对cpu profiling中使用的方法对memory profiling同样有效，另外`tree`命令可以在命令行下以文本方式显示调用关系图。

### 火焰图

火焰图（Flame Graph）是Bredan Gregg 创建的一种性能分析图表，前面介绍的profiling数据结果也可以转换成火焰图来分析，优点是它是动态的，可以通过点击每个方块来分析它上面的内容。火焰图的调用顺序从下到上，每个方块代表一个函数，它上面一层表示这个函数会调用哪些函数，方块的大小代表了占用 CPU 使用的长短。火焰图的配色并没有特殊的意义，默认的红、黄配色是为了更像火焰而已。这里介绍使用go-torch工具来生成火焰图。安装go-torth及其依赖,克隆FlameGraph后需要将该文件夹添加到$PATH以便于go-torch调用依赖的脚本。

```shell
go get github.com/uber/go-torch
git clone git@github.com:brendangregg/FlameGraph.git
```

go-torch在没有任何参数的情况下会尝试从 http://localhost:8080/debug/pprof/profile 获取profiling数据。它有三个常用的参数可以调整：

* -u --url：要访问的 URL，这里只是主机和端口部分
* -s --suffix：pprof profile 的路径，默认为 /debug/pprof/profile
* --seconds：要执行 profiling 的时间长度，默认为 30s

运行命令`go-torch http://localhost:8080/debug/pprof/heap`后，默认输出svg文件`torch.svg`到当前目录下。mac下可以直接运行`open torch.svg`打开，如果希望用浏览器可以执行`open -a 'Firefox' torch.svg`.

```shell
go-torch http://localhost:8080/debug/pprof/heap
INFO[21:11:46] Run pprof command: go tool pprof -raw -seconds 30 http://localhost:8080/debug/pprof/heap
INFO[21:11:50] Writing svg to torch.svg`
```

## 参考

* [使用pprof和火焰图调试golang应用](https://cizixs.com/2017/09/11/profiling-golang-program/)
* [Profiling Go programs with pprof](https://jvns.ca/blog/2017/09/24/profiling-go-with-pprof/)
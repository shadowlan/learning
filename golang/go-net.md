# Go net

net包属于golang里的[标准库](https://golang.org/pkg/#stdlib)，根据标准库里的描述,net包提供一个轻便的网络IO接口，包括TCP/IP，UDP，域名解析和Unix域套接字。中文版的库详情也可以在腾讯云的[开发者手册](https://cloud.tencent.com/developer/section/1143223)找到。最近项目里用到net包比较多，决定近距离看一看这个包能做什么，目标是提供一个client和server示例代码，以展示如何利用net来构建一个可用的简单服务器。

net中常用的方法如下：

1. func Dial(network, address string) (Conn, error)
2. func DialTCP(network string, laddr, raddr *TCPAddr) (*TCPConn, error)
3. func Listen(network, address string) (Listener, error
4. func ListenUDP(network string, laddr *UDPAddr) (*UDPConn, error)
5. func DialUDP(network string, laddr, raddr *UDPAddr) (*UDPConn, error)

所有版本的`net.Dial()`和`net.Listen()`返回的数据类型都实现了`io.Reader`和`io.Writer`接口，这意味着我们可以用常规的文件I/O函数从一个tcp/ip连接发送
接受数据。

## 创建tcp服务端

建立一个tcp服务器，监听端口为8080，接受用户的数据并返回用户数据和当前的调用时间，具体代码在[这里](go-net/tcp/tcpServer.go).通过`go build tcpServer.go`构建二进制文件后，用`./tcpServer` 来启动该服务。用netcat简单测试下是否如预期工作：`nc "hello" | nc localhost 8080` 或者 `nc localhost 8080 <<< hello`  
ps: 像net的子包http里常被用来建立http server端的方法`http.ListenAndServe()`也是调用了`net.Listen("tcp", addr)`来创建一个tcp服务监听端。

## 创建tcp客户端

客户端代码提供两种方式来和tcp服务端交互，发送和接收数据分别使用了常规I/O函数方式和tcpconn本身提供的`write`&`read`方法。具体代码在[这里](go-net/tcp/tcpClient.go)

## 创建udp服务端

建立一个udp服务器，监听端口为8088，接受用户的数据并返回用户数据和当前的调用时间，具体代码在[这里](go-net/udp/udpServer.go). 通过`go build udpServer.go`构建二进制文件后，用`./udpServer` 来启动该服务。用netcat简单测试下是否如预期工作：`echo "hi" | nc -u localhost 8088`
ps: 如果地址使用`:0`,那么将随机分配端口给服务端，可以通过命令`ps aux | grep udpServer | grep -v grep | awk '{print $2}' | xargs -n 1 lsof -p | grep UDP`来查看具体的端口。

## 创建udp客户端

客户端示例代码在[这里](go-net/udp/udpClient.go),类似tcp客户端，udp客户端接受标准输入的数据并发送给服务器，接收到服务器数据后打印到标准输出。

## 参考

* [developing-udp-and-tcp-clients-and-servers-in-go](https://www.linode.com/docs/development/go/developing-udp-and-tcp-clients-and-servers-in-go/)  
* [深入Go UDP编程](https://colobu.com/2016/10/19/Go-UDP-Programming/)  
* [stack-overflow-Q&A](https://stackoverflow.com/questions/24933352/difference-between-listen-read-and-write-functions-in-the-net-package)  
* [golang net包基础解析](https://blog.csdn.net/Wu_Roc/article/details/77169838)  
* [Go语言TCP Socket编程](https://tonybai.com/2015/11/17/tcp-programming-in-golang/)  

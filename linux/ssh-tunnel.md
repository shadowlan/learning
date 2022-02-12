# ssh隧道

有些环境网络限制比较严格，常常遇到仅仅开放ssh端口22的情况，在这种情况下目标机器上的很多端口都无法通过本地工作机器直接访问，必须借ssh隧道来完成。这里简单记录经常使用的隧道转发方式，便于以后参考使用。

## 本地转发

`ssh -CfNg -L listen_port:DST_Host:DST_port user@Tunnel_Host`

1. 建立本机独享隧道

```bash
ssh -L 8080:10.0.0.12:80 -Nf user@10.0.0.11
# -L 将本地8080端与的80端口建立映射关系
# -N 不执行命令或者脚本，否则会等待用户输入命令
# -f 不登录到主机，一般-Nf搭配使用即后台运行且不登录
ssh -L 8080:localhost:80 -Nf user@10.0.0.11
直接将隧道机10.0.0.11的80端口映射到本地8080
```

2. 建立共享隧道

```bash
ssh  -g  -L 8080:10.0.0.12:80 -Nf user@10.0.0.11
# -g 共享8080端口:别人访问本地工作机端口也可以直接使用隧道,否则只能本机使用。
```

## 远程转发

`ssh -CfNg -R listen_port:DST_Host:DST_port user@Tunnel_Host`

1. 有时由于防火墙的原因，本地host1无法访问host2，但是host2可以访问host1时，可以使用远程转发,从而使host1 可以访问host2的服务
在host2上执行`ssh -R 8080:localhost:80 host1`，使得用户在host1上可以使用通过访问8080端口来访问host2上的80端口。
2. 如果用户在外网，试图访问公司内网的服务器A的端口389但无法直接连接，同时用户可以访问公司外网服务器B，可以利用服务器B，在A和B之间创建一条ssh隧道，让用户本地能够从外面访问A。在内网服务器A上执行：
ssh -CfNg -R listen_port:ServerA:389 user@ServerB
此时用户再访问ServerB的listen_port，就可以访问到ServerA上389端口的服务。

## 动态转发

`ssh -CfNg -D listen_port user@Tunnel_Host`

本地转发，远程转发的前提都是要求有一个固定的应用服务端的端口号，如果没有这个端口号，例如用浏览器进行Web浏览的时候，可以使用动态转发来保护浏览行为。
`ssh -CfNg -D listen_port user@Tunnel_Host`,这里SSH实际是创建了一个SOCKS代理服务。在浏览器上可以直接使用localhost:listen_port 来作为正常的SOCKS代理来使用

## 参考

https://www.ibm.com/developerworks/cn/linux/l-cn-sshforward/  
https://www.jianshu.com/p/f2156c444fd6

_20191014_
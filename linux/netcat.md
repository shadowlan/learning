# Linux Netcat

netcat 作为网络工具中的瑞士军刀，能做的事情非常之多，在这里记录下平时遇到问题时用过的方式，便于以后需要的时候再翻一翻。

* 端口扫描 `nc -z -v -n $ip 21-500`
* 发送消息   
  服务器端: `nc -l 1234`,在1234端口启动一个tcp服务器  
  客户端: `nc $ip 1234`,在客户端的键入的任何消息都会显示在服务器端
* 文件传输  
  `nc -l 1234 < file`  
  `nc -n $ip 1234 > file`
* 打开一个shell  
  nc支持-e参数:  
    server: `nc -l 1234 -e /bin/bash -i`  
    client: `nc $ip 1234`  
  nc不支持-e参数:  
    server:  
    ```
    mkfifo /tmp/tmp_fifo
    cat /tmp/tmp_fifo | /bin/sh -i 2>&1 | nc -l 1234 > /tmp/tmp_fifo
    ```
    client: `nc -n $ip 1234`

## 参考
[cheetsheet](https://www.sans.org/security-resources/sec560/netcat_cheat_sheet_v1.pdf)  
[Linux Netcat 命令](https://www.oschina.net/translate/linux-netcat-command?lang=chs&p=1)

_20191008_
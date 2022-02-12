# git hook

给自己定的目标是每天的记录至少100字，于是想着应该让检查自动化，git是提供客户端和服务端钩子功能的，在commit之前我需要检查checkin的文件字数是否大于100.于是决定用pre-commit来实现这个功能。

git hook 文件默认都放在.git/hooks下面，打印文件列表发现还是有挺多类型的hook的，今天就专注在实现一个pre-commit吧。
平时统计文件字数一般都用wc，但是发现wc不带任何参数的时候，统计中文字数变成了三倍，比如统计`中文`二字：
```
echo -n 中文 | wc
6
```
原来不带参数的时候默认是以字节数计算的，而在不同的字符集编码下，中文对应的字节数是不同的，查了下本地mac系统的字符集`echo $LANG`, 结果是`en_US.UTF-8`,而在utf-8编码格式下，单个英文字符占一个字节，单个中文汉字占3个字节(参考的[这里](https://blog.csdn.net/yaomingyang/article/details/79374209))。所以统计汉字个数的时候两个汉字就变成了6，看了下wc的帮助文档，找到`wc -m`,支持统计多字节字符。万事俱备，可以继续我的pre-commit逻辑了。
看pre-commit.sample里的内容，首行是`#!/bin/sh`解析,所以shell语言语法就行.

在.git/hooks文件夹下，添加count文件内容如下：
```
#!/bin/sh
total=$(cat $1 | wc -m)
if [ $total -lt 100 ];then
echo "commit file length is less than 100"
exit 1
fi
```

将pre-commit.sample重命名为pre-commit,默认内容保留，添加如下检查：
```
basedir=$(dirname "$0")
git diff --cached --name-only -z $against | xargs -n 1 -0 -I {} "${basedir}"/count {}
```

_20190930_
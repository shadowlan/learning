# 如何查看linux os版本

每次需要检查linux版本的时候第一个想起来的总是`lsb_release`,但是发现有时候并不好使，大概率是相关的package没有安装。这个时候又要求助于搜索。但实际上查看linux os版本的方式无非也就几种，为了未来不费力快速得到相关信息，我参考了nixCraft里的[一篇文章](https://www.cyberciti.biz/faq/how-to-check-os-version-in-linux-command-line/)，在目前工作里常用的ubuntu 16.04和alpine 3.10里测试了相关命令。记录如下。

## 常用方式

* cat /etc/os-release
* lsb_release -a
* hostnamectl
* cat /proc/version

最后发现只有`lsb_release -a`在alpine下是不work的，其他方式在两个系统里都能顺利找到版本信息。我在ubuntu16.04里用`dpkg -S lsb_release`命令查看了下lsb_release命令所属包是`lsb-release`,看来在alpine 3.10里默认并没有装这个包。

## 小确幸

`dpkg -S`是在已经安装的包里查找文件，如果要查找并没有安装的包，可以利用`apt-file`,安装和使用如下：
```shell
apt-get install -y apt-file
apt-file update
apt-file search lsb_release
```
另附alpine里apk的[常用命令介绍指南](https://www.cyberciti.biz/faq/10-alpine-linux-apk-command-examples/)
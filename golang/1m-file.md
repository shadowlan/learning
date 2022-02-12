# 1M文件

看到kafka默认的消息大小是1M，于是想到如果要生成一个1M的大小的文件在linux或者通过golang要怎么做呢？今天就来实验实验吧，1M按照1024*1024字节来计算。

## golang

在golang里怎么生成一个大小为1M的文件呢？golang本身提供byte字节类型，1M=1024*1024byte，那么我们只需要给文件写入1024*1024字节的字符就行，用26个英文单词来替代的话就是1024*1024/26,大概是重复40329次字母表。a对应的unicode编码值为97，z的编码值为122，为了不手动输入字母表，遍历97到122之间的数值转换为字母，然后重复输入40329到一个文件，写完后再来读取文件大小。

示例代码在[这里](1m-file/main.go),没想到挺简单的功能也遇到了一些坑，说明自己对golang还是不太熟悉，在这里顺便记录下遇到的坑：

1. 本来尝试用ioutil.WriteFile的方式写入，`ioutil.WriteFile("1mfile", （[]byte)("content"), os.ModeAppend)`,但是发现这种模式下无法在循环里实现内容追加，因为每次之前的内容都会被清空。官方文档里写的是"WriteFile writes data to a file named by filename. If the file does not exist, WriteFile creates it with permissions perm; otherwise WriteFile truncates it before writing."
`
1. 调用os.openfilefile的时候先是写的`file, err := os.OpenFile("1mfile", os.O_CREATE, 0644)`,结果得到`write 1mfile: bad file descriptor`的报错。原来这个命令仅仅是创建了文件，但是却没有明确对这个文件是否能写入，改成`file, err := os.OpenFile("1mfile", os.O_CREATE|os.O_RDWR, 0644)`就修复了。

## linux

在linux里一般用`dd`来生成制定大小的文件。 要产生1m大小的文件，执行`dd if=/dev/zero of=1mfile bs=1M count=1`,`bs=1M`指定了input和output的blocksize都为1M，`count=1`表示拷贝1个输入块。


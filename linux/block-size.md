# 硬盘块

## ls -s

回想了一下为啥忽然要看硬盘块的内容，block其实是磁盘相关的一个非常基础的知识，只不过从没仔细研究过。昨天创建了一个1M大小的文件后，用ls -alts 1mfile去显示文件具体信息的时候，发现mac和linux输出不太一样，于是想八一八两者之间的区别。

先把两者的输出贴在这里：

```bash
#linux
ls -alts 1mfile
1028 -rw-r--r-- 1 root root 1048580 Nov  6 07:15 1mfile
#macos
ls -alts 1mfile
2056 -rw-r--r--  1 mhub  wheel  1048580 Nov  6 19:49 1mfile
```

鉴于不同操作系统上的同一个命令经常有差异，估计是`-s`这个参数导致了最前面数据的差别，于是看了下帮助文档后发现了两者的区别:

* MacOS: 显示实际使用的文件系统块数量，单位是512bytes，如果非整的话向上取整。如果是向终端输出，最上面会有所有文件的总和。而环境变量BLOCKSIZE可以覆盖默认的512字节单位。
* Linux: 打印每个文件分配的块大小。

这样看来这个数据也很好理解，应该为块基本单位大小的整数倍，但是按照文件大小1048580字节来算，MacOS的值也不对，除的结果是2048.0078,怎么也到不了2056。于是建了一个空文件来瞧一瞧大小是多少,结果发现在Mac上touch建出来的文件最前显示的值为0，只有当我在文件里输入一个字母或者空字符的时候，前面的值才发生变化，变成了8，一个空字符不可能占有8*512个字节的，想起上次学习Linux性能优化课程的时候有提到`Linux文件系统为每个文件都分配两个数据结构，索引节点（index node）和目录项（directory entry）。它们主要用来记录文件的元信息和目录结构。`,MacOS应该也会为文件做类似的事情，我猜想可能8个存储块就是被用来存储这些内容。而2048+8=2056，貌似也证实了这一点。 而在我的linux机器上，block size是1024，单个空字符文件默认占有4个块，1024+4=1028，这个值也符合输出。

## block size

磁盘块/簇（虚拟出来的）。块是操作系统中最小的逻辑存储单位。操作系统与磁盘打交道的最小单位是磁盘块。那么既然块是操作系统层面的，那么不同操作系统完全可能用不同尺寸的block size，这里就记录下在linux和mac下，怎么查看block size。

```bash
# mac
diskutil list
diskutil info /dev/disk1 | grep "Block Size"
   Device Block Size:         4096 Bytes
   Allocation Block Size:     4096 Bytes
# linux
df -l
tune2fs -l /dev/sda1 | grep Block
Block count:              498688
Block size:               1024
Blocks per group:         8192
```

## 参考

[硬盘基本知识](https://www.jianshu.com/p/9aa66f634ed6)
[扇区与块](https://blog.csdn.net/terryliu98/article/details/26479359)
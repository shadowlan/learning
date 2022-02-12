# 索引节点 vs 文件描述符

## inode
Linux 文件系统为每个文件都分配两个数据结构，索引节点（index node）和目录项（directory entry）。它们主要用来记录文件的元信息和目录结构
* 索引节点，简称为 inode，用来记录文件的元数据，比如 inode 编号、文件大小、访问权限、修改日期、数据的位置等。索引节点和文件一一对应，它跟文件内容一样，都会被持久化存储到磁盘中。所以记住，索引节点同样占用磁盘空间。
* 目录项，简称为 dentry，用来记录文件的名字、索引节点指针以及与其他目录项的关联关系。多个关联的目录项，就构成了文件系统的目录结构。不过，不同于索引节点，目录项是由内核维护的一个内存数据结构，所以通常也被叫做目录项缓存。

索引节点是每个文件的唯一标志，而目录项维护的正是文件系统的树状结构。目录项和索引节点的关系是多对一，你可以简单理解为，一个文件可以有多个别名。

### 常用命令

* `stat $file` 查看文件数据信息
* `stat --format=%i $file` `ls -i $file` 仅打印inode号
* `df -i` 查看文件系统inode空间信息，当一个分区有大量的小文件时，可能会先耗尽inode数量而不是磁盘空间，该命令有助于理解该类情况。

## 文件描述符

一个打开的文件可能是常规文件，文件夹，块文件，库，流或者网络文件等。而文件描述符是程序用于获得一个文件句柄的数据结构，最被熟知的有：
```
0 标准输入
1 标准输出
2 标准错误
```

当前打开的文件数不一定等于打开文件描述符数，例如当前的工作文件夹，内存映射文件和执行文本文件等。

### 常用命令
* `lsof -p $pid` 查看某进程打开的文件。
* `ls -l /proc/$pid/fd` 查看该进程使用的文件描述符。
* `cat /proc/sys/fs/file-max`，`sysctl fs.file-max` 查看操作系统的最大文件描述符限制。
* `cat /proc/sys/fs/file-nr`，`sysctl fs.file-nr` 查看操作系统当前在使用的文件描述符。
* `ulimit -a` 获得针对每个进程的文件描述符限制。

# 两者区别

一个有效的文件描述符关联文件模式标志和偏移量，根据文件描述符的如何获得的，给予进程读写的权限。同时也记录文件内的一些读写位置等，但是不包含任何文件本身的元数据信息，例如时间戳，unix权限比特位。一个inode包含时间戳和unix权限比特位，但是不包含文件模式标志或偏移量。

# 参考

[Open Files / Open File Descriptors](https://www.thegeekdiary.com/linux-interview-questions-open-files-open-file-descriptors/)  
[Linux性能优化](https://time.geekbang.org/column/article/76876)  
[what's an inode in linux](https://linoxide.com/linux-command/linux-inode/#What_is_an_inode_in_Linux)  
[inode vs file descriptor](https://www.quora.com/Whats-the-difference-between-inode-number-and-file-descriptor)  
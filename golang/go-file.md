# Go 文件操作

1. 打开文件`os.Open(name string)`,如果文件不存在，则返回错误，如果想在文件不存在的时候创建文件可以使用方法`OpenFile(name string, flag int, perm FileMode)`，通过OpenFile返回的错误可以用os.IsPermission(err)来判断对该文件是否有相应的权限
1. 获得文件的基本信息`os.Stat()`,该方法返回os.FileInfo，包括文件的名字，大小，读写权限，修改时间等信息，如果文件不存在，会返回`os.IsNotExist`的error，所以可以通过判断`os.Stat()`的error来确定文件是否存在
1. 从指定位置读取文件内容，可以通过`Seek(offset int64, whence int)`方法获得新的指定位置，然后利用`ReadAt(b []byte, off int64)`开始读取文件。
1. 按照行/字符等读取，可以利用`bufio.Scanner`和`scanner.Split(bufio.ScanLines)`,`scanner.Split(bufio.ScanWords)`，结合`scanner.Scan()`来读取文件。bufio默认提供四种分解方式：
   * ScanLines (default)
   * ScanWords
   * ScanRunes (highly useful for iterating over UTF-8 codepoints, as opposed to bytes)
   * ScanBytes
1. 写文件，基础方式可以利用`os.OpenFile`和`file.Write()`,或者利用ioutil包中一个非常有用的方法WriteFile()处理创建／打开文件
1. 缓存写`bufio.NewWriter()`,利用缓存来减少磁盘io的时间，提高效率。


关于如何实现io包的Reader和Writer，以及io包里的一些基本函数可以参考[Go中io包的使用方法](https://segmentfault.com/a/1190000015591319)  
io下面的ioutil子包对io做了一些封装，提供了一些便利功能，但是有一次在处理拷贝一个文件的时候，使用io.ReadAll()出现了OOM的问题,根本原因是我在容器中分配的内存有限，只有100M，而当时要拷贝的文件超过200M，使用io.ReadAll方法直接一次性将内存撑爆了，这篇文章介绍了在遇到大文件时候io.ReadAll()非常容易暴雷，建议使用io.Copy来替代,io.Copy固定使用32KB内存来拷贝文件直到遇到io.EOF。[Be careful with ioutil.ReadAll in Golang](https://haisum.github.io/2017/09/11/golang-ioutil-readall/)。  
关于bufio、ioutil读文件的速度比较（性能测试）和影响因素分析可以看看[这篇文章](https://segmentfault.com/a/1190000011680507)

## 参考

[Go文件操作大全](https://colobu.com/2016/10/12/go-file-operations/)
[Reading files in Go](https://kgrz.io/reading-files-in-go-an-overview.html)
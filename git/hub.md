# hub

github原生提供的工具[hub](https://github.com/github/hub)，好早以前英国同事就推荐过，一直也没想着用，最近每次提交PR都需要折腾一番，先打开浏览器，然后到目标repo，再到branch上去开PR。这几个简单动作重复一两次不觉得麻烦，
多做几次发现挺浪费时间，于是又想起了这个工具。找了出来果断给安装上:`brew install hub`。   
简单执行`hub`就出来一堆help信息，找到了我要的命令`hub pull-request`,不过到目标工作目录输入`hub pull-request`后，遇到第一个坑，要求认证信息，输入用户密码后却得到401错误。不明所以。也不想纠缠为什么，于是上google找了个解决方法。在用户工作目录下的`.config`文件夹下建一个hub文件，输入如下内容,替换认证用户和token即可：
```
github.*.com:
- user: *****@cn.**.com
  oauth_token: **********
  protocol: https
```


_20190929_
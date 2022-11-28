# Kubernetes 集群证书失效

## 问题及修复
手里有个实验环境部署了一年多，最近重新启用想做一些测试验证，信心满满的准备先跑下`kubectl get pods`来看下环境是不是正常，
结果就看到了下面的错误：

```txt
Unable to connect to the server: x509: certificate has expired or is not yet valid.
```

哈？证书无效？失效？其实错误非常明显了，但是怎么解决却有点蒙圈。因为一开始是用kubeadm安装的cluster，并不记得有手动配置什么证书...
搜了一圈才发现，kubeadm安装的cluster里用的证书默认就是一年，这个环境超过一年了，证书自然也就过期了。当前集群的证书信息可以通过
`kubeadm certs check-expiration`来查看，有些老版本需要运行`kubeadm alpha certs check-expiration`。更新证书的命令
kubeadm也直接提供了，只需要运行`kubeadm certs renew all`即可，老版本是`kubeadm alpha certs renew all`。当然为了保险
起见，也可以先备份下现有的配置文件：`cp -p /etc/kubernetes/*.conf $HOME/k8s-old-certs`。

本地的kubeconfig文件因为证书的更新也会失效，需要将`/etc/kubernetes/admin.conf`拷贝至`$HOME/.kube`。如果没有其他cluster的配置信息，那么直接拷贝并重命名为config即可。

## 参考链接

- [Renewing Kubernetes cluster certificates](https://www.ibm.com/docs/en/fci/1.1.0?topic=kubernetes-renewing-cluster-certificates)
- [Certificate Management with kubeadm](https://kubernetes.io/docs/tasks/administer-cluster/kubeadm/kubeadm-certs/#automatic-certificate-renewal)

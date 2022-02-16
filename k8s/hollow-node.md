# K8s Hollow Node

在做K8s相关开发时，有时会需要通过Node的增删改的事件来测试代码功能，而对已有环境进行Node的真实增删改可能会对现有开发测试有影响。
通过模拟Node的增删改事件是一个更好的方式。在K8s中，为模拟大型集群进行性能测试提供了一个名为[Kubemark](https://github.com/kubernetes/kubernetes/tree/master/test/kubemark)的工具。本文就将利用Kubemark工具来模拟构建一个Hollow Node。

## 创建Hollow Node

在Kubemark中，模拟真实Node的过程是通过创建Hollow Node来实现。如Hollow这个前缀名所提示的，“Hollow”表示组件的实现/实例化中，所有“移动”部分都被模拟出来。Hollow Nodes 使用现有的fake Kubelet（称为 SimpleKubelet），它使用 pkg/kubelet/fake-docker-manager.go 模拟其运行时管理器，因此大多数处理逻辑都有模拟。下面介绍如何在一个现有集群里添加Hollow Node。

1. 创建 kubemark namespace

```bash
kubectl create ns kubemark
```

2. 创建secret

将集群的kubeconfig文件拷贝并命名为config
```bash
kubectl create secret generic kubeconfig --type=Opaque --namespace=kubemark --from-file=kubelet.kubeconfig=config --from-file=kubeproxy.kubeconfig=config
```

3. 创建Hollow Node

创建Hollow Node的yaml文件可以是[hollow-node_simplified_template.yaml](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-scalability/hollow-node_simplified_template.yaml),或者在`kubernetes/test/kubemark/resources/`里的[hollow-node_template.yaml](https://github.com/kubernetes/kubernetes/blob/master/test/kubemark/resources/hollow-node_template.yaml)。

`hollow-node_simplified_template.yaml`文件中`{{numreplicas}}`, `{{kubemark_image_registry}}`和`{{kubemark_image_tag}}`是必须替换的参数。`hollow-node_template.yaml`文件中的变量替换，可以参考脚本[startup.sh](https://github.com/kubernetes/kubernetes/blob/master/test/kubemark/iks/startup.sh#L227-L250)。

这里贴出一个可用的yaml文件内容供参考。
```yaml
apiVersion: v1
kind: ReplicationController
metadata:
  name: hollow-node
  namespace: kubemark
spec:
  replicas: 1
  selector:
      name: hollow-node
  template:
    metadata:
      labels:
        name: hollow-node
    spec:
      initContainers:
      - name: init-inotify-limit
        image: docker.io/busybox:latest
        command: ['sysctl', '-w', 'fs.inotify.max_user_instances=200']
        securityContext:
          privileged: true
      volumes:
      - name: kubeconfig-volume
        secret:
          secretName: kubeconfig
      - name: logs-volume
        hostPath:
          path: /var/log
      containers:
      - name: hollow-kubelet
        image: antrea/kubemark:v1.18.4
        ports:
        - containerPort: 4194
        - containerPort: 10250
        - containerPort: 10255
        env:
        - name: NODE_NAME
          valueFrom:
            fieldRef:
              fieldPath: metadata.name
        command:
        - /kubemark
        args:
        - --morph=kubelet
        - --name=$(NODE_NAME)
        - --kubeconfig=/kubeconfig/kubelet.kubeconfig
        - --alsologtostderr
        - --v=2
        volumeMounts:
        - name: kubeconfig-volume
          mountPath: /kubeconfig
          readOnly: true
        - name: logs-volume
          mountPath: /var/log
        resources:
          requests:
            cpu: 20m
            memory: 50M
        securityContext:
          privileged: true
      - name: hollow-proxy
        image: antrea/kubemark:v1.18.4
        env:
        - name: NODE_NAME
          valueFrom:
            fieldRef:
              fieldPath: metadata.name
        command:
        - /kubemark
        args:
        - --morph=proxy
        - --name=$(NODE_NAME)
        - --use-real-proxier=false
        - --kubeconfig=/kubeconfig/kubeproxy.kubeconfig
        - --alsologtostderr
        - --v=2
        volumeMounts:
        - name: kubeconfig-volume
          mountPath: /kubeconfig
          readOnly: true
        - name: logs-volume
          mountPath: /var/log
        resources:
          requests:
            cpu: 20m
            memory: 50M
      tolerations:
      - effect: NoExecute
        key: node.kubernetes.io/unreachable
        operator: Exists
      - effect: NoExecute
        key: node.kubernetes.io/not-ready
        operator: Exists
```

## 参考资料

* [Kubemark集群性能测试](https://supereagle.github.io/2017/03/09/kubemark/)
* [kubemark-guide](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-scalability/kubemark-guide.md)
* [kubemark-startup-guide](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-scalability/kubemark-setup-guide.md)
* [build-kubemark](https://antrea.io/docs/main/docs/maintainers/build-kubemark/)
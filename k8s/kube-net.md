# kube-net

A repo about useful kubernetes network related codes, tools, docs, blogs etc.

## kubectl

```bash
# get service cidr
kubectl cluster-info dump | grep -m 1 service-cluster-ip-range
# get cluster cidr
kubectl cluster-info dump | grep -m 1 cluster-cidr
# Get each node' pod subnet
kubectl get nodes -A -o jsonpath='{range .items[*]}{.spec.podCIDR}{"\n"}{end}'
# Get each pod's ip, use " | " as delimiter
kubectl get pods -A -o jsonpath='{range .items[*]}{.metadata.namespace}{"/"}{.metadata.name}{" | "}{.status.podIP}{"\n"}{end}'
```

## docker

```bash
# add eth1 to a pod when there is no ip command inside the pod.
id_of_pod=`docker ps | grep <pod_container_name> | awk '{print $1}'`
pid_of_pod=`docker inspect --format "{{ .State.Pid }}" $id_of_pod`
nsenter -t $pid_of_pod -n ip route add 172.16.2.0/24 dev eth1
```

## tools

* IP calculation tools in ubuntu
```shell
# install tools
sudo apt-get install sipcalc
sudo apt install netmask
# show CIDR information
sipcalc 192.168.1.0/24 -a
# show eth0's CIDR information
sipcalc eth0
# figuring out minimal sets of subnets for a particular IP range
netmask -c 10.32.0.0:10.255.255.255
```
* hex to IP
```shell
ip=0A000021;printf '%d.%d.%d.%d\n' $(echo $ip | sed 's/../0x& /g')
```
* IP to hex
```shell
# option 1
apt install syslinux-utils
gethostip -x 10.0.0.33
# option 2
printf '%02x%02x%02x%02x' $(echo $1 | awk -F. '{print $1" "$2" "$3" "$4}')
```

* check NIC speed: `ethtool eth0 | grep Speed` or `cat /sys/class/net/eth0/speed`(Mb/s)
* [netshoot](https://github.com/nicolaka/netshoot) includes a set of powerful tools
* [kubectl-debug](https://github.com/aylei/kubectl-debug) is an out-of-tree solution for troubleshooting running pods, which allows you to run a new container in running pods for debugging purpose 
* [vegeta](https://github.com/tsenart/vegeta) is HTTP load testing tool and library


## blogs

* [Debugging network stalls on Kubernetes](https://github.blog/2019-11-21-debugging-network-stalls-on-kubernetes/)
* [Kubernetes Networking Demystified: A Brief Guide](https://www.stackrox.com/post/2020/01/kubernetes-networking-demystified/)


# Grafana+Graphite+Statsd

## Grafana

因为项目要用Grafana来展示通过statsd收集的metrics数据，最近都在研究Grafana，为了满足内部需求，对官方的[Grafana镜像](https://hub.docker.com/r/grafana/grafana/)做了一些基本的配置改动，为了便于以后回顾和查看，把一些配置记录在这里。

Grafana相关的文件都放到了[这里](grafana-docker),这里记录下主要的改动：

1. 禁用登陆页面： 官方grafana镜像启动后是有登陆页面的，默认账号是admin:admin，因为我们是内部使用，不想有多余的登陆操作，于是禁用登陆页面，这个需要在graphite.ini文件里配置，路径是/etc/grafana/grafana.ini。具体的配置参考[这里](grafana-docker/grafana.ini)
2. 添加默认[dashboard供应商](grafana-docker/default-provider.yaml),更多的内容可以参考[官方文档](https://grafana.com/docs/administration/provisioning/#dashboards)
3. 添加默认启动时的dashboards：我这里只是拷贝了一个[kafka数据模版](grafana-docker/dashboards),同时在构建新镜像时拷贝进去[Dockerfile](grafana-docker/Dockerfile)，[官方站点](https://grafana.com/grafana/dashboards)上有很多已发布的或者社区的dashboards，可以去搜索满足个人需要的。
4. 安装plugin：Grafana官方镜像貌似没加任何plugin，像很基本的Pie图插件也没有. 要查看有那些plugin，可以在grafana容器里运行`grafana-cli plugins list-remote`,安装pie chart需要运行`grafana-cli plugins install grafana-piechart-panel`.默认plugin是安装在路径`/var/lib/grafana/plugins`.如果在grafana.ini中修改了默认加载plugin的路径，则需要在Dockerfile中将`/var/lib/grafana/plugins`下的所有文件拷贝到新路径。

## Graphite+statsd

Graphite+Statsd也是基于[官方的镜像](https://hub.docker.com/r/graphiteapp/graphite-statsd/)做了一些改动。改动后的基本结构和配置放到[graphite-statsd](graphite-statsd-docker),官方的镜像提供的配置信息比较全，我在本地只是因为端口冲突重新配置了statsd的tcp和udp监听端口。另外为了保证没有vulnerability的问题，加了一行`apk upgrade --no-cache`，保证系统始终会更新到最新。
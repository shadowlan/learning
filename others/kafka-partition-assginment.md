# kakfa 分区分配策略

不同版本kafka的分区策略有所不同，从官方java doc文档里看，最新版本里提供四种分区策略：

* [RangeAssignor](https://kafka.apache.org/24/javadoc/org/apache/kafka/clients/consumer/RangeAssignor.html)
* [RoundRobinAssignor](https://kafka.apache.org/24/javadoc/org/apache/kafka/clients/consumer/RoundRobinAssignor.html)
* [StickyAssignor](https://kafka.apache.org/24/javadoc/org/apache/kafka/clients/consumer/StickyAssignor.html)
* [CooperativeStickyAssignor](https://kafka.apache.org/24/javadoc/org/apache/kafka/clients/consumer/CooperativeStickyAssignor.html)

这四类分区策略都实现了接口[ConsumerPartitionAssignor](https://kafka.apache.org/24/javadoc/org/apache/kafka/clients/consumer/ConsumerPartitionAssignor.html),而在kafka 2.3里面这个结构还叫做PartitionAssignor，如果用户要实现自定义的Assignor，需要留意是在那个版本下，需要实现那个接口类型。

当用户创建一个新kafka消费者时，可以配置具体使用哪个分区策略，配置项是`partition.assignment.strategy`, 从[kafka官方文档](https://kafka.apache.org/documentation/#consumerconfigs)可以查到默认的分配策略是`Range`，而该类型是一个列表，说明可以指定多个分区策略，需要注意的是，在同一个消费组下的所有消费者都必须至少有一个共同的策略，如果一个消费者试图加入一个组，但是策略列表里没有和组内其他消费者一样的分区策略，那么会收到下面的异常：

```
org.apache.kafka.common.errors.InconsistentGroupProtocolException: The group member’s supported protocols are incompatible with those of existing members or first group member tried to join with empty protocol type or empty protocol list.
```

关于Range和RoundRobin类型的分配策略的具体分配逻辑介绍，可以参考文章[Kafka分区分配策略](https://www.okcode.net/article/45817),懒得搬过来了。

如果用户在订阅主题时，加入同一个组的消费者既有使用分配策略的，又有使用直接指定分区号来订阅消息的，则会出现冲突而导致收到如下异常：

```
java.lang.IllegalStateException: Subscription to topics, partitions and pattern are mutually exclusive
```
针对这个错误的扩展阅读如下：  
[Don't Use Apache Kafka Consumer Groups the Wrong Way!](https://paolopatierno.wordpress.com/2017/07/27/apache-kafka-consumer-groups-dont-use-them-in-the-wrong-way/)  
[IllegalStateException](https://blog.csdn.net/Dongguabai/article/details/86543698)

## 参考

[Partition Assignment Strategies](https://medium.com/streamthoughts/understanding-kafka-partition-assignment-strategies-and-how-to-write-your-own-custom-assignor-ebeda1fc06f3)
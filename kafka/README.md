

## kafka
[架构图](http://topgoer.com/static/9.2/1.jpg)

- Partition

Topic的分区，每个topic可以有多个分区，分区的作用是做负载，提高kafka的吞 吐量。同一个topic在不同的分区的数据是不重复的，partition的表现形式就是一个一个的⽂件夹！

Partition在服务器上的表现形式就是⼀个⼀个的⽂件夹，每个partition的⽂件夹下⾯会有多组segment ⽂件，每组segment⽂件⼜包含 .index ⽂件、 .log ⽂件、 .timeindex ⽂件三个⽂件，其中 .log ⽂ 件就是实际存储message的地⽅，⽽ .index 和 .timeindex ⽂件为索引⽂件，⽤于检索消息


- Replication

每一个分区都有多个副本，副本的作用是做备胎。当主分区（Leader）故障的 时候会选择一个备胎（Follower）上位，成为Leader。在kafka中默认副本的最大数量是10 个，且副本的数量不能大于Broker的数量，follower和leader绝对是在不同的机器，同一机 器对同一个分区也只可能存放一个副本（包括自己）。

- 保留期限

Kafka可以配置⼀个保留期限，⽤来标识⽇志会在Kafka集群内保留多⻓时间。Kafka集群会保留在保留 期限内所有被发布的消息，不管这些消息是否被消费过。⽐如保留期限设置为两天，那么数据被发布到 Kafka集群的两天以内，所有的这些数据都可以被消费。当超过两天，这些数据将会被清空，以便为后 续的数据腾出空间。由于Kafka会将数据进⾏持久化存储（即写⼊到硬盘上），所以保留的数据⼤⼩可 以设置为⼀个⽐较⼤的值。


## 发送端流程

```
1.⽣产者从Kafka集群获取分区leader信息
2.⽣产者将消息发送给leader
3.leader将消息写入本地磁盘
4.follower从leader拉取消息数据
5.follower将消息写入本地磁盘后向leader发送ACK
6.leader收到所有的follower的ACK之后向生产者发送ACK

1.partition在写入的时候可以指定需要写入的partition，如果有指定，则写入对应的partition。
2.如果没有指定partition，但是设置了数据的key，则会根据key的值hash出一个partition。
3.如果既没指定partition，又没有设置key，则会采用轮询⽅式，即每次取一小段时间的数据写入某个partition，下一小段的时间写入下一个partition
```

- ACK应答机制
producer在向kafka写入消息的时候，可以设置参数来确定是否确认kafka接收到数据，这个参数可设置 的值为 0,1,all
```
0代表producer往集群发送数据不需要等到集群的返回，不确保消息发送成功。安全性最低但是效率最高。
1代表producer往集群发送数据只要leader应答就可以发送下一条，只确保leader发送成功。
all代表producer往集群发送数据需要所有的follower都完成从leader的同步才会发送下一条，确保 leader发送成功和所有的副本都完成备份。安全性最⾼高，但是效率最低。
```


## 接收端流程

在同⼀个消费者组中，每个消费者实例可以消费多个分区，但是每个分区最多只 能被消费者组中的⼀个实例消费。其实这也很好理解，如果允许两个消费者实例同 时消费同⼀个分区，那么就⽆法记录这个分区被这个消费者组消费的offset了。如果在消费者组中动态 的上线或下线消费者，那么Kafka集群会⾃动调整分区与消费者实例间的对应关系。


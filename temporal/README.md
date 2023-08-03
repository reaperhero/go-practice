### temporal原理

- Start：工作流的创建者/发起者。可以将不同的Activity(也就是Worker实现的具体的执行逻辑模块，每一个Activity都有一个名字)组织成一个Workflow，并且开启workflow。
- Temporal Server：存储所有工作流的数据、状态的中间件，整个工作依赖于该server（后续简写为TS）
- Worker：实际进行逻辑处理的执行者。该模块实现具体的执行逻辑，并且在启动的时候注册到TS。
- Bank：官方给的示例，可以理解为DB，TS的存储模块。

具体的流程描述：
```
1.启动Temporal Server。
2.启动Temporal Client，也就是Worker，监听TS，循环获取待执行的工作流。
3.Start创建一个工作流，封装参数，调用sdk的api(rpc)发送到TS。Worker拉取到工作流开始逻辑处理
```

worker是Activity和Workflow的包装，worker的唯一工作就是执行Activity和Workflow并将结果返回给TS。
一个Workflow包含多个Activity，对Activity进行编排，多个Activity可以并行，也可以同步（阻塞到都某个Activity执行完毕）。其底层会阻塞到Future.Get()方法上。
Activity一定是在worker里面的，因为逻辑的具体执行者一定是worker。workflow不一定要封装在worker里，可以单独拿出来，实现Activity的编排，worker只管一个一个Activity的具体实现即可。

### Steps to run this sample:
1) Run a [Temporal Server](https://github.com/temporalio/samples-go/tree/main/#how-to-use).
2) Run the following command to start the worker
```
> go run worker.go
2023/08/03 17:38:12 INFO  No logger configured for temporal client. Created default one.
2023/08/03 17:38:12 INFO  Started Worker Namespace default TaskQueue hello-world WorkerID 12661@chenqiangjundeMacBook-Pro-2.local@
2023/08/03 17:38:18 INFO  Activity Namespace default TaskQueue hello-world WorkerID 12661@chenqiangjundeMacBook-Pro-2.local@ ActivityID 5 ActivityType Activity Attempt 4 WorkflowType Workflow WorkflowID hello_world_workflowID RunID 5e5d5e9a-9394-445c-a4a1-aa171e02a5b4 name Temporal
2023/08/03 17:38:18 INFO  Activity complete after timeout. Namespace default TaskQueue hello-world WorkerID 12661@chenqiangjundeMacBook-Pro-2.local@ WorkflowID hello_world_workflowID RunID 5e5d5e9a-9394-445c-a4a1-aa171e02a5b4 ActivityType Activity Attempt 4 Result &Payloads{Payloads:[]*Payload{&Payload{Metadata:map[string][]byte{encoding: [106 115 111 110 47 112 108 97 105 110],},Data:[34 72 101 108 108 111 32 84 101 109 112 111 114 97 108 33 34],},},} Error <nil>

```
3) Run the following command to start the example
```
> go run starter.go
2023/08/03 17:38:47 INFO  No logger configured for temporal client. Created default one.
2023/08/03 17:38:47 Started workflow WorkflowID hello_world_workflowID RunID 5e5d5e9a-9394-445c-a4a1-aa171e02a5b4

```

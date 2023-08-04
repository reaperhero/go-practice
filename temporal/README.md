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
2023/08/04 18:04:46 INFO  No logger configured for temporal client. Created default one.
2023/08/04 18:04:46 INFO  Started Worker Namespace default TaskQueue TRANSFER_MONEY_TASK_QUEUE WorkerID 1476@node13@
2023/08/04 18:05:16 Withdrawing $250 from account 85-150.

2023/08/04 18:05:16 DEBUG ExecuteActivity Namespace default TaskQueue TRANSFER_MONEY_TASK_QUEUE WorkerID 1476@node13@ WorkflowType MoneyTransfer WorkflowID pay-invoice-701 RunID b0b950aa-b8c7-4784-8fbc-6b383a80cd42 Attempt 1 ActivityID 11 ActivityType Deposit
2023/08/04 18:05:16 Depositing $250 into account 43-812.
```
3) Run the following command to start the example
```
> go run starter.go
2023/08/04 18:04:54 INFO  No logger configured for temporal client. Created default one.
2023/08/04 18:04:54 Starting transfer from account 85-150 to account 43-812 for 250
2023/08/04 18:04:54 WorkflowID: pay-invoice-701 RunID: b0b950aa-b8c7-4784-8fbc-6b383a80cd42
2023/08/04 18:05:16 Transfer complete (transaction IDs: W1779185060, D4129841576)
```

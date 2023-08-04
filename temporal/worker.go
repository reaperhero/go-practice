package main

import (
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"log"
	"temporlservice/app"
)

func main() {
	// Create the client object just once per process
	c, err := client.Dial(client.Options{
		HostPort: "172.16.94.221:7233",
	})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	w := worker.New(c, app.MoneyTransferTaskQueueName, worker.Options{})

	// This worker hosts both Workflow and Activity functions.
	w.RegisterWorkflow(app.MoneyTransfer)
	w.RegisterActivity(app.Withdraw)
	w.RegisterActivity(app.Deposit)
	w.RegisterActivity(app.Refund)

	// Start listening to the Task Queue.
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}

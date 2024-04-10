package main

import (
	"log"

	"github.com/nidhey27/cqrs-go/activities"
	"github.com/nidhey27/cqrs-go/workflows"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	w := worker.New(c, "account", worker.Options{})
	w.RegisterWorkflow(workflows.CalculateBalance)
	w.RegisterActivity(activities.GetTransactions)
	w.RegisterActivity(activities.GetBalance)
	w.RegisterActivity(activities.SetBalance)

	// Start listening to the Task Queue
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}

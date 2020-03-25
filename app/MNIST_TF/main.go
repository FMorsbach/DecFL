package main

import (
	"github.com/FMorsbach/DecFL/control"
	"github.com/FMorsbach/dlog"
)

// TODO: make CL argument
const CLIENTS = 3
const TARGET_ACC = 0.90

func init() {

	// TODO: Check if chain reachable

	// TODO: Check if storage reachable

	dlog.SetDebug(false)

}

func main() {

	dlog.Println("Starting MNIST-TF scenario")

	_, err := control.Initialize()
	if err != nil {
		dlog.Fatal(err)
	}

	dlog.Println("Created and deployed model")

	iteration := 0

	status, err := control.Status()
	dlog.Printf("Initial Status %s\n", status)

	dlog.Printf("Starting sequential taining with %d local clients untill %f accuracy reached", CLIENTS, TARGET_ACC)
	for status.Accuracy < TARGET_ACC {
		for i := 0; i < CLIENTS; i++ {
			err = control.Iterate()
			if err != nil {
				dlog.Fatal(err)
			}
			dlog.Printf("Client %d finished and submitted training", i)
		}

		iteration++
		dlog.Printf("Clients finished iteration %d, starting aggregation", iteration)
		err = control.Aggregate()
		if err != nil {
			dlog.Fatal(err)
		}

		status, err = control.Status()
		dlog.Printf("Status after iteration %d: %s\n", iteration, status)
	}

	dlog.Printf("Finished training after %d iterations and reached %s", iteration, status)
}

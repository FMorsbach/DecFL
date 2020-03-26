package main

import (
	"github.com/FMorsbach/DecFL/communcation/chain"
	"github.com/FMorsbach/DecFL/communcation/storage"
	"github.com/FMorsbach/DecFL/control"
	"github.com/FMorsbach/DecFL/training/tensorflow"
	"github.com/FMorsbach/dlog"
)

// TODO: make CL argument
const CLIENTS = 3
const TARGET_ACC = 0.90

func init() {

	// TODO: Check if chain reachable

	// TODO: Check if storage reachable

	dlog.SetDebug(false)
	dlog.SetPrefix("Main: ")

	chain.EnableDebug(false)
	storage.EnableDebug(false)
	tensorflow.EnableDebug(false)
	control.EnableDebug(true)
}

func main() {

	dlog.Debugln("Starting MNIST-TF scenario")

	_, err := control.Initialize()
	if err != nil {
		dlog.Fatal(err)
	}

	dlog.Debugln("Created and deployed model")

	iteration := 0

	status, err := control.Status()
	dlog.Debugf("Initial Status %s\n", status)

	dlog.Debugf("Starting sequential taining with %d local clients untill %f accuracy reached", CLIENTS, TARGET_ACC)
	for status.Accuracy < TARGET_ACC {
		for i := 0; i < CLIENTS; i++ {
			err = control.Iterate()
			if err != nil {
				dlog.Fatal(err)
			}
			dlog.Debugf("Client %d finished and submitted training", i)
		}

		iteration++
		dlog.Debugf("Clients finished iteration %d, starting aggregation", iteration)
		err = control.Aggregate()
		if err != nil {
			dlog.Fatal(err)
		}

		status, err = control.Status()
		dlog.Debugf("Status after iteration %d: %s\n", iteration, status)
	}

	dlog.Debugf("Finished training after %d iterations and reached %s", iteration, status)
}

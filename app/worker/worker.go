package main

import (
	"log"
	"os"
	"time"

	md "github.com/FMorsbach/DecFL/model"
	"github.com/FMorsbach/DecFL/model/chain/ethereum"
	"github.com/FMorsbach/DecFL/model/common"
	"github.com/FMorsbach/DecFL/model/storage"
	"github.com/FMorsbach/DecFL/model/training/tensorflow"
	"github.com/FMorsbach/dlog"
)

var model md.Model

var logB = log.New(os.Stdout, "BENCH: ", log.LstdFlags|log.Lmsgprefix)
var start time.Time

func init() {

	dlog.SetDebug(true)
	dlog.SetPrefix("Worker: ")

	dlog.Debug("Init...")

	md.EnableDebug(true)
	storage.EnableDebug(true)
	ethereum.EnableDebug(true)

	chain, store, modelID, err := md.ParseCLIConfig()
	if err != nil {
		dlog.Fatalf("Error parsing CLI config: %s", err)
	}

	// setup trainer
	trainer := tensorflow.NewTensorflowMLF()
	dlog.Debug("Created trainer")

	model, err = md.NewModel(chain, store, trainer, modelID)
	if err != nil {
		dlog.Fatalf("Error creating model: %s", err)
	}

	//dlog.Printf("Working on model %s as node %s connected to %s\n", modelID, nodeID, chainConnection)
}

func main() {

	dlog.Print("Start working...")
	trainings := 0
	aggregations := 0

	for !isFinished() {

		state, err := model.State()
		if err != nil {
			dlog.Fatal(err)
		}

		switch state {
		case common.Training:

			start = time.Now()
			err = model.Iterate()
			if err != nil {
				if err.Error() == "VM Exception while processing transaction: revert Not valid at this state" {
					continue
				}
				dlog.Fatalf("Error training: %s", err)
			}
			logB.Printf("TRAINING_ROUND %d %.3f\n", trainings, time.Since(start).Seconds())
			trainings++

			start = time.Now()
			waitForStateTransitionFrom(state)
			logB.Printf("WAITING_AFTER_TRAINING %d %.3f\n", trainings-1, time.Since(start).Seconds())

		case common.Aggregation:
			start = time.Now()
			err = model.Aggregate()
			if err != nil {
				if err.Error() == "VM Exception while processing transaction: revert Not valid at this state" {
					continue
				}
				dlog.Fatalf("Error aggregating: %s", err)
			}
			logB.Printf("AGGREGATION_ROUND %d %.3f\n", aggregations, time.Since(start).Seconds())
			aggregations++

			start = time.Now()
			waitForStateTransitionFrom(state)
			logB.Printf("WAITING_AFTER_AGGREGATION %d %.3f\n", aggregations-1, time.Since(start).Seconds())
		}
	}

	dlog.Printf("Finished working, Trainings: %d, Aggregations: %d\n", trainings, aggregations)
}

func isFinished() bool {
	state, err := model.State()
	if err != nil {
		dlog.Fatal(err)
	}
	return (state == common.Finished)
}

func waitForStateTransitionFrom(currentState uint8) {

	state, err := model.State()
	if err != nil {
		dlog.Fatal(err)
	}

	for state == currentState {

		time.Sleep(time.Second)

		state, err = model.State()
		if err != nil {
			dlog.Fatal(err)
		}
	}

	return
}

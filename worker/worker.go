package main

import (
	md "github.com/FMorsbach/DecFL/model"
	"github.com/FMorsbach/DecFL/model/storage"
	"github.com/FMorsbach/DecFL/model/training/tensorflow"
	"github.com/FMorsbach/dlog"
)

var model md.Model

func init() {

	dlog.SetDebug(true)
	md.EnableDebug(true)
	storage.EnableDebug(true)

	chain, store, modelID, err := md.ParseCLIConfig()
	if err != nil {
		dlog.Fatal(err)
	}

	// setup trainer
	trainer := tensorflow.NewTensorflowMLF()
	dlog.Debug("Created trainer")

	model, err = md.NewControl(chain, store, trainer, modelID)
	if err != nil {
		dlog.Fatal(err)
	}

	//dlog.Printf("Working on model %s as node %s connected to %s\n", modelID, nodeID, chainConnection)
}

func main() {

	dlog.Print("Starting training...")
	for true {

		model.WaitForNewEpoch()

		model.Iterate()

		model.WaitForAggregation()

		model.Aggregate()

		break
	}
}

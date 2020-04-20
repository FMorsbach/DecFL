package main

import (
	md "github.com/FMorsbach/DecFL/model"
	"github.com/FMorsbach/DecFL/model/training/tensorflow"
	"github.com/FMorsbach/dlog"
)

var model md.Model

func init() {

	chain, store, modelID, err := md.ParseCLIConfig()
	if err != nil {
		dlog.Fatal(err)
	}

	// setup trainer
	trainer := tensorflow.NewTensorflowMLF()

	model, err = md.NewControl(chain, store, trainer, modelID)
	if err != nil {
		dlog.Fatal(err)
	}

	//dlog.Printf("Working on model %s as node %s connected to %s\n", modelID, nodeID, chainConnection)
	md.EnableDebug(false)
}

func main() {

	for true {

		model.WaitForNewEpoch()

		model.Iterate()

		model.WaitForAggregation()

		model.Aggregate()
	}
}

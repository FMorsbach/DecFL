package main

import (
	"flag"
	"time"

	"github.com/FMorsbach/DecFL/communication"
	"github.com/FMorsbach/DecFL/communication/mocks"
	"github.com/FMorsbach/DecFL/control"
	"github.com/FMorsbach/dlog"
)

var modelID communication.ModelIdentifier
var trainerID communication.TrainerIdentifier

func init() {

	modelPtr := flag.String("model", "", "The model the worker should work on")
	trainerPtr := flag.String("trainer", "", "The id of this worker node")

	flag.Parse()

	modelID = communication.ModelIdentifier(*modelPtr)
	trainerID = communication.TrainerIdentifier(*trainerPtr)

	if modelID == "" {
		dlog.Fatal("No model id provided")
	}
	if trainerID == "" {
		dlog.Fatal("No trainer id provided")
	}

	control.EnableDebug(true)
}

func main() {

	chain := mocks.NewRedis()

	localEpoch, err := chain.ModelEpoch(modelID)
	if err != nil {
		panic(err)
	}

	for true {

		if globalEpoch, err := chain.ModelEpoch(modelID); err != nil {
			panic(err)
		} else if localEpoch > globalEpoch {
			time.Sleep(time.Second)
			continue
		} else {
			err := control.Iterate(modelID, trainerID)
			if err != nil {
				dlog.Println(err)
			}
			localEpoch = globalEpoch + 1
		}
	}
}

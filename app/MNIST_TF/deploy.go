package main

import (
	"flag"

	"github.com/FMorsbach/DecFL/communication"
	"github.com/FMorsbach/DecFL/communication/chain"
	"github.com/FMorsbach/DecFL/communication/storage"
	"github.com/FMorsbach/DecFL/control"
	"github.com/FMorsbach/DecFL/training/tensorflow"
	"github.com/FMorsbach/dlog"
)

func init() {

	// TODO: Check if chain reachable

	// TODO: Check if storage reachable

	dlog.SetDebug(true)
	dlog.SetPrefix("Main: ")

	chain.EnableDebug(false)
	storage.EnableDebug(false)
	tensorflow.EnableDebug(false)
	control.EnableDebug(false)
}

func main() {

	var command string
	var argument string

	flag.StringVar(&command, "cmd", "", "What do you want to do? deploy / aggregate")
	flag.StringVar(&argument, "argument", "", "What argument? model type / model id")
	flag.Parse()

	switch command {
	case "deploy":

		modelID, err := control.Initialize()
		if err != nil {
			dlog.Fatal(err)
		}

		dlog.Printf("Deployed MNIST model with id %s to redis \n", modelID)

	case "aggregate":

		if argument == "" {
			dlog.Fatal("You need to specify a model id")
		}

		err := control.Aggregate(communication.ModelIdentifier(argument))
		if err != nil {
			dlog.Fatal(err)
		}
		dlog.Println("Aggregated model and upated global weights")

	default:
		dlog.Fatal("No valid command provided")
	}

}

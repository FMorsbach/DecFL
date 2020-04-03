package main

import (
	"flag"

	"github.com/FMorsbach/DecFL/communication"
	"github.com/FMorsbach/DecFL/communication/mocks"
	"github.com/FMorsbach/DecFL/control"
	"github.com/FMorsbach/DecFL/models/MNIST"
	"github.com/FMorsbach/DecFL/training/tensorflow"
	"github.com/FMorsbach/dlog"
)

var ctl control.Control

func init() {

	dlog.SetDebug(true)
	dlog.SetPrefix("Main: ")

	redis := mocks.NewRedis("localhost:6379")

	trainer := tensorflow.NewTensorflowTrainer()

	ctl = control.NewControl(redis, redis, trainer)
}

func main() {

	var command string
	var argument string

	flag.StringVar(&command, "cmd", "", "What do you want to do? deploy / aggregate")
	flag.StringVar(&argument, "argument", "", "What argument? model type / model id")
	flag.Parse()

	switch command {
	case "deploy":

		config, weights := MNIST.GenerateInitialModel()
		modelID, err := ctl.Initialize(config, weights)
		if err != nil {
			dlog.Fatal(err)
		}

		dlog.Printf("Deployed MNIST model with id %s to redis \n", modelID)

	case "aggregate":

		if argument == "" {
			dlog.Fatal("You need to specify a model id")
		}

		err := ctl.Aggregate(communication.ModelIdentifier(argument))
		if err != nil {
			dlog.Fatal(err)
		}
		dlog.Println("Aggregated model and upated global weights")

		status, err := ctl.Status(communication.ModelIdentifier(argument))
		if err != nil {
			dlog.Fatal(err)
		}
		dlog.Println(status)

	default:
		dlog.Fatal("No valid command provided")
	}

}

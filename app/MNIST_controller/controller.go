package main

import (
	"flag"

	"github.com/FMorsbach/DecFL/app/MNIST"
	"github.com/FMorsbach/DecFL/model"
	"github.com/FMorsbach/DecFL/model/common"
	"github.com/FMorsbach/DecFL/model/mocks"
	"github.com/FMorsbach/DecFL/model/training/tensorflow"
	"github.com/FMorsbach/dlog"
)

var ctl model.Model

func init() {

	dlog.SetDebug(true)
	dlog.SetPrefix("Main: ")

	redis := mocks.NewRedis("localhost:6379")
	if ok, err := redis.IsReachable(); !ok {
		dlog.Fatal(err)
	}

	trainer := tensorflow.NewTensorflowMLF()

	ctl = model.NewControl(redis, redis, trainer)
}

func main() {

	var command string
	var argument string

	flag.StringVar(&command, "cmd", "", "What do you want to do? deploy / aggregate")
	flag.StringVar(&argument, "arg", "", "What argument? model type / model id")
	flag.Parse()

	config, weights := MNIST.GenerateInitialModel()
	modelID, err := ctl.Initialize(config, weights, common.Hyperparameters{UpdatesTillAggregation: 3})
	if err != nil {
		dlog.Fatal(err)
	}

	dlog.Printf("Deployed MNIST model with id %s to redis \n", modelID)

}

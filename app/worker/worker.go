package main

import (
	"flag"
	"time"

	"github.com/FMorsbach/DecFL/communication"
	bc "github.com/FMorsbach/DecFL/communication/chain"
	"github.com/FMorsbach/DecFL/communication/mocks"
	"github.com/FMorsbach/DecFL/communication/storage"
	"github.com/FMorsbach/DecFL/control"
	"github.com/FMorsbach/DecFL/training/tensorflow"
	"github.com/FMorsbach/dlog"
)

var modelID communication.ModelIdentifier
var trainerID communication.TrainerIdentifier

var chain bc.Chain
var store storage.Storage
var ctl control.Control

func init() {

	modelPtr := flag.String("model", "", "The model the worker should work on")
	trainerPtr := flag.String("trainer", "", "The id of this worker node")
	redisConnectionPtr := flag.String("redis", "", "The connection identifier to redis")

	flag.Parse()

	modelID = communication.ModelIdentifier(*modelPtr)
	trainerID = communication.TrainerIdentifier(*trainerPtr)
	redisConnection := *redisConnectionPtr

	if modelID == "" {
		dlog.Fatal("No model id provided")
	}
	if trainerID == "" {
		dlog.Fatal("No trainer id provided")
	}
	if redisConnection == "" {
		dlog.Fatal("No redis connection provided")
	}

	redis := mocks.NewRedis(redisConnection)

	if ok, err := redis.IsReachable(); !ok {
		dlog.Fatal("Cant reach redis", err)
	}

	chain = redis
	store = redis
	trainer := tensorflow.NewTensorflowTrainer()
	ctl = control.NewControl(chain, store, trainer)

	control.EnableDebug(true)
}

func main() {

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
			err := ctl.Iterate(modelID, trainerID)
			if err != nil {
				dlog.Println(err)
			}
			localEpoch = globalEpoch + 1
		}
	}
}

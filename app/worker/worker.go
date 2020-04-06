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
var nodeID communication.TrainerIdentifier

var chain bc.Chain
var store storage.Storage
var ctl control.Control

func init() {

	modelPtr := flag.String("model", "", "The model the worker should work on")
	nodeIDPtr := flag.String("trainer", "", "The id of this worker node")
	redisConnectionPtr := flag.String("redis", "", "The connection identifier to redis")

	flag.Parse()

	modelID = communication.ModelIdentifier(*modelPtr)
	nodeID = communication.TrainerIdentifier(*nodeIDPtr)
	redisConnection := *redisConnectionPtr

	if modelID == "" {
		dlog.Fatal("No model id provided")
	}
	if nodeID == "" {
		dlog.Fatal("No node id provided")
	}
	if redisConnection == "" {
		dlog.Fatal("No redis connection provided")
	}

	redis := mocks.NewRedis(redisConnection)

	if ok, err := redis.IsReachable(); !ok {
		dlog.Fatal("Cant reach redis: ", err)
	}

	dlog.Printf("Working on model %s as node %s connected to %s\n", modelID, nodeID, redisConnection)

	chain = redis
	store = redis
	trainer := tensorflow.NewTensorflowMLF()
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
			err := ctl.Iterate(modelID, nodeID)
			if err != nil {
				dlog.Println(err)
			}
			localEpoch = globalEpoch + 1
		}
	}
}

package main

import (
	"flag"
	"time"

	md "github.com/FMorsbach/DecFL/model"
	ch "github.com/FMorsbach/DecFL/model/chain"
	"github.com/FMorsbach/DecFL/model/common"
	"github.com/FMorsbach/DecFL/model/mocks"
	st "github.com/FMorsbach/DecFL/model/storage"
	"github.com/FMorsbach/DecFL/model/training/tensorflow"
	"github.com/FMorsbach/dlog"
)

var modelID common.ModelIdentifier

var chain ch.Chain
var store st.Storage
var ctl md.Model

func init() {

	var chainType string
	var storageType string
	var chainConnection string
	var storageConnection string

	modelPtr := flag.String("model", "", "The model the worker should work on")
	keyPtr := flag.String("key", "", "The private key of this worker node")
	_ = keyPtr

	flag.StringVar(&chainType, "chainType", "ethereum", "Wheter to choose ethereum or redis as chain")
	flag.StringVar(&storageType, "storageType", "redis", "Wheter to choose IPFS or redis as storage")
	flag.StringVar(&chainConnection, "chain", "", "The connection identifier to the chain")
	flag.StringVar(&storageConnection, "storage", "", "The connection identifier to the storage service")

	flag.Parse()

	modelID = common.ModelIdentifier(*modelPtr)
	//	nodeID = common.TrainerIdentifier()

	if modelID == "" {
		dlog.Fatal("No model id provided")
	}
	//	if nodeID == "" {
	//		dlog.Fatal("No node id provided")
	//	}
	if chainConnection == "" {
		dlog.Fatal("No redis connection provided")
	}

	redis := mocks.NewRedis(chainConnection)

	if ok, err := redis.IsReachable(); !ok {
		dlog.Fatal("Cant reach redis: ", err)
	}

	//dlog.Printf("Working on model %s as node %s connected to %s\n", modelID, nodeID, chainConnection)

	chain = redis
	store = redis
	trainer := tensorflow.NewTensorflowMLF()
	ctl = md.NewControl(chain, store, trainer)

	md.EnableDebug(true)
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
			//err := ctl.Iterate(modelID, nodeID)
			if err != nil {
				dlog.Println(err)
			}
			localEpoch = globalEpoch + 1
		}
	}
}

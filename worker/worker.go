package main

import (
	"flag"

	md "github.com/FMorsbach/DecFL/model"
	ch "github.com/FMorsbach/DecFL/model/chain"
	"github.com/FMorsbach/DecFL/model/chain/ethereum"
	"github.com/FMorsbach/DecFL/model/common"
	"github.com/FMorsbach/DecFL/model/mocks"
	st "github.com/FMorsbach/DecFL/model/storage"
	"github.com/FMorsbach/DecFL/model/training/tensorflow"
	"github.com/FMorsbach/dlog"
)

var model md.Model

func init() {

	var chainType string
	var storageType string
	var chainConnection string
	var storageConnection string
	var privateKey string
	var modelAddress string
	var redisPassword string

	flag.StringVar(&chainType, "chainType", "ethereum", "Wheter to choose ethereum or redis as chain")
	flag.StringVar(&storageType, "storageType", "redis", "Wheter to choose IPFS or redis as storage")
	flag.StringVar(&chainConnection, "c", "", "The connection identifier to the chain")
	flag.StringVar(&storageConnection, "s", "", "The connection identifier to the storage service")
	flag.StringVar(&privateKey, "k", "", "the private key for the chain")
	flag.StringVar(&modelAddress, "m", "", "The address of the model")
	flag.StringVar(&redisPassword, "redisPW", "", "(Optional) The password for redis if used")

	required := []string{"c", "s", "k", "m"}
	flag.Parse()

	seen := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) { seen[f.Name] = true })
	for _, req := range required {
		if !seen[req] {
			dlog.Fatalf("missing required -%s argument/flag\n", req)
		}
	}

	// setup chain
	var chain ch.Chain
	switch chainType {
	case "ethereum":
		var err error
		chain, err = ethereum.NewEthereum(chainConnection, privateKey)
		if err != nil {
			dlog.Fatal(err)
		}

	case "redis":
		redis := mocks.NewRedis(chainConnection, redisPassword)
		if ok, err := redis.IsReachable(); !ok {
			dlog.Fatal("Cant reach redis: ", err)
		}
		chain = redis

	default:
		dlog.Fatal("Invalid chain type")
	}

	// setup storage
	var store st.Storage
	switch storageType {
	case "redis":
		redis := mocks.NewRedis(storageConnection, redisPassword)
		if ok, err := redis.IsReachable(); !ok {
			dlog.Fatal("Cant reach redis: ", err)
		}
		store = redis

	default:
		dlog.Fatal("Invalid storage type")
	}

	// setup trainer
	trainer := tensorflow.NewTensorflowMLF()

	// setup model
	modelID := common.ModelIdentifier(modelAddress)

	var err error
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

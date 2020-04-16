package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/FMorsbach/DecFL/model"
	"github.com/FMorsbach/DecFL/model/chain/ethereum"
	"github.com/FMorsbach/DecFL/model/common"
	"github.com/FMorsbach/DecFL/model/mocks"
	"github.com/FMorsbach/DecFL/model/storage"
	"github.com/FMorsbach/dlog"
)

func init() {
	dlog.SetPrefix("MNIST: ")
	dlog.SetDebug(false)
	dlog.SetFlags(dlog.Flags() | log.Lshortfile)
}

func main() {

	config, weights := GenerateInitialModel()

	var st storage.Storage
	redis := mocks.NewRedis("localhost:6379", "")
	if ok, err := redis.IsReachable(); !ok {
		dlog.Fatal(err)
	}
	st = redis

	ch, err := ethereum.NewEthereum("http://localhost:8545", "3b3a098805d048bab52b82b8767da2117af104cc97ec820acbe1b63e768ebba7")
	if err != nil {
		dlog.Fatal(err)
	}

	modelID, err := model.Deploy(
		config,
		weights,
		st,
		ch,
		common.Hyperparameters{UpdatesTillAggregation: 3},
	)
	if err != nil {
		dlog.Fatal(err)
	}

	fmt.Println(string(modelID))
}

func GenerateInitialModel() (configuration string, weights string) {

	configuration = loadDataFromDisk("configuration.txt")
	weights = loadDataFromDisk("weights.txt")
	return
}

func loadDataFromDisk(file string) (data string) {
	var content []byte

	content, err := ioutil.ReadFile(file)
	if err != nil {
		dlog.Fatal(err)
	}
	dlog.Debugf("Read %d bytes from %s", len(content), file)

	data = string(content)
	return
}

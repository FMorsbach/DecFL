package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/FMorsbach/DecFL/model"
	"github.com/FMorsbach/DecFL/model/common"
	"github.com/FMorsbach/dlog"
	ethCommon "github.com/ethereum/go-ethereum/common"
)

func init() {
	dlog.SetPrefix("MNIST: ")
	dlog.SetDebug(false)
	dlog.SetFlags(dlog.Flags() | log.Lshortfile)
}

func main() {

	config, weights := generateInitialModel()

	chain, store, _, err := model.ParseCLIConfig()
	if err != nil {
		dlog.Fatal(err)
	}

	modelID, err := model.Deploy(
		config,
		weights,
		store,
		chain,
		common.Hyperparameters{
			UpdatesTillAggregation: 3,
			Epochs:                 3,
		},
	)
	if err != nil {
		dlog.Fatal(err)
	}

	// Read trainer accounts and add to contract
	accounts := strings.Split(loadDataFromDisk("accounts.txt"), "\n")
	for _, acc := range accounts {
		address := ethCommon.HexToAddress(acc)
		err := chain.AddTrainer(modelID, common.TrainerIdentifier(address))
		if err != nil {
			dlog.Fatal(err)
		}
	}

	fmt.Println(string(modelID))
}

func generateInitialModel() (configuration string, weights string) {

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

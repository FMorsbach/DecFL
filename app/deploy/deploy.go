package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/FMorsbach/DecFL/model"
	"github.com/FMorsbach/DecFL/model/common"
	"github.com/FMorsbach/dlog"
	ethCommon "github.com/ethereum/go-ethereum/common"
)

var logger = dlog.New(os.Stderr, "Deploy: ", log.LstdFlags|log.Lshortfile, false)

func main() {

	var chainConnection string
	var storageConnection string
	var privateKey string
	var storageType string
	var scriptsArchive string
	var configFile string
	var weightsFile string
	var trainerFile string

	required := []string{"chain", "storage", "key", "storageType", "scripts", "config", "weights", "trainer"}
	flag.StringVar(&chainConnection, "chain", "", "The connection identifier to the chain")
	flag.StringVar(&storageConnection, "storage", "", "The connection identifier to the storage service")
	flag.StringVar(&privateKey, "key", "", "the private key for the chain")
	flag.StringVar(&storageType, "storageType", "", "Wheter to choose IPFS or redis as storage")
	flag.StringVar(&scriptsArchive, "scripts", "", "Which model to use")
	flag.StringVar(&configFile, "config", "", "Which model to use")
	flag.StringVar(&weightsFile, "weights", "", "Which model to use")
	flag.StringVar(&trainerFile, "trainer", "", "Which trainers to allow")

	var redisPassword string
	flag.StringVar(&redisPassword, "redisPW", "", "(Optional) The password for redis if used")

	flag.Parse()

	seen := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) { seen[f.Name] = true })
	for _, req := range required {
		if !seen[req] {
			err := fmt.Errorf("missing required -%s argument/flag\n", req)
			logger.Fatal(err)
		}
	}

	chain, store, err := model.CreateNetwork(chainConnection, storageConnection, storageType, privateKey, redisPassword)
	if err != nil {
		logger.Fatal(err)
	}

	config := loadDataFromDisk(configFile)
	weights := loadDataFromDisk(weightsFile)
	scripts := loadDataFromDisk(scriptsArchive)
	trainers := strings.Split(loadDataFromDisk(trainerFile), "\n")

	modelID, err := model.Deploy(
		config,
		weights,
		scripts,
		store,
		chain,
		common.Hyperparameters{
			UpdatesTillAggregation: 3,
			Epochs:                 3,
		},
	)
	if err != nil {
		logger.Fatal(err)
	}
	fmt.Println("Contract ID: " + string(modelID))

	// Add allowed trainers to contract
	for _, acc := range trainers {
		address := ethCommon.HexToAddress(acc)
		err := chain.AddTrainer(modelID, common.TrainerIdentifier(address))
		logger.Debugln("Added " + string(acc) + " as trainer")
		if err != nil {
			logger.Fatal(err)
		}
	}

}

func loadDataFromDisk(file string) (data string) {
	var content []byte

	content, err := ioutil.ReadFile(file)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Debugf("Read %d bytes from %s", len(content), file)

	data = string(content)
	return
}

package control

import (
	"log"
	"time"

	"github.com/FMorsbach/DecFL/communcation/chain"
	"github.com/FMorsbach/DecFL/communcation/storage"
	"github.com/FMorsbach/DecFL/models/MNIST"
	"github.com/FMorsbach/DecFL/training/tensorflow"
)

func Initialize() {
	config, weights := MNIST.GenerateInitialModel()

	configAddress, weightsAddress := storage.StoreInitialModel(config, weights)
	chain.DeployInitialModel(configAddress, weightsAddress)
}

func Iterate() {
	// load the storage addresses from the chain
	configAddress := chain.ModelConfigurationAddress()
	weightsAddress := chain.GlobalWeightsAddress()

	// load the model from the storage
	config, err := storage.LoadGlobalState(configAddress)
	if err != nil {
		log.Fatal(err)
	}

	weights, err := storage.LoadGlobalState(weightsAddress)
	if err != nil {
		log.Fatal(err)
	}

	// train locally
	localUpdate, err := tensorflow.Train(config, weights)
	if err != nil {
		log.Fatal(err)
	}

	// write the update to the storage
	updateAddress, err := storage.StoreUpdates(localUpdate)

	// write the address of the stored update to the chain
	chain.AppendUpdateAddress(string(time.Now().Unix()), updateAddress)

}

func Aggregate() {
	// load the local udpate addresses from the chain
	updateAddresses := chain.LocalUpdateAddresses()

	// load the local updates from storage
	updates := storage.LocalUpdates(updateAddresses)

	// aggregate the local updates
	globalWeights := tensorflow.Aggregate(updates)

	// write the new global weights to storage
	globalWeightsAddress, err := storage.StoreUpdates(globalWeights)
	if err != nil {
		log.Fatal(err)
	}

	// write the new global weights storage address to the chain
	chain.UpdateGlobalWeightsAddress(globalWeightsAddress)

	// empty the local update storage
	chain.CleanLocalUpdateStore()
}

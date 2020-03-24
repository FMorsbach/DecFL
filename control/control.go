package control

import (
	"log"
	"time"

	"github.com/FMorsbach/DecFL/communcation/chain"
	"github.com/FMorsbach/DecFL/communcation/storage"
	"github.com/FMorsbach/DecFL/models/MNIST"
	"github.com/FMorsbach/DecFL/training/tensorflow"
)

func Initialize() (trainingsID string, err error) {

	config, weights := MNIST.GenerateInitialModel()

	configAddress, weightsAddress := storage.StoreInitialModel(config, weights)
	chain.DeployInitialModel(configAddress, weightsAddress)

	return
}

func Iterate() (err error) {

	config, weights, err := globalModel()
	if err != nil {
		return err
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

	return
}

func Aggregate() (err error) {

	// load the local udpate addresses from the chain
	updateAddresses := chain.LocalUpdateAddresses()

	// load the local updates from storage
	updates := storage.LocalUpdates(updateAddresses)

	// aggregate the local updates
	globalWeights, err := tensorflow.Aggregate(updates)
	if err != nil {
		log.Fatal(err)
	}

	// write the new global weights to storage
	globalWeightsAddress, err := storage.StoreUpdates(globalWeights)
	if err != nil {
		log.Fatal(err)
	}

	// write the new global weights storage address to the chain
	chain.SetGlobalWeightsAddress(globalWeightsAddress)

	// empty the local update storage
	chain.CleanLocalUpdateStore()

	return
}

func Status() (status string, err error) {

	config, weights, err := globalModel()
	if err != nil {
		return "", err
	}

	results, err := tensorflow.Evaluate(config, weights)
	if err != nil {
		return "", err
	}

	status = results.String()
	return
}

func CleanUpEnvironment() {
	//cleanUpStorage()
	//cleanUpChain()
}

func globalModel() (config string, weights string, err error) {
	// load the storage addresses from the chain
	configAddress := chain.ModelConfigurationAddress()
	weightsAddress := chain.GlobalWeightsAddress()

	// load the model from the storage
	config, err = storage.LoadGlobalState(configAddress)
	if err != nil {
		return "", "", err
	}

	weights, err = storage.LoadGlobalState(weightsAddress)
	if err != nil {
		return "", "", err
	}

	return
}

package control

import (
	"log"
	"os"
	"time"

	"github.com/FMorsbach/DecFL/communcation/chain"
	"github.com/FMorsbach/DecFL/communcation/storage"
	"github.com/FMorsbach/DecFL/models/MNIST"
	"github.com/FMorsbach/DecFL/training"
	"github.com/FMorsbach/DecFL/training/tensorflow"
	"github.com/FMorsbach/dlog"
)

var logger = dlog.New(os.Stderr, "Control: ", log.LstdFlags, false)

func EnableDebug(b bool) {
	logger.SetDebug(b)
}

func Initialize() (trainingsID string, err error) {

	config, weights := MNIST.GenerateInitialModel()
	logger.Debug("Created initial model")

	configAddress, weightsAddress := storage.StoreInitialModel(config, weights)
	logger.Debugf("Wrote initial model to storage at %s and %s", configAddress, weightsAddress)

	chain.DeployInitialModel(configAddress, weightsAddress)
	logger.Debug(("Wrote initial model addresses to chain"))
	return
}

func Iterate() (err error) {

	config, weights, err := globalModel()
	if err != nil {
		return err
	}
	logger.Debug("Loaded model from network")

	// train locally
	localUpdate, err := tensorflow.Train(config, weights)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Debug("Trained local model")

	// write the update to the storage
	updateAddress, err := storage.StoreUpdate(localUpdate)
	logger.Debugf("Wrote local update to storage at %s", updateAddress)

	// write the address of the stored update to the chain
	chain.AppendUpdateAddress(string(time.Now().Unix()), updateAddress)
	logger.Debug("Wrote local update address to chain")

	return
}

func Aggregate() (err error) {

	// load the local udpate addresses from the chain
	updateAddresses := chain.LocalUpdateAddresses()
	logger.Debug("Loaded update addresses from chain")

	// load the local updates from storage
	updates := storage.LocalUpdates(updateAddresses)
	logger.Debug("Loaded updates from storage")

	// aggregate the local updates
	globalWeights, err := tensorflow.Aggregate(updates)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Debug("Aggregated updates")

	// write the new global weights to storage
	globalWeightsAddress, err := storage.StoreUpdate(globalWeights)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Debugf("Wrote new weights to storage at %s", globalWeightsAddress)

	// write the new global weights storage address to the chain
	chain.SetGlobalWeightsAddress(globalWeightsAddress)
	logger.Debug("Wrote new weight address to chain")

	// empty the local update storage
	chain.CleanLocalUpdateStore()
	logger.Debug("Cleaned local update list on chain")

	return
}

func Status() (status training.EvaluationResults, err error) {

	config, weights, err := globalModel()
	if err != nil {
		return
	}
	logger.Debug("Loaded model from network")

	status, err = tensorflow.Evaluate(config, weights)
	if err != nil {
		return
	}
	logger.Debug("Evaluated model")

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

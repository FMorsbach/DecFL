package control

import (
	"log"
	"os"

	c "github.com/FMorsbach/DecFL/communication"
	bc "github.com/FMorsbach/DecFL/communication/chain"
	"github.com/FMorsbach/DecFL/communication/mocks"
	"github.com/FMorsbach/DecFL/communication/storage"
	"github.com/FMorsbach/DecFL/models/MNIST"
	"github.com/FMorsbach/DecFL/training"
	"github.com/FMorsbach/DecFL/training/tensorflow"
	"github.com/FMorsbach/dlog"
)

var logger = dlog.New(os.Stderr, "Control: ", log.LstdFlags, false)
var chain bc.Chain = mocks.NewRedis()
var store storage.Storage = mocks.NewRedis()

func EnableDebug(b bool) {
	logger.SetDebug(b)
}

func Initialize() (modelID c.ModelIdentifier, err error) {

	config, weights := MNIST.GenerateInitialModel()
	logger.Debug("Created initial model")

	configAddress, err := store.Store(config)
	if err != nil {
		return
	}

	weightsAddress, err := store.Store(weights)
	if err != nil {
		return
	}
	logger.Debugf("Wrote initial model to storage at %s and %s", configAddress, weightsAddress)

	modelID, err = chain.DeployModel(configAddress, weightsAddress)
	logger.Debug(("Wrote initial model addresses to chain"))
	return
}

func Iterate(modelID c.ModelIdentifier) (err error) {

	config, weights, err := globalModel(modelID)
	if err != nil {
		return err
	}
	logger.Debug("Loaded model from network")

	// train locally
	localUpdate, err := tensorflow.Train(config, weights)
	if err != nil {
		return
	}
	logger.Debug("Trained local model")

	// write the update to the storage
	updateAddress, err := store.Store(localUpdate)
	if err != nil {
		return
	}
	logger.Debugf("Wrote local update to storage at %s", updateAddress)

	// write the address of the stored update to the chain
	err = chain.SubmitLocalUpdate(modelID, c.StorageAddress(updateAddress))
	if err != nil {
		return
	}
	logger.Debug("Wrote local update address to chain")

	return
}

func Aggregate(modelID c.ModelIdentifier) (err error) {

	// load the local udpate addresses from the chain
	updateAddresses, err := chain.LocalUpdateAddresses(modelID)
	if err != nil {
		return
	}
	logger.Debug("Loaded update addresses from chain")

	// load the local updates from storage
	updates, err := store.Loads(updateAddresses)
	if err != nil {
		return
	}
	logger.Debug("Loaded updates from storage")

	// aggregate the local updates
	globalWeights, err := tensorflow.Aggregate(updates)
	if err != nil {
		return
	}
	logger.Debug("Aggregated updates")

	// write the new global weights to storage
	globalWeightsAddress, err := store.Store(globalWeights)
	if err != nil {
		return
	}
	logger.Debugf("Wrote new weights to storage at %s", globalWeightsAddress)

	// write the new global weights storage address to the chain
	err = chain.SetGlobalWeightsAddress(modelID, globalWeightsAddress)
	if err != nil {
		return
	}
	logger.Debug("Wrote new weight address to chain")

	// empty the local update storage
	err = chain.ClearLocalUpdateAddresses(modelID)
	if err != nil {
		return
	}
	logger.Debug("Cleaned local update list on chain")

	return
}

func Status(modelID c.ModelIdentifier) (status training.EvaluationResults, err error) {

	config, weights, err := globalModel(modelID)
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

func globalModel(modelID c.ModelIdentifier) (config string, weights string, err error) {
	// load the storage addresses from the chain
	configAddress, err := chain.ModelConfigurationAddress(modelID)
	if err != nil {
		return
	}

	weightsAddress, err := chain.GlobalWeightsAddress(modelID)
	if err != nil {
		return
	}

	// load the model from the storage
	config, err = store.Load(configAddress)
	if err != nil {
		return
	}

	weights, err = store.Load(weightsAddress)
	if err != nil {
		return
	}

	return
}

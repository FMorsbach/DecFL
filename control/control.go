package control

import (
	"log"
	"os"

	c "github.com/FMorsbach/DecFL/communication"
	bc "github.com/FMorsbach/DecFL/communication/chain"
	"github.com/FMorsbach/DecFL/communication/storage"
	"github.com/FMorsbach/DecFL/training"
	"github.com/FMorsbach/dlog"
)

type Control interface {
	Initialize(configuration string, weights string, params bc.Hyperparameters) (modelID c.ModelIdentifier, err error)
	Iterate(modelID c.ModelIdentifier, trainerID c.TrainerIdentifier) (err error)
	Aggregate(modelID c.ModelIdentifier) (err error)
	Status(modelID c.ModelIdentifier) (status training.EvaluationResults, err error)
}

type ctlImpl struct {
	chain bc.Chain
	store storage.Storage
	mlf   training.MLFramework
}

var logger = dlog.New(os.Stderr, "Control: ", log.LstdFlags, false)

func EnableDebug(b bool) {
	logger.SetDebug(b)
}

func NewControl(ch bc.Chain, st storage.Storage, mlf training.MLFramework) Control {
	return &ctlImpl{
		chain: ch,
		store: st,
		mlf:   mlf,
	}
}

func (ctl *ctlImpl) Initialize(configuration string, weights string, params bc.Hyperparameters) (modelID c.ModelIdentifier, err error) {

	logger.Debug("Created initial model")

	configAddress, err := ctl.store.Store(configuration)
	if err != nil {
		return
	}

	weightsAddress, err := ctl.store.Store(weights)
	if err != nil {
		return
	}
	logger.Debugf("Wrote initial model to storage at %s and %s", configAddress, weightsAddress)

	modelID, err = ctl.chain.DeployModel(configAddress, weightsAddress, params)
	logger.Debug(("Wrote initial model addresses to chain"))
	return
}

func (ctl *ctlImpl) Iterate(modelID c.ModelIdentifier, trainerID c.TrainerIdentifier) (err error) {

	config, weights, err := ctl.globalModel(modelID)
	if err != nil {
		return err
	}
	logger.Debug("Loaded model from network")

	// train locally
	localUpdate, err := ctl.mlf.Train(config, weights)
	if err != nil {
		return
	}
	logger.Debug("Trained local model")

	// write the update to the storage
	updateAddress, err := ctl.store.Store(localUpdate)
	if err != nil {
		return
	}
	logger.Debugf("Wrote local update to storage at %s", updateAddress)

	update := c.Update{
		Trainer: trainerID,
		Address: updateAddress,
	}
	// write the address of the stored update to the chain
	err = ctl.chain.SubmitLocalUpdate(modelID, update)
	if err != nil {
		return
	}
	logger.Debug("Wrote local update address to chain")

	return
}

func (ctl *ctlImpl) Aggregate(modelID c.ModelIdentifier) (err error) {

	// load the local udpate addresses from the chain
	localUpdates, err := ctl.chain.LocalUpdates(modelID)
	if err != nil {
		return
	}
	logger.Debug("Loaded update addresses from chain")

	// load the local updates from storage
	updates := make([]string, len(localUpdates))
	for i, localUpdate := range localUpdates {
		updates[i], err = ctl.store.Load(localUpdate.Address)
		if err != nil {
			return
		}
	}
	logger.Debug("Loaded updates from storage")

	// aggregate the local updates
	globalWeights, err := ctl.mlf.Aggregate(updates)
	if err != nil {
		return
	}
	logger.Debug("Aggregated updates")

	// write the new global weights to storage
	globalWeightsAddress, err := ctl.store.Store(globalWeights)
	if err != nil {
		return
	}
	logger.Debugf("Wrote new weights to storage at %s", globalWeightsAddress)

	// write the new global weights storage address to the chain
	err = ctl.chain.SubmitAggregation(modelID, globalWeightsAddress)
	if err != nil {
		return
	}
	logger.Debug("Wrote new weight address to chain")

	return
}

func (ctl *ctlImpl) Status(modelID c.ModelIdentifier) (status training.EvaluationResults, err error) {

	config, weights, err := ctl.globalModel(modelID)
	if err != nil {
		return
	}
	logger.Debug("Loaded model from network")

	status, err = ctl.mlf.Evaluate(config, weights)
	if err != nil {
		return
	}
	logger.Debug("Evaluated model")

	return
}

func (ctl *ctlImpl) globalModel(modelID c.ModelIdentifier) (config string, weights string, err error) {
	// load the storage addresses from the chain
	configAddress, err := ctl.chain.ModelConfigurationAddress(modelID)
	if err != nil {
		return
	}

	weightsAddress, err := ctl.chain.GlobalWeightsAddress(modelID)
	if err != nil {
		return
	}

	// load the model from the storage
	config, err = ctl.store.Load(configAddress)
	if err != nil {
		return
	}

	weights, err = ctl.store.Load(weightsAddress)
	if err != nil {
		return
	}

	return
}

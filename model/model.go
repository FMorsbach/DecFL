package model

import (
	"log"
	"os"
	"time"

	"github.com/FMorsbach/DecFL/model/chain"
	"github.com/FMorsbach/DecFL/model/common"
	"github.com/FMorsbach/DecFL/model/storage"
	"github.com/FMorsbach/DecFL/model/training"
	"github.com/FMorsbach/dlog"
)

type Model interface {
	Iterate() (err error)
	Aggregate() (err error)
	Status() (status training.EvaluationResults, err error)
	WaitForNewEpoch() (err error)
	WaitForAggregation() (err error)
}

type ctlImpl struct {
	chain      chain.Chain
	store      storage.Storage
	mlf        training.MLFramework
	modelID    common.ModelIdentifier
	localEpoch int
}

var logger = dlog.New(os.Stderr, "Model: ", log.LstdFlags, false)

func EnableDebug(b bool) {
	logger.SetDebug(b)
}

func NewControl(ch chain.Chain, st storage.Storage, mlf training.MLFramework, modelID common.ModelIdentifier) (Model, error) {

	localEpoch, err := ch.ModelEpoch(modelID)
	if err != nil {
		return nil, err
	}

	return &ctlImpl{
		chain:      ch,
		store:      st,
		mlf:        mlf,
		modelID:    modelID,
		localEpoch: localEpoch,
	}, nil
}

func Deploy(configuration string, weights string, store storage.Storage, ch chain.Chain, params common.Hyperparameters) (modelID common.ModelIdentifier, err error) {

	logger.Debug("Created initial model")

	configAddress, err := store.Store(configuration)
	if err != nil {
		return
	}

	weightsAddress, err := store.Store(weights)
	if err != nil {
		return
	}
	logger.Debugf("Wrote initial model to storage at %s and %s", configAddress, weightsAddress)

	modelID, err = ch.DeployModel(configAddress, weightsAddress, params)
	logger.Debug(("Wrote initial model addresses to chain"))
	return
}

func (mod *ctlImpl) Iterate() (err error) {

	config, weights, err := mod.globalModel()
	if err != nil {
		return err
	}
	logger.Debug("Loaded model from network")

	// train locally
	localUpdate, err := mod.mlf.Train(config, weights)
	if err != nil {
		return
	}
	logger.Debug("Trained local model")

	eval, err := mod.mlf.Evaluate(config, localUpdate)
	if err != nil {
		return
	}
	logger.Printf("LOCAL: %f Accuracy\n", eval.Accuracy)

	// write the update to the storage
	updateAddress, err := mod.store.Store(localUpdate)
	if err != nil {
		return
	}
	logger.Debugf("Wrote local update to storage at %s", updateAddress)

	// write the address of the stored update to the chain
	err = mod.chain.SubmitLocalUpdate(mod.modelID, updateAddress)
	if err != nil {
		return
	}
	logger.Debug("Wrote local update address to chain")

	return
}

func (mod *ctlImpl) Aggregate() (err error) {

	// load the local udpate addresses from the chain
	localUpdates, err := mod.chain.LocalUpdates(mod.modelID)
	if err != nil {
		return
	}
	logger.Debug("Loaded update addresses from chain")

	// load the local updates from storage
	updates := make([]string, len(localUpdates))
	for i, localUpdate := range localUpdates {
		updates[i], err = mod.store.Load(localUpdate.Address)
		if err != nil {
			return
		}
	}
	logger.Debug("Loaded updates from storage")

	// aggregate the local updates
	globalWeights, err := mod.mlf.Aggregate(updates)
	if err != nil {
		return
	}
	logger.Debug("Aggregated updates")

	// write the new global weights to storage
	globalWeightsAddress, err := mod.store.Store(globalWeights)
	if err != nil {
		return
	}
	logger.Debugf("Wrote new weights to storage at %s", globalWeightsAddress)

	// write the new global weights storage address to the chain
	err = mod.chain.SubmitAggregation(mod.modelID, globalWeightsAddress)
	if err != nil {
		return
	}
	logger.Debug("Wrote new weight address to chain")

	mod.localEpoch++

	return
}

func (mod *ctlImpl) Status() (status training.EvaluationResults, err error) {

	config, weights, err := mod.globalModel()
	if err != nil {
		return
	}
	logger.Debug("Loaded model from network")

	status, err = mod.mlf.Evaluate(config, weights)
	if err != nil {
		return
	}
	logger.Debug("Evaluated model")

	return
}

func (mod *ctlImpl) globalModel() (config string, weights string, err error) {

	// load the storage addresses from the chain
	configAddress, err := mod.chain.ModelConfigurationAddress(mod.modelID)
	if err != nil {
		return
	}

	weightsAddress, err := mod.chain.GlobalWeightsAddress(mod.modelID)
	if err != nil {
		return
	}

	// load the model from the storage
	config, err = mod.store.Load(configAddress)
	if err != nil {
		return
	}

	weights, err = mod.store.Load(weightsAddress)
	if err != nil {
		return
	}

	return
}

func (mod *ctlImpl) WaitForNewEpoch() (err error) {

	globalEpoch, err := mod.chain.ModelEpoch(mod.modelID)
	if err != nil {
		return
	}

	for mod.localEpoch > globalEpoch {

		time.Sleep(time.Second)

		globalEpoch, err = mod.chain.ModelEpoch(mod.modelID)
		if err != nil {
			return err
		}
	}
	return
}

func (mod *ctlImpl) WaitForAggregation() (err error) {

	ready, err := mod.chain.AggregationReady(mod.modelID)
	if err != nil {
		return
	}

	for !ready {

		time.Sleep(time.Second)

		ready, err = mod.chain.AggregationReady(mod.modelID)
		if err != nil {
			return
		}
	}

	return
}

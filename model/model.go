package model

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
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
	Evaluate() (evaluation training.EvaluationResults, err error)
	Epoch() (epoch int, err error)
	State() (state uint8, err error)
}

type modelImpl struct {
	chain       chain.Chain
	store       storage.Storage
	mlf         training.MLFramework
	modelID     common.ModelIdentifier
	localEpoch  int
	modelConfig string
}

var logger = dlog.New(os.Stderr, "Model: ", log.LstdFlags, false)
var logB = log.New(os.Stdout, "BENCH: ", log.LstdFlags|log.Lmsgprefix)
var start time.Time

const LOAD_WEIGHTS_ADDRESS int = 0

func EnableDebug(b bool) {
	logger.SetDebug(b)
}

func NewModel(ch chain.Chain, st storage.Storage, mlf training.MLFramework, modelID common.ModelIdentifier) (Model, error) {

	localEpoch, err := ch.ModelEpoch(modelID)
	if err != nil {
		return nil, err
	}
	logger.Debugf("Retrieved model epoch %d", localEpoch)

	modelAddress, err := ch.ModelConfigurationAddress(modelID)
	if err != nil {
		return nil, err
	}
	logger.Debugf("Retrieved model config address %s", string(modelAddress))

	config, err := st.Load(modelAddress)
	if err != nil {
		return nil, err
	}
	logger.Debugf("Loaded model config from storage")

	scriptsAddress, err := ch.ScriptsAddress(modelID)
	if err != nil {
		return nil, err
	}

	scripts, err := st.Load(scriptsAddress)
	if err != nil {
		return nil, err
	}
	logger.Debugf("Loaded scripts from storage")

	path, exists := os.LookupEnv("DECFL_ROOT")
	if !exists {
		logger.Fatal("DECFL_ROOT is not set.")
	}
	scriptsArchive := filepath.Join(path, "scripts.tar.gz")
	err = ioutil.WriteFile(scriptsArchive, []byte(scripts), 0644)
	if err != nil {
		return nil, err
	}
	logger.Debugf("Wrote scripts to %s", scriptsArchive)

	cmd := exec.Command(
		"tar",
		"-xzvf",
		scriptsArchive,
		"-C",
		path,
	)

	logger.Debugln("Executing: ", cmd.Args)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("Could not extract scripts: %s, %s", err, string(out))
	}
	logger.Debug(string(out))

	return &modelImpl{
		chain:       ch,
		store:       st,
		mlf:         mlf,
		modelID:     modelID,
		localEpoch:  localEpoch,
		modelConfig: config,
	}, nil
}

func Deploy(configuration string, weights string, scripts string, store storage.Storage, ch chain.Chain, params common.Hyperparameters) (modelID common.ModelIdentifier, err error) {

	logger.Debug("Created initial model")

	configAddress, err := store.Store(configuration)
	if err != nil {
		return
	}

	weightsAddress, err := store.Store(weights)
	if err != nil {
		return
	}

	scriptsAddress, err := store.Store(scripts)
	if err != nil {
		return
	}

	logger.Debugf("Wrote initial model to storage at %s, %s and %s", configAddress, weightsAddress, scriptsAddress)

	modelID, err = ch.DeployModel(configAddress, weightsAddress, scriptsAddress, params)
	logger.Debug(("Wrote initial model addresses to chain"))
	return
}

func (mod *modelImpl) Iterate() (err error) {

	weights, err := mod.globalWeights()
	if err != nil {
		return err
	}
	logger.Debug("Loaded model from network")

	// train locally
	start = time.Now()
	localUpdate, err := mod.mlf.Train(mod.modelConfig, weights)
	if err != nil {
		return
	}
	logger.Debug("Trained local model")
	logB.Printf("TRAIN_MODEL %.3f\n", time.Since(start).Seconds())

	start = time.Now()
	eval, err := mod.mlf.Evaluate(mod.modelConfig, localUpdate)
	if err != nil {
		return
	}
	logB.Printf("EVALUATE_MODEL %.3f\n", time.Since(start).Seconds())

	start = time.Now()
	// write the update to the storage
	updateAddress, err := mod.store.Store(localUpdate)
	if err != nil {
		return
	}
	logger.Debugf("Wrote local update to storage at %s", updateAddress)
	logB.Printf("STORE_TRAINING_UPDATE %.3f\n", time.Since(start).Seconds())

	logger.Printf("LOCAL TRAINING: %s with %f Accuracy\n", string(updateAddress)[0:6], eval.Accuracy)

	// write the address of the stored update to the chain
	start = time.Now()
	err = mod.chain.SubmitLocalUpdate(mod.modelID, updateAddress)
	if err != nil {
		return
	}
	logger.Debug("Wrote local update address to chain")
	logB.Printf("SUBMIT_TRAINING_UPDATE %.3f\n", time.Since(start).Seconds())

	return
}

func (mod *modelImpl) Aggregate() (err error) {

	start = time.Now()
	// load the local udpate addresses from the chain
	localUpdates, err := mod.chain.LocalUpdates(mod.modelID)
	if err != nil {
		return
	}
	logger.Debug("Loaded update addresses from chain")
	logB.Printf("LOAD_AGGREGATION_WEIGHTS_ADDRESSES %.3f\n", time.Since(start).Seconds())

	start = time.Now()
	// load the local updates from storage
	updates := make([]string, len(localUpdates))
	for i, localUpdate := range localUpdates {
		updates[i], err = mod.store.Load(localUpdate.Address)
		if err != nil {
			return
		}
	}
	logger.Debug("Loaded updates from storage")
	logB.Printf("DOWNLOAD_AGGREGATION_WEIGHTS %.3f\n", time.Since(start).Seconds())

	updateHashes := make([]string, len(localUpdates))
	for i, update := range updates {
		h := sha256.Sum256([]byte(update))
		updateHashes[i] = hex.EncodeToString(h[0:32])[0:6]
	}
	logger.Printf("AGGREGATING: %s\n", updateHashes)

	start = time.Now()
	// aggregate the local updates
	globalWeights, err := mod.mlf.Aggregate(updates)
	if err != nil {
		return
	}
	logger.Debug("Aggregated updates")
	logB.Printf("AGGREGATE_MODELS %.3f\n", time.Since(start).Seconds())

	start = time.Now()
	eval, err := mod.mlf.Evaluate(mod.modelConfig, globalWeights)
	if err != nil {
		return
	}
	logB.Printf("EVALUTE_MODEL %.3f\n", time.Since(start).Seconds())

	start = time.Now()
	// write the new global weights to storage
	globalWeightsAddress, err := mod.store.Store(globalWeights)
	if err != nil {
		return
	}
	logger.Debugf("Wrote new weights to storage at %s", globalWeightsAddress)
	logB.Printf("STORE_AGGREGATION_CANDIDATE %.3f\n", time.Since(start).Seconds())

	logger.Printf("AGGREGATION: %s with %f Accuracy\n", string(globalWeightsAddress)[0:6], eval.Accuracy)

	start = time.Now()
	// write the new global weights storage address to the chain
	err = mod.chain.SubmitAggregation(mod.modelID, globalWeightsAddress)
	if err != nil {
		return
	}
	logger.Debug("Wrote new weight address to chain")
	logB.Printf("SUBMIT_AGGREGATION_CANDIDATE %.3f\n", time.Since(start).Seconds())

	mod.localEpoch++

	return
}

func (mod *modelImpl) Evaluate() (evaluation training.EvaluationResults, err error) {

	weights, err := mod.globalWeights()
	if err != nil {
		return
	}
	logger.Debug("Loaded model from network")

	evaluation, err = mod.mlf.Evaluate(mod.modelConfig, weights)
	if err != nil {
		return
	}
	logger.Debug("Evaluated model")

	return
}

func (mod *modelImpl) Epoch() (epoch int, err error) {
	return mod.chain.ModelEpoch(mod.modelID)
}

func (mod *modelImpl) State() (state uint8, err error) {
	return mod.chain.State(mod.modelID)
}

func (mod *modelImpl) globalWeights() (weights string, err error) {

	start = time.Now()
	weightsAddress, err := mod.chain.GlobalWeightsAddress(mod.modelID)
	if err != nil {
		return
	}
	logB.Printf("LOAD_WEIGHT_ADDRESS %.3f\n", time.Since(start).Seconds())

	start = time.Now()
	weights, err = mod.store.Load(weightsAddress)
	if err != nil {
		return
	}
	logB.Printf("DOWNLOAD_WEIGHTS %.3f\n", time.Since(start).Seconds())

	return
}

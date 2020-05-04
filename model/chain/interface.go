package chain

import "github.com/FMorsbach/DecFL/model/common"

type Chain interface {
	DeployModel(configAddress common.StorageAddress, weightsAddress common.StorageAddress, params common.Hyperparameters) (id common.ModelIdentifier, err error)
	ModelEpoch(id common.ModelIdentifier) (epoch int, err error)
	ModelConfigurationAddress(id common.ModelIdentifier) (address common.StorageAddress, err error)
	GlobalWeightsAddress(id common.ModelIdentifier) (address common.StorageAddress, err error)
	SubmitLocalUpdate(modelID common.ModelIdentifier, updateAddress common.StorageAddress) (err error)
	LocalUpdates(id common.ModelIdentifier) (updates []common.Update, err error)
	SubmitAggregation(id common.ModelIdentifier, address common.StorageAddress) (err error)
	State(id common.ModelIdentifier) (state uint8, err error)
}

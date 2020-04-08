package chain

import (
	c "github.com/FMorsbach/DecFL/communication"
)

type Hyperparameters struct {
	UpdatesTillAggregation int
}

type Chain interface {
	DeployModel(configAddress c.StorageAddress, weightsAddress c.StorageAddress, params Hyperparameters) (id c.ModelIdentifier, err error)
	ModelEpoch(id c.ModelIdentifier) (epoch int, err error)
	ModelConfigurationAddress(id c.ModelIdentifier) (address c.StorageAddress, err error)
	GlobalWeightsAddress(id c.ModelIdentifier) (address c.StorageAddress, err error)
	SubmitLocalUpdate(modelID c.ModelIdentifier, update c.Update) (err error)
	LocalUpdates(id c.ModelIdentifier) (updates []c.Update, err error)
	SubmitAggregation(id c.ModelIdentifier, address c.StorageAddress) (err error)
}

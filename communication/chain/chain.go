package chain

import (
	"log"
	"os"

	c "github.com/FMorsbach/DecFL/communication"
	"github.com/FMorsbach/dlog"
)

type ModelIdentifier string

type Chain interface {
	DeployModel(configAddress c.StorageAddress, weightsAddress c.StorageAddress) (id ModelIdentifier, err error)
	ModelConfigurationAddress(id ModelIdentifier) (address c.StorageAddress, err error)
	GlobalWeightsAddress(id ModelIdentifier) (address c.StorageAddress, err error)
	SetGlobalWeightsAddress(id ModelIdentifier, address c.StorageAddress) (err error)
	SubmitLocalUpdate(id ModelIdentifier, address c.StorageAddress) (err error)
	LocalUpdateAddresses(id ModelIdentifier) (addresses []c.StorageAddress, err error)
	ClearLocalUpdateAddresses(id ModelIdentifier) (err error)
}

var logger = dlog.New(os.Stderr, "Chain: ", log.LstdFlags, false)

func EnableDebug(b bool) {
	logger.SetDebug(b)
}

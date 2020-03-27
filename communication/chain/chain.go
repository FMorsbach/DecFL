package chain

import (
	"log"
	"os"

	"github.com/FMorsbach/dlog"
)

type StorageAddress string
type ModelIdentifier string

type Chain interface {
	DeployModel(configAddress StorageAddress, weightsAddress StorageAddress) (id ModelIdentifier, err error)
	ModelConfigurationAddress(id ModelIdentifier) (address StorageAddress, err error)
	GlobalWeightsAddress(id ModelIdentifier) (address StorageAddress, err error)
	SetGlobalWeightsAddress(id ModelIdentifier, address StorageAddress) (err error)
	SubmitLocalUpdate(id ModelIdentifier, address StorageAddress) (err error)
	LocalUpdateAddresses(id ModelIdentifier) (addresses []StorageAddress, err error)
	ClearLocalUpdateAddresses(id ModelIdentifier) (err error)
}

var logger = dlog.New(os.Stderr, "Chain: ", log.LstdFlags, false)

func EnableDebug(b bool) {
	logger.SetDebug(b)
}

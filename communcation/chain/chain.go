package chain

import (
	"log"
	"os"

	"github.com/FMorsbach/dlog"
)

type storageAddress string
type trainerIdentifier string
type modelIdentifier string

type Chain interface {
	DeployModel(configAddress storageAddress, weightsAddress storageAddress) (id modelIdentifier, err error)
	ModelConfigurationAddress(id modelIdentifier) (address storageAddress, err error)
	GlobalWeightsAddress(id modelIdentifier) (address storageAddress, err error)
	SetGlobalWeightsAddress(id modelIdentifier, address storageAddress) (err error)
	SubmitLocalUpdate(id modelIdentifier, address storageAddress) (err error)
	LocalUpdateAddresses(id modelIdentifier) (addresses []storageAddress, err error)
	ClearLocalUpdateAddresses(id modelIdentifier) (err error)
}

var logger = dlog.New(os.Stderr, "Chain: ", log.LstdFlags, false)

func EnableDebug(b bool) {
	logger.SetDebug(b)
}

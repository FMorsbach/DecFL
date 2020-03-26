package chain

import (
	"log"
	"os"

	"github.com/FMorsbach/dlog"
)

type storageAddress string
type trainerIdentification string

type Chain interface {
	DeployModel(configuration string, weights string) (modelIdentitifier string, err error)
	ModelConfigurationAddress(modelIdentitifier string) (configurationAddress storageAddress, err error)
	GlobalWeightsAddress(modelIdentitifier string) (weightsAddress storageAddress, err error)
	SetGlobalWeightsAddress(address storageAddress) (err error)
	SubmitLocalUpdate(trainer trainerIdentification, address storageAddress) (err error)
	LocalUpdateAddresses() (localUpdateAddresses []storageAddress, err error)
	ClearLocalUpdateAddresses() (err error)
}

var logger = dlog.New(os.Stderr, "Chain: ", log.LstdFlags, false)

func EnableDebug(b bool) {
	logger.SetDebug(b)
}

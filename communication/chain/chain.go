package chain

import (
	"log"
	"os"

	c "github.com/FMorsbach/DecFL/communication"
	"github.com/FMorsbach/dlog"
)

type Chain interface {
	DeployModel(configAddress c.StorageAddress, weightsAddress c.StorageAddress) (id c.ModelIdentifier, err error)
	ModelConfigurationAddress(id c.ModelIdentifier) (address c.StorageAddress, err error)
	GlobalWeightsAddress(id c.ModelIdentifier) (address c.StorageAddress, err error)
	SetGlobalWeightsAddress(id c.ModelIdentifier, address c.StorageAddress) (err error)
	SubmitLocalUpdate(id c.ModelIdentifier, address c.StorageAddress) (err error)
	LocalUpdateAddresses(id c.ModelIdentifier) (addresses []c.StorageAddress, err error)
	ClearLocalUpdateAddresses(id c.ModelIdentifier) (err error)
}

var logger = dlog.New(os.Stderr, "Chain: ", log.LstdFlags, false)

func EnableDebug(b bool) {
	logger.SetDebug(b)
}

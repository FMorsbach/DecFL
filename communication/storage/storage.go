package storage

import (
	"log"
	"os"

	"github.com/FMorsbach/DecFL/communication"
	"github.com/FMorsbach/dlog"
)

type Storage interface {
	Store(content string) (address communication.StorageAddress, err error)
	Load(address communication.StorageAddress) (content string, err error)
	Loads(addresses []communication.StorageAddress) (content []string, err error)
}

var logger = dlog.New(os.Stderr, "Storage: ", log.LstdFlags, false)

func EnableDebug(b bool) {
	logger.SetDebug(b)
}

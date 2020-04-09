package storage

import (
	"log"
	"os"

	"github.com/FMorsbach/DecFL/model/common"
	"github.com/FMorsbach/dlog"
)

type Storage interface {
	Store(content string) (address common.StorageAddress, err error)
	Load(address common.StorageAddress) (content string, err error)
}

var logger = dlog.New(os.Stderr, "Storage: ", log.LstdFlags, false)

func EnableDebug(b bool) {
	logger.SetDebug(b)
}

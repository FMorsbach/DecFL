package storage

import (
	"log"
	"os"

	"github.com/FMorsbach/dlog"
)

var logger = dlog.New(os.Stderr, "Storage: ", log.LstdFlags, false)

func EnableDebug(b bool) {
	logger.SetDebug(b)
}

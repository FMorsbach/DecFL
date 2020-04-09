package MNIST

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/FMorsbach/dlog"
)

var logger = dlog.New(os.Stderr, "Model: ", log.Flags(), false)

func init() {
	if _, exists := os.LookupEnv("DECFL_ROOT"); !exists {
		logger.Fatal("DECFL_ROOT is not set.")
	}
}

func GenerateInitialModel() (configuration string, weights string) {

	program_root := os.Getenv("DECFL_ROOT")
	configuration = loadDataFromDisk(filepath.Join(program_root, "models/MNIST/configuration.txt"))
	weights = loadDataFromDisk(filepath.Join(program_root, "models/MNIST/weights.txt"))
	return
}

func loadDataFromDisk(file string) (data string) {
	var content []byte

	content, err := ioutil.ReadFile(file)
	if err != nil {
		dlog.Fatal(err)
	}
	logger.Debugf("Read %d bytes from %s", len(content), file)

	data = string(content)
	return
}

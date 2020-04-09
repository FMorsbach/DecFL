package tensorflow

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/FMorsbach/DecFL/model/training"
	"github.com/FMorsbach/dlog"
)

type TensorflowError struct {
	Err          error
	Description  string
	PythonOutput string
}

func (e *TensorflowError) Error() string {
	return e.Description + " " + e.Err.Error() + "\n" + e.PythonOutput
}

type TensorflowFramework struct {
}

func NewTensorflowMLF() training.MLFramework {
	return &TensorflowFramework{}
}

const OUTPUT_FILE string = "output.out"
const CONFIG_FILE string = "config.in"
const WEIGHTS_FILE string = "weights.in"

var pythonPath = findExternalDependency("DECFL_PYTHON")
var resourcePath = findExternalDependency("DECFL_RES")
var scriptsPath = findExternalDependency("DECFL_SCRIPTS")

var logger = dlog.New(os.Stderr, "Training - TF: ", log.LstdFlags, false)

func findExternalDependency(envVar string) string {
	path, exists := os.LookupEnv(envVar)
	if !exists {
		logger.Fatal(envVar, " is not set.")
	}

	if _, err := os.Stat(path); err == nil {
		return path
	} else if os.IsNotExist(err) {
		panic(fmt.Sprintf("%s %s", envVar, "does not point to a valid location"))
	} else {
		panic(err)
	}
}

func EnableDebug(b bool) {
	logger.SetDebug(b)
}

func readUpdatesFromDisk() (output string, err error) {

	var content []byte

	content, err = ioutil.ReadFile(filepath.Join(resourcePath, OUTPUT_FILE))
	if err != nil {
		return
	}

	logger.Debugf("Read %d bytes from %s", len(content), OUTPUT_FILE)
	output = string(content)

	return
}

func writeModelToDisk(configuration string, weights string) (err error) {

	err = ioutil.WriteFile(filepath.Join(resourcePath, CONFIG_FILE), []byte(configuration), 0644)
	if err != nil {
		return
	}
	logger.Debugf("Wrote %d bytes as %s to disk.", len([]byte(configuration)), CONFIG_FILE)

	err = ioutil.WriteFile(filepath.Join(resourcePath, WEIGHTS_FILE), []byte(weights), 0644)
	if err != nil {
		return
	}
	logger.Debugf("Wrote %d bytes as %s to disk.", len([]byte(weights)), WEIGHTS_FILE)

	return
}

func cleanUpRessources() {

	dir, err := ioutil.ReadDir(resourcePath)
	if err != nil {
		panic(fmt.Sprintf("Tried deleting contents of %s but while reading its content got %s", resourcePath, err))
	}

	for _, d := range dir {
		if d.Name()[0:1] == "." {
			continue
		}

		err := os.RemoveAll(path.Join([]string{resourcePath, d.Name()}...))
		if err != nil && !os.IsNotExist(err) {
			panic(fmt.Sprintf("Tried deleting %s after training but got %s", d, err))
		}
	}

	logger.Debug("Cleaned up resources from disk")
}

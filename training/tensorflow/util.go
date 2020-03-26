package tensorflow

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/FMorsbach/dlog"
)

var prefix string = findPrefix()
var pythonPath string
var resourcePath string
var configPath string
var weightsPath string
var outputPath string

var logger = dlog.New(os.Stderr, "Training - TF: ", log.LstdFlags, false)

type TensorflowError struct {
	Err          error
	Description  string
	PythonOutput string
}

func (e *TensorflowError) Error() string {
	return e.Description + " " + e.Err.Error() + "\n" + e.PythonOutput
}

func findPrefix() string {

	project_root, exists := os.LookupEnv("DECFL_ROOT")
	if !exists {
		logger.Fatal("DECFL_ROOT is not set.")
	}

	prefixes := []string{
		"",
		filepath.Join(project_root, "training/tensorflow"),
	}

	for _, pre := range prefixes {
		if _, err := os.Stat(filepath.Join(pre, "venv/bin/python")); err == nil {
			return pre
		} else if os.IsNotExist(err) {
			continue
		} else {
			panic(err)
		}
	}

	panic("Cant find python environment for training")
}

func init() {

	pythonPath = filepath.Join(prefix, "venv/bin/python")
	resourcePath = filepath.Join(prefix, "res/")
	configPath = filepath.Join(resourcePath, "configuration.in")
	weightsPath = filepath.Join(resourcePath, "weights.in")
	outputPath = filepath.Join(resourcePath, "output.out")
}

func EnableDebug(b bool) {
	logger.SetDebug(b)
}

func readUpdatesFromDisk() (output string, err error) {

	var content []byte

	content, err = ioutil.ReadFile(outputPath)
	if err != nil {
		return
	}

	logger.Debugf("Read %d bytes from %s", len(content), outputPath)
	output = string(content)

	return
}

func writeModelToDisk(configuration string, weights string) (err error) {

	err = ioutil.WriteFile(configPath, []byte(configuration), 0644)
	if err != nil {
		return
	}
	logger.Debugf("Wrote %d bytes as %s to disk.", len([]byte(configuration)), configPath)

	ioutil.WriteFile(weightsPath, []byte(weights), 0644)
	if err != nil {
		return
	}
	logger.Debugf("Wrote %d bytes as %s to disk.", len([]byte(weights)), weightsPath)

	return
}

func cleanUpRessources() {

	for _, res := range []string{configPath, weightsPath, outputPath} {
		err := os.Remove(res)

		if err != nil && !os.IsNotExist(err) {
			panic(fmt.Sprintf("Tried deleting %s after training but got %s", res, err))
		}
	}
	logger.Debug("Cleaned up resources from disk")
}

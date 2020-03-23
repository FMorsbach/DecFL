package tensorflow

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var prefix string
var pythonPath string
var resourcePath string
var configPath string
var weightsPath string
var outputPath string

type TensorflowError struct {
	Err          error
	Description  string
	PythonOutput string
}

func (e *TensorflowError) Error() string {
	return e.Description + " " + e.Err.Error() + " " + e.PythonOutput
}

func init() {

	defer func() {
		pythonPath = prefix + "venv/bin/python"
		resourcePath = prefix + "res/"
		configPath = resourcePath + "configuration.in"
		weightsPath = resourcePath + "weights.in"
		outputPath = resourcePath + "output.out"
	}()

	// Check if python is reachable without prefix
	if _, err := os.Stat("venv/bin/python"); err == nil {
		return
	} else if !os.IsNotExist(err) {
		panic(err)
	}

	// python is not found, check if prefix helps
	if _, err := os.Stat("training/tensorflow/venv/bin/python"); err == nil {
		prefix = "training/tensorflow/"
		return
	} else if os.IsNotExist(err) {
		panic("Cant find python environment for training")
	} else {
		panic(err)
	}
}

func readUpdatesFromDisk() (output string, err error) {

	var content []byte

	content, err = ioutil.ReadFile(outputPath)
	if err != nil {
		return
	}

	log.Printf("Read %d bytes from %s", len(content), outputPath)
	output = string(content)

	return
}

func writeModelToDisk(configuration string, weights string) (err error) {

	err = ioutil.WriteFile(configPath, []byte(configuration), 0644)
	if err != nil {
		return
	}
	log.Printf("Wrote %d bytes as %s to disk.", len([]byte(configuration)), configPath)

	ioutil.WriteFile(weightsPath, []byte(weights), 0644)
	if err != nil {
		return
	}
	log.Printf("Wrote %d bytes as %s to disk.", len([]byte(weights)), weightsPath)

	log.Println("Wrote model to disk")
	return
}

func cleanUpRessources() {

	for _, res := range []string{configPath, weightsPath, outputPath} {
		err := os.Remove(res)

		if err != nil && !os.IsNotExist(err) {
			panic(fmt.Sprintf("Tried deleting %s after training but got %s", res, err))
		}
	}
}

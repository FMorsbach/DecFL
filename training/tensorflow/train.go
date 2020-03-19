/*
This package provides an abstraction to tensorflow training which is carried
out in python rather than go.
*/
package tensorflow

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

var prefix string
var pythonPath string
var scriptPath string
var configPath string
var weightsPath string
var outputPath string

func init() {

	defer func() {
		pythonPath = prefix + "venv/bin/python"
		scriptPath = prefix + "train.py"
		configPath = prefix + "res/configuration.in"
		weightsPath = prefix + "res/weights.in"
		outputPath = prefix + "res/output.out"
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

type TrainError struct {
	Err         error
	Description string
	Details     string
}

func (e *TrainError) Error() string {
	return e.Description + " " + e.Err.Error() + " " + e.Details
}

// Trains the supplied model with local data and returns the updated weights
func Train(configuration string, weights string) (updatedWeights string, err error) {

	updatedWeights, err = trainByFile(configuration, weights)
	return
}

func trainByFile(configuration string, weights string) (updatedWeights string, err error) {

	err = writeModelToDisk(configuration, weights)
	if err != nil {
		return "", &TrainError{err, "Could not write model to disk", ""}
	}
	log.Println("Wrote model to disk")

	cmd := exec.Command(pythonPath, scriptPath, configPath, weightsPath, outputPath)

	log.Print("Executing: ", cmd.Args)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", &TrainError{err, "Could not run training script", string(out)}
	}
	log.Println("Training completed")

	updatedWeights, err = readModelUpdatesFromDisk()
	if err != nil {
		return "", &TrainError{err, "Could not read training results from disk", ""}
	}
	log.Println("Read model back from disk")

	cleanUpRessources()
	return
}

func readModelUpdatesFromDisk() (weights string, err error) {

	var content []byte

	content, err = ioutil.ReadFile(outputPath)
	if err != nil {
		return
	}

	log.Printf("Read %d bytes from %s", len(content), outputPath)
	weights = string(content)

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

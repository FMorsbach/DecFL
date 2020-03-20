/*
This package provides an abstraction to tensorflow training which is carried
out in python rather than go.
*/
package tensorflow

import (
	"log"
	"os/exec"
)

type TrainError struct {
	Err         error
	Description string
	Details     string
}

func (e *TrainError) Error() string {
	return e.Description + " " + e.Err.Error() + " " + e.Details
}

var trainScript string

func init() {
	trainScript = prefix + "train.py"
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

	cmd := exec.Command(pythonPath, trainScript, configPath, weightsPath, outputPath)

	log.Print("Executing: ", cmd.Args)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", &TrainError{err, "Could not run training script", string(out)}
	}
	log.Println("Training completed")

	updatedWeights, err = readOutputFromDisk()
	if err != nil {
		return "", &TrainError{err, "Could not read training results from disk", ""}
	}
	log.Println("Read model back from disk")

	cleanUpRessources()
	return
}

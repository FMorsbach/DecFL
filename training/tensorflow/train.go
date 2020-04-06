/*
This package provides an abstraction to tensorflow training which is carried
out in python rather than go.
*/
package tensorflow

import (
	"os/exec"
	"path/filepath"
)

var trainScript string

func init() {
	trainScript = filepath.Join(scriptsPath, "train.py")
}

// Trains the supplied model with local data and returns the updated weights
func (t *TensorflowFramework) Train(configuration string, weights string) (updatedWeights string, err error) {

	updatedWeights, err = trainByFile(configuration, weights)
	return
}

func trainByFile(configuration string, weights string) (updatedWeights string, err error) {

	defer cleanUpRessources()

	err = writeModelToDisk(configuration, weights)
	if err != nil {
		return "", &TensorflowError{err, "Could not write model to disk", ""}
	}
	logger.Debugln("Wrote model to disk")

	cmd := exec.Command(
		pythonPath,
		trainScript,
		filepath.Join(resourcePath, CONFIG_FILE),
		filepath.Join(resourcePath, WEIGHTS_FILE),
		filepath.Join(resourcePath, OUTPUT_FILE),
	)

	logger.Debugln("Executing: ", cmd.Args)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", &TensorflowError{err, "Could not run training script", string(out)}
	}
	logger.Debugln("Training completed")

	updatedWeights, err = readUpdatesFromDisk()
	if err != nil {
		return "", &TensorflowError{err, "Could not read training results from disk", ""}
	}
	logger.Debugln("Read model back from disk")

	return
}

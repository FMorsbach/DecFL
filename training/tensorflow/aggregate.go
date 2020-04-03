package tensorflow

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

var aggregateScript string

func init() {
	aggregateScript = filepath.Join(scriptsPath, "aggregate.py")
}

func (t *TensorflowTrainer) Aggregate(updates []string) (aggregatedWeights string, err error) {

	defer cleanUpRessources()

	for i, update := range updates {
		path := filepath.Join(resourcePath, strconv.Itoa(i)+"_trainingWeights.in")
		err := ioutil.WriteFile(path, []byte(update), 0644)
		if err != nil {
			return "", &TensorflowError{err, fmt.Sprintf("Can't write update %d to %s", i, path), ""}
		}
		logger.Debugf("Wrote update %d to %s", i, path)

		defer func() {
			err := os.Remove(path)

			if err != nil && !os.IsNotExist(err) {
				panic(fmt.Sprintf("Tried deleting %s after aggregation but got %s", path, err))
			}
		}()
	}

	outputPath := filepath.Join(resourcePath, OUTPUT_FILE)

	cmd := exec.Command(
		pythonPath,
		aggregateScript,
		resourcePath,
		outputPath,
	)

	defer func() {
		err := os.Remove(outputPath)
		if err != nil && !os.IsNotExist(err) {
			panic(fmt.Sprintf("Tried deleting %s after aggregation but got %s", outputPath, err))
		}
	}()

	logger.Debug("Executing: ", cmd.Args)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", &TensorflowError{err, "Could not run aggregation script", string(out)}
	}
	logger.Debug("Aggregation completed")

	aggregatedWeights, err = readUpdatesFromDisk()
	if err != nil {
		return "", &TensorflowError{err, "Could not read aggregated weights from disk", ""}
	}
	logger.Debug("Read aggregated weights back from disk")

	return
}

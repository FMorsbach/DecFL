package tensorflow

import (
	"encoding/json"
	"os/exec"
	"path/filepath"

	"github.com/FMorsbach/DecFL/training"
	"github.com/FMorsbach/dlog"
)

var evaluateScript string

func init() {
	evaluateScript = filepath.Join(prefix, "evaluate.py")
}

func Evaluate(configuration string, weights string) (results training.EvaluationResults, err error) {

	defer cleanUpRessources()

	err = writeModelToDisk(configuration, weights)
	if err != nil {
		return results, &TensorflowError{err, "Could not write model to disk", ""}
	}

	cmd := exec.Command(pythonPath, evaluateScript, configPath, weightsPath, outputPath)

	dlog.Debug("Executing: ", cmd.Args)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return results, &TensorflowError{err, "Could not run evaluation script", string(out)}
	}
	dlog.Debug("Evaluation complete")

	output, err := readUpdatesFromDisk()
	if err != nil {
		return results, &TensorflowError{err, "Could not read evaluation results from disk", ""}
	}

	results, err = parseOutput(output)
	if err != nil {
		return results, &TensorflowError{err, "Could not parse evaluation results", ""}
	}

	return
}

func parseOutput(input string) (result training.EvaluationResults, err error) {

	var i []interface{}

	err = json.Unmarshal([]byte(input), &i)

	if err != nil {
		return
	}

	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	loss := i[0].(float64)
	accuracy := i[1].(float64)

	result = training.EvaluationResults{loss, accuracy}
	return
}

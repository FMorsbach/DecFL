package tensorflow

import (
	"encoding/json"
	"log"
	"os/exec"
)

type EvaluationResults struct {
	Loss     float64
	Accuracy float64
}

type EvaluationError struct {
	Err         error
	Description string
	Details     string
}

func (e EvaluationError) Error() string {
	return e.Description + " " + e.Err.Error() + " " + e.Details
}

var evaluateScript string

func init() {
	evaluateScript = prefix + "evaluate.py"
}

func Evaluate(configuration string, weights string) (results EvaluationResults, err error) {

	err = writeModelToDisk(configuration, weights)
	if err != nil {
		return
	}

	cmd := exec.Command(pythonPath, evaluateScript, configPath, weightsPath, outputPath)

	log.Print("Executing: ", cmd.Args)
	out, err := cmd.CombinedOutput()
	if err != nil {
		_ = out
		return
	}

	output, err := readOutputFromDisk()
	if err != nil {
		return
	}

	results, err = parseOutput(output)
	if err != nil {
		return
	}

	return
}

func parseOutput(input string) (result EvaluationResults, err error) {

	var i []interface{}

	err = json.Unmarshal([]byte(input), &i)

	if err != nil {
		return
	}

	defer func() {
		if r := recover(); r != nil {
			err = EvaluationError{r.(error), "Could not parse output", ""}
		}
	}()
	loss := i[0].(float64)
	accuracy := i[1].(float64)

	result = EvaluationResults{loss, accuracy}
	return
}

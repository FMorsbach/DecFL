package tensorflow

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

var aggregateScript string

func init() {
	aggregateScript = filepath.Join(prefix, "aggregate.py")
}

func Aggregate(updates []string) (aggregatedWeights string, err error) {

	for i, update := range updates {
		path := filepath.Join(resourcePath, strconv.Itoa(i)+"_trainingWeights.in")
		err := ioutil.WriteFile(path, []byte(update), 0644)
		if err != nil {
			return "", &TensorflowError{err, fmt.Sprintf("Can't write update %d to %s", i, path), ""}
		}
		log.Printf("Wrote update %d to %s", i, path)

		defer func() {
			err := os.Remove(path)

			if err != nil && !os.IsNotExist(err) {
				panic(fmt.Sprintf("Tried deleting %s after aggregation but got %s", path, err))
			}
		}()
	}

	cmd := exec.Command(pythonPath, aggregateScript, resourcePath, outputPath)

	defer func() {
		err := os.Remove(outputPath)
		if err != nil && !os.IsNotExist(err) {
			panic(fmt.Sprintf("Tried deleting %s after aggregation but got %s", outputPath, err))
		}
	}()

	log.Print("Executing: ", cmd.Args)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", &TensorflowError{err, "Could not run aggregation script", string(out)}
	}
	log.Println("Aggregation completed")

	aggregatedWeights, err = readUpdatesFromDisk()
	if err != nil {
		return "", &TensorflowError{err, "Could not read aggregated weights from disk", ""}
	}
	log.Println("Read aggregated weights back from disk")

	return
}

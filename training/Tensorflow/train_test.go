package train

import (
	"io/ioutil"
	"testing"
)

const configuration string = "config1"
const weights string = "weights1"

func TestWriteToDisk(t *testing.T) {
	err := writeModelToDisk(configuration, weights)
	if err != nil {
		t.Error(err)
	}
}

func TestReadModelUpdatesFromDisk(t *testing.T) {

	err := ioutil.WriteFile("res/output.txt", []byte(weights), 0644)
	if err != nil {
		t.Error(err)
	}

	content, err := readModelUpdatesFromDisk()
	if err != nil {
		t.Error(err)
	}

	if content != weights {
		t.Errorf("Read %s as weights but wanted %s", content, weights)
	}

}

func TestTrainModel(t *testing.T) {

	content, err := ioutil.ReadFile("testData/configuration.txt")
	if err != nil {
		t.Error(err)
	}

	configuration := string(content)

	content, err = ioutil.ReadFile("testData/weights.txt")
	if err != nil {
		t.Error(err)
	}

	weights := string(content)

	updatedWeights := Run(configuration, weights)

	if updatedWeights == weights {
		t.Error("Weights didn't change")
	}

	if updatedWeights == "" {
		t.Error("Returned weights are empty")
	}

}

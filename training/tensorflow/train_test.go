package tensorflow

import (
	"testing"
)

func TestTrainModel(t *testing.T) {

	updatedWeights, err := Train(testConfiguration, testWeights)
	if err != nil {
		t.Error(err)
	}

	if updatedWeights == testWeights {
		t.Error("Weights didn't change")
	}

	if updatedWeights == "" {
		t.Error("Returned weights are empty")
	}

	cleanUpRessources()
}

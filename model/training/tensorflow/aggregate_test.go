package tensorflow

import (
	"io/ioutil"
	"testing"

	"github.com/FMorsbach/dlog"
)

func TestAggregate(t *testing.T) {

	inputWeights0, err := ioutil.ReadFile("testdata/0_trainingWeights.in")
	if err != nil {
		dlog.Fatal(err)
	}

	inputWeights1, err := ioutil.ReadFile("testdata/1_trainingWeights.in")
	if err != nil {
		dlog.Fatal(err)
	}

	inputWeights := []string{string(inputWeights0), string(inputWeights1)}

	result, err := trainer.Aggregate(inputWeights)
	if err != nil {
		t.Error((err))
	}

	outputWeights, err := ioutil.ReadFile("testdata/aggregation_output.out")

	if result != string(outputWeights) {
		t.Errorf("Aggregation results do not match expected weights from %s", "testdata/aggregation_output.out")
	}
}

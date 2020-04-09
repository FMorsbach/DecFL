package tensorflow

import (
	"testing"

	"github.com/FMorsbach/DecFL/model/training"
)

func TestEvaluate(t *testing.T) {

	expectedResult := training.EvaluationResults{Loss: 13.646389878750453, Accuracy: 0.9757}

	result, err := trainer.Evaluate(testConfiguration, testWeights)

	if err != nil {
		t.Error(err)
	}

	if result.Loss != expectedResult.Loss {
		t.Errorf("Loss mismatch: Expected %f but got %f", expectedResult.Loss, result.Loss)
	}

	if result.Accuracy != expectedResult.Accuracy {
		t.Errorf("Accuracy mismatch: Expected %f but got %f", expectedResult.Accuracy, result.Accuracy)
	}
}

func TestParseOutput(t *testing.T) {

	m := "[13.646389878750453, 0.9757]"
	expectedResult := training.EvaluationResults{Loss: 13.646389878750453, Accuracy: 0.9757}

	result, err := parseOutput(m)
	if err != nil {
		t.Error(err)
	}

	if result.Loss != expectedResult.Loss {
		t.Errorf("Loss mismatch: Expected %f but got %f", expectedResult.Loss, result.Loss)
	}

	if result.Accuracy != expectedResult.Accuracy {
		t.Errorf("Accuracy mismatch: Expected %f but got %f", expectedResult.Accuracy, result.Accuracy)
	}
}

// Malformed json
func TestParseOutput_Error(t *testing.T) {

	m := "[13.64638m9878750453, f0.9757]"

	_, err := parseOutput(m)
	if err == nil {
		t.Error("Expected error but returned nil as error")
	}
}

// Valid json, but incorrect data
func TestParseOutput_Error2(t *testing.T) {

	m := "[13.646389878750453]"

	_, err := parseOutput(m)
	if err == nil {
		t.Error("Expected error but returned nil as error")
	}
}

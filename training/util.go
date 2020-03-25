package training

import "fmt"

type EvaluationResults struct {
	Loss     float64
	Accuracy float64
}

func (e EvaluationResults) String() string {
	return fmt.Sprintf("Loss: %f - Accuracy %f", e.Loss, e.Accuracy)
}

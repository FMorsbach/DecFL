package training

import "fmt"

type EvaluationResults struct {
	Loss     float64
	Accuracy float64
}

func (e EvaluationResults) String() string {
	return fmt.Sprintf("Loss: %f - Accuracy %f", e.Loss, e.Accuracy)
}

type Trainer interface {
	Train(configuration string, weights string) (updatedWeights string, err error)
	Aggregate(updates []string) (aggregatedWeights string, err error)
	Evaluate(configuration string, weights string) (results EvaluationResults, err error)
}

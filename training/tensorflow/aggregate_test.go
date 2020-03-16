package tensorflow

import (
	"io/ioutil"
	"testing"
)

func TestSmokeTest(t *testing.T) {

	content1, err := ioutil.ReadFile("testData/aggregation/weights1.txt")
	if err != nil {
		t.Error(err)
	}

	content2, err := ioutil.ReadFile("testData/aggregation/weights2.txt")
	if err != nil {
		t.Error(err)
	}

	updates := []string{string(content1), string(content2)}

	result := Aggregate(updates)

	if result == "" {
		t.Error("Nil")
	}
}

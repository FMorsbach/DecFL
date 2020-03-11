package storage

import (
	"testing"
)

func TestSmokeTest(t *testing.T) {

	udpate := "SomeRandomUpdate"

	key, err := SaveLocalUpdate(udpate)
	if err != nil {
		t.Error(err)
	}

	weights, err := LoadGlobalState(key)
	if err != nil {
		t.Error(err)
	}

	if weights != udpate {
		t.Errorf("Write/Read failed. Expected %s but got %s", udpate, weights)
	}

}

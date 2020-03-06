package storage

import (
	"testing"
)

func TestSmokeTest(t *testing.T) {

	udpate := "SomeRandomUpdate"
	
	key, err := saveLocalUpdate(udpate)
	if err != nil {
		t.Error(err)
	}
	
	weights, err := loadGlobalState(key)
	if err != nil {
		t.Error(err)
	}

	if weights != udpate {
		t.Errorf("Write/Read failed. Expected %s but got %s", udpate, weights)
	}

}
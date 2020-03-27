package storage

import (
	"testing"

	"github.com/FMorsbach/DecFL/communication/chain"
)

func TestSmokeTest(t *testing.T) {

	udpate := "SomeRandomUpdate"

	key, err := StoreUpdate(udpate)
	if err != nil {
		t.Error(err)
	}

	weights, err := LoadGlobalState(chain.StorageAddress(key))
	if err != nil {
		t.Error(err)
	}

	if weights != udpate {
		t.Errorf("Write/Read failed. Expected %s but got %s", udpate, weights)
	}

}

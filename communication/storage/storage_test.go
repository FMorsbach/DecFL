package storage

import (
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/FMorsbach/DecFL/communication/mocks"
)

var content0 string = generateRandomContent()
var content1 string = generateRandomContent()

func generateRandomContent() string {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(rand.Int())
}

var testObjects []Storage

func init() {
	testObjects = []Storage{
		mocks.NewRedis(),
	}
}

func TestStoreAndLoad(t *testing.T) {

	for _, store := range testObjects {

		add, err := store.Store(content0)
		if err != nil {
			t.Error(err)
		}
		if add == "" {
			t.Error("Got empty address")
		}

		res, err := store.Load(add)
		if err != nil {
			t.Error(err)
		}
		if res != content0 {
			t.Errorf("Expected %s but got %s", content0, res)
		}
	}
}

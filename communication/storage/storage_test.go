package storage

import (
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/FMorsbach/DecFL/communication"
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
		NewRedis(),
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

func TestMultipleLoad(t *testing.T) {

	for _, store := range testObjects {

		add0, err := store.Store(content0)
		if err != nil {
			t.Error(err)
		}

		add1, err := store.Store(content1)
		if err != nil {
			t.Error(err)
		}

		con, err := store.Loads([]communication.StorageAddress{add0, add1})
		if err != nil {
			t.Error(err)
		}
		if len(con) == 2 {
			if con[0] != content0 {
				t.Errorf("Expected %s but got %s", content0, con[0])
			}
			if con[1] != content1 {
				t.Errorf("Expected %s but got %s", content1, con[1])
			}
		} else {
			t.Errorf("Expected result of length 2 but got %d", len(con))
		}
	}
}

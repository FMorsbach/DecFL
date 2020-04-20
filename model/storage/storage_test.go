package storage

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
)

var content0 string = generateRandomContent()
var content1 string = generateRandomContent()

func generateRandomContent() string {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(rand.Int())
}

func StoreAndLoad(store Storage, t *testing.T) {

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

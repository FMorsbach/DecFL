package storage

import "testing"

var store Storage

func init() {
	store = NewIPFS("localhost:5001")
}

func TestIPFSStoreAndLoad(t *testing.T) {
	StoreAndLoad(store, t)
}

package mocks

import (
	"math/rand"
	"strconv"
	"testing"
	"time"

	c "github.com/FMorsbach/DecFL/communication"
)

// TODO: check if can be imported
func FillAddress() c.StorageAddress {
	rand.Seed(time.Now().UnixNano())
	return c.StorageAddress(strconv.Itoa(rand.Int()))
}

func TestNewRedis(t *testing.T) {

	redis := NewRedis()

	if redis.client.Options().Addr != connection {
		t.Errorf("Expected %s but got %s", connection, redis.client.Options().Addr)
	}
}

func TestFlushRedis(t *testing.T) {

	redis := NewRedis()

	id, err := redis.DeployModel(FillAddress(), FillAddress())
	if err != nil {
		t.Error(err)
	}

	err = redis.FlushRedis()
	if err != nil {
		t.Error(err)
	}

	if ad, err := redis.ModelConfigurationAddress(id); err == nil {
		t.Errorf("Config: Got %s but expected error", ad)
	}

	if ad, err := redis.GlobalWeightsAddress(id); err == nil {
		t.Errorf("Weights: Got %s but expected error", ad)
	}
}

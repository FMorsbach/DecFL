package mocks

import (
	"math/rand"
	"strconv"
	"testing"
	"time"

	c "github.com/FMorsbach/DecFL/communication"
	"github.com/FMorsbach/dlog"
)

const connection string = "localhost:6379"

// TODO: check if can be imported
func FillAddress() c.StorageAddress {
	rand.Seed(time.Now().UnixNano())
	return c.StorageAddress(strconv.Itoa(rand.Int()))
}

func init() {
	if ok, err := NewRedis(connection).IsReachable(); !ok {
		dlog.Fatal("Cant reach redis: ", err)
	}
}

func TestNewRedis(t *testing.T) {

	redis := NewRedis(connection)

	if redis.client.Options().Addr != connection {
		t.Errorf("Expected %s but got %s", connection, redis.client.Options().Addr)
	}
}

func TestFlushRedis(t *testing.T) {

	redis := NewRedis(connection)

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

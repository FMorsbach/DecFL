package chain

import (
	"testing"
)

func TestNewRedis(t *testing.T) {

	redis := NewRedis()

	if redis.client.Options().Addr != connection {
		t.Errorf("Expected %s but got %s", connection, redis.client.Options().Addr)
	}
}

func TestFlushRedis(t *testing.T) {

	redis := NewRedis()

	id, err := redis.DeployModel(fillAddress(), fillAddress())
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

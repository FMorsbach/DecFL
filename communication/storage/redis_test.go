package storage

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

	add, err := redis.Store(generateRandomContent())
	if err != nil {
		t.Error(err)
	}

	err = redis.FlushRedis()
	if err != nil {
		t.Error(err)
	}

	if ad, err := redis.Load(add); err == nil {
		t.Errorf("Config: Got %s but expected error", ad)
	}
}

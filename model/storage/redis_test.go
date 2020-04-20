package storage

import (
	"testing"
)

var redis1 *Redis

func init() {

	connection := "localhost:6379"

	redis1 = NewRedis(connection, "")

	if redis1.client.Options().Addr != connection {
		logger.Fatalf("Expected %s but got %s", connection, redis1.client.Options().Addr)
	}

	if ok, err := redis1.IsReachable(); !ok {
		logger.Fatal("Cant reach redis: ", err)
	}
}

func TestRedisStoreAndLoad(t *testing.T) {
	StoreAndLoad(redis1, t)
}

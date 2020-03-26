package storage

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/FMorsbach/dlog"
	"github.com/go-redis/redis"
)

const connection string = "localhost:6379"
const modelID = "model1"

func StoreInitialModel(config string, weights string) (configKey string, weightsKey string) {

	client := redis.NewClient(&redis.Options{
		Addr:     connection,
		Password: "",
		DB:       0,
	})

	configKey = modelID + ":config"
	err := client.Set(configKey, config, 0).Err()
	if err != nil {
		logger.Fatal("Store initial config", err)
	}
	logger.Debugf("Stored config to %s", configKey)

	weightsKey = modelID + ":startWeights"
	err = client.Set(weightsKey, weights, 0).Err()
	if err != nil {
		logger.Fatal("Store initial weights", err)
	}
	logger.Debugf("Stored weights to %s", weightsKey)

	return
}

func LoadGlobalState(key string) (weights string, err error) {

	client := redis.NewClient(&redis.Options{
		Addr:     connection,
		Password: "",
		DB:       0,
	})

	weights, err = client.Get(key).Result()
	if err == redis.Nil {
		logger.Fatalf("Key %s does not exist\n", key)
		weights = ""
		return
	} else if err != nil {
		return
	}

	return
}

func StoreUpdate(weights string) (key string, err error) {

	client := redis.NewClient(&redis.Options{
		Addr:     connection,
		Password: "",
		DB:       0,
	})

	h := sha256.Sum256([]byte(weights))
	dh := h[0:32]
	key = hex.EncodeToString(dh)

	err = client.Set(key, weights, 0).Err()
	if err != nil {
		return
	}

	logger.Debugf("Saved update with key %s\n", key)

	err = nil
	return
}

func LocalUpdates(addresses []string) (updates []string) {

	client := redis.NewClient(&redis.Options{
		Addr:     connection,
		Password: "",
		DB:       0,
	})

	for _, address := range addresses {
		update, err := client.Get(address).Result()
		if err != nil {
			dlog.Fatal(err)
		}
		updates = append(updates, update)
	}
	return
}

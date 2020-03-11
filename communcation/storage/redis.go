package storage

import (
	"crypto/sha256"
	"encoding/hex"
	"log"

	"github.com/go-redis/redis/v7"
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
		log.Fatal("Store initial config", err)
	}

	weightsKey = modelID + ":startWeights"
	err = client.Set(weightsKey, weights, 0).Err()
	if err != nil {
		log.Fatal("Store initial weights", err)
	}
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
		log.Fatalf("Key %s does not exist\n", key)
		weights = ""
		return
	} else if err != nil {
		return
	}

	return
}

func SaveLocalUpdate(weights string) (key string, err error) {

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

	log.Printf("Saved update with key %s\n", key)

	err = nil
	return
}

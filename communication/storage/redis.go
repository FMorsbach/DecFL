package storage

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/FMorsbach/DecFL/communication"
	"github.com/FMorsbach/dlog"
	"github.com/go-redis/redis"
)

const connection string = "localhost:6379"
const modelID = "model1"

type Redis struct {
	client *redis.Client
}

func NewRedis() (instance *Redis) {

	client := redis.NewClient(&redis.Options{
		Addr:     connection,
		Password: "",
		DB:       0,
	})

	return &Redis{client: client}
}

func (r *Redis) Store(content string) (address communication.StorageAddress, err error) {

	h := sha256.Sum256([]byte(content))
	dh := h[0:32]
	address = communication.StorageAddress(hex.EncodeToString(dh))

	err = r.client.Set(string(address), content, 0).Err()
	if err != nil {
		return
	}
	logger.Debugf("Saved update with key %s\n", address)

	return
}

func (r *Redis) Load(address communication.StorageAddress) (content string, err error) {

	content, err = r.client.Get(string(address)).Result()
	if err == redis.Nil {
		err = fmt.Errorf("Key %s does not exist: %s", address, err)
		return
	} else if err != nil {
		return
	}

	return
}

func (r *Redis) Loads(addresses []communication.StorageAddress) (content []string, err error) {

	for _, address := range addresses {
		update, err := r.client.Get(string(address)).Result()
		if err != nil {
			return nil, err
		}
		content = append(content, update)
	}

	return
}

func (r *Redis) FlushRedis() (err error) {
	err = r.client.FlushAll().Err()
	if err != nil {
		logger.Fatal(err)
	}
	logger.Debugf("Flushed redis at %s", r.client.Options().Addr)

	return
}

func (r *Redis) IsReachable() (reachable bool, err error) {
	pong, err := r.client.Ping().Result()

	if err != nil || pong != "PONG" {
		return false, fmt.Errorf("%s, expected PONG but got %s", err, pong)
	} else {
		return true, err
	}
}

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

func LoadGlobalState(address communication.StorageAddress) (weights string, err error) {

	client := redis.NewClient(&redis.Options{
		Addr:     connection,
		Password: "",
		DB:       0,
	})

	weights, err = client.Get(string(address)).Result()
	if err == redis.Nil {
		logger.Fatalf("Key %s does not exist\n", address)
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

func LocalUpdates(addresses []communication.StorageAddress) (updates []string) {

	client := redis.NewClient(&redis.Options{
		Addr:     connection,
		Password: "",
		DB:       0,
	})

	for _, address := range addresses {
		update, err := client.Get(string(address)).Result()
		if err != nil {
			dlog.Fatal(err)
		}
		updates = append(updates, update)
	}
	return
}

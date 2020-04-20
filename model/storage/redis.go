package storage

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/FMorsbach/DecFL/model/common"
	"github.com/go-redis/redis"
)

const MODEL_CONFIG_KEY string = "globalModelConfiguration"
const MODEL_WEIGHTS_KEY string = "globalModelWeights"
const ITERATIONS_KEY string = "iteration"
const LOCAL_UPDATES_KEY string = "localUpdates"
const MODEL_EPOCH_KEY string = "epoch"

type Redis struct {
	client *redis.Client
}

func NewRedis(connection string, password string) (instance *Redis) {

	client := redis.NewClient(&redis.Options{
		Addr:     connection,
		Password: password,
		DB:       0,
	})

	return &Redis{client: client}
}

func (r *Redis) FlushRedis() (err error) {
	err = r.client.FlushAll().Err()
	if err != nil {
		return
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

func (r *Redis) Store(content string) (address common.StorageAddress, err error) {

	h := sha256.Sum256([]byte(content))
	dh := h[0:32]
	address = common.StorageAddress(hex.EncodeToString(dh))

	err = r.client.Set(string(address), content, 0).Err()
	if err != nil {
		return
	}
	logger.Debugf("Saved update with key %s\n", address)

	return
}

func (r *Redis) Load(address common.StorageAddress) (content string, err error) {

	content, err = r.client.Get(string(address)).Result()
	if err == redis.Nil {
		err = fmt.Errorf("Key %s does not exist: %s", address, err)
		return
	} else if err != nil {
		return
	}

	return
}

package storage

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/FMorsbach/DecFL/communication"
	"github.com/go-redis/redis"
)

const connection string = "localhost:6379"

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

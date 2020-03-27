package chain

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	c "github.com/FMorsbach/DecFL/communication"
	"github.com/go-redis/redis"
)

const connection string = "localhost:6379"

const MODEL_CONFIG_KEY string = "globalModelConfiguration"
const MODEL_WEIGHTS_KEY string = "globalModelWeights"
const ITERATIONS_KEY string = "iteration"
const LOCAL_UPDATES_KEY string = "localUpdates"

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

func (r *Redis) DeployModel(configAddress c.StorageAddress, weightsAddress c.StorageAddress) (id ModelIdentifier, err error) {

	rand.Seed(time.Now().UnixNano())
	id = ModelIdentifier(strconv.Itoa(rand.Intn(10000)))
	logger.Debugf("Generated %s as model id", id)

	err = r.client.Set(key(id, MODEL_CONFIG_KEY), string(configAddress), 0).Err()
	if err != nil {
		logger.Fatal(err)
	}
	logger.Debugf("Wrote configuration to %s", key(id, MODEL_CONFIG_KEY))

	err = r.client.Set(key(id, MODEL_WEIGHTS_KEY), string(weightsAddress), 0).Err()
	if err != nil {
		logger.Fatal(err)
	}
	logger.Debugf("Wrote weights to %s", key(id, MODEL_WEIGHTS_KEY))

	return
}

func (r *Redis) ModelConfigurationAddress(id ModelIdentifier) (address c.StorageAddress, err error) {

	temp, err := r.client.Get(key(id, MODEL_CONFIG_KEY)).Result()
	if err != nil {
		return
	}
	address = c.StorageAddress(temp)

	return
}

func (r *Redis) GlobalWeightsAddress(id ModelIdentifier) (address c.StorageAddress, err error) {

	temp, err := r.client.Get(key(id, MODEL_WEIGHTS_KEY)).Result()
	if err != nil {
		return
	}
	address = c.StorageAddress(temp)

	return
}

func (r *Redis) SetGlobalWeightsAddress(id ModelIdentifier, address c.StorageAddress) (err error) {

	err = r.client.Set(key(id, MODEL_WEIGHTS_KEY), string(address), 0).Err()
	if err != nil {
		return
	}
	logger.Debugf("Wrote weights to %s", key(id, MODEL_WEIGHTS_KEY))

	return
}

func (r *Redis) SubmitLocalUpdate(id ModelIdentifier, address c.StorageAddress) (err error) {

	err = r.client.SAdd(key(id, LOCAL_UPDATES_KEY), string(address)).Err()
	if err != nil {
		logger.Fatal(err)
	}
	logger.Debugf("Appended %s to %s", address, key(id, LOCAL_UPDATES_KEY))

	return
}

func (r *Redis) LocalUpdateAddresses(id ModelIdentifier) (addresses []c.StorageAddress, err error) {

	temp, err := r.client.SMembers(key(id, LOCAL_UPDATES_KEY)).Result()
	if err != nil {
		logger.Fatal(err)
	}

	addresses = make([]c.StorageAddress, len(temp))
	for i, t := range temp {
		addresses[i] = c.StorageAddress(t)
	}

	return
}

func (r *Redis) ClearLocalUpdateAddresses(id ModelIdentifier) (err error) {

	err = r.client.Del(key(id, LOCAL_UPDATES_KEY)).Err()
	if err != nil {
		logger.Fatal(err)
	}
	logger.Debugf("Reset local update store at %s", key(id, LOCAL_UPDATES_KEY))

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

func key(id ModelIdentifier, key string) string {
	return fmt.Sprintf("%s-%s", string(id), key)
}

package chain

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

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

func (r *Redis) DeployModel(configAddress storageAddress, weightsAddress storageAddress) (id modelIdentifier, err error) {

	rand.Seed(time.Now().UnixNano())
	id = modelIdentifier(strconv.Itoa(rand.Intn(10000)))
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

func (r *Redis) ModelConfigurationAddress(id modelIdentifier) (address storageAddress, err error) {

	temp, err := r.client.Get(key(id, MODEL_CONFIG_KEY)).Result()
	if err != nil {
		return
	}
	address = storageAddress(temp)

	return
}

func (r *Redis) GlobalWeightsAddress(id modelIdentifier) (address storageAddress, err error) {

	temp, err := r.client.Get(key(id, MODEL_WEIGHTS_KEY)).Result()
	if err != nil {
		return
	}
	address = storageAddress(temp)

	return
}

func (r *Redis) SetGlobalWeightsAddress(id modelIdentifier, address storageAddress) (err error) {

	err = r.client.Set(key(id, MODEL_WEIGHTS_KEY), string(address), 0).Err()
	if err != nil {
		return
	}
	logger.Debugf("Wrote weights to %s", key(id, MODEL_WEIGHTS_KEY))

	return
}

func (r *Redis) SubmitLocalUpdate(id modelIdentifier, address storageAddress) (err error) {

	err = r.client.SAdd(key(id, LOCAL_UPDATES_KEY), string(address)).Err()
	if err != nil {
		logger.Fatal(err)
	}
	logger.Debugf("Appended %s to %s", address, key(id, LOCAL_UPDATES_KEY))

	return
}

func (r *Redis) LocalUpdateAddresses(id modelIdentifier) (addresses []storageAddress, err error) {

	temp, err := r.client.SMembers(key(id, LOCAL_UPDATES_KEY)).Result()
	if err != nil {
		logger.Fatal(err)
	}

	addresses = make([]storageAddress, len(temp))
	for i, t := range temp {
		addresses[i] = storageAddress(t)
	}

	return
}

func (r *Redis) ClearLocalUpdateAddresses(id modelIdentifier) (err error) {

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

func DeployInitialModel(configuration string, weights string) {

	client := redis.NewClient(&redis.Options{
		Addr:     connection,
		Password: "",
		DB:       0,
	})

	err := client.Set(MODEL_CONFIG_KEY, configuration, 0).Err()
	if err != nil {
		logger.Fatal(err)
	}
	logger.Debugf("Wrote configuration to %s", MODEL_CONFIG_KEY)

	err = client.Set(MODEL_WEIGHTS_KEY, weights, 0).Err()
	if err != nil {
		logger.Fatal(err)
	}
	logger.Debugf("Wrote weights to %s", MODEL_WEIGHTS_KEY)

	return
}

func ModelConfigurationAddress() (configAddress string) {

	client := redis.NewClient(&redis.Options{
		Addr:     connection,
		Password: "",
		DB:       0,
	})

	configAddress, err := client.Get(MODEL_CONFIG_KEY).Result()
	if err != nil {
		logger.Fatal(err)
	}

	return
}

func GlobalWeightsAddress() (globalWeightsAddress string) {

	client := redis.NewClient(&redis.Options{
		Addr:     connection,
		Password: "",
		DB:       0,
	})

	globalWeightsAddress, err := client.Get(MODEL_WEIGHTS_KEY).Result()
	if err != nil {
		logger.Fatal(err)
	}

	return
}

func key(id modelIdentifier, key string) string {
	return fmt.Sprintf("%s-%s", string(id), key)
}

/***
*	Appends the address to the list of local updates
 */
func AppendUpdateAddress(id string, address string) {

	client := redis.NewClient(&redis.Options{
		Addr:     connection,
		Password: "",
		DB:       0,
	})

	err := client.SAdd(LOCAL_UPDATES_KEY, address).Err()
	if err != nil {
		logger.Fatal(err)
	}
	logger.Debugf("Appended %s to %s", address, LOCAL_UPDATES_KEY)

	return
}

func LocalUpdateAddresses() (addresses []string) {
	client := redis.NewClient(&redis.Options{
		Addr:     connection,
		Password: "",
		DB:       0,
	})

	addresses, err := client.SMembers(LOCAL_UPDATES_KEY).Result()
	if err != nil {
		logger.Fatal(err)
	}
	return
}

func SetGlobalWeightsAddress(address string) {
	client := redis.NewClient(&redis.Options{
		Addr:     connection,
		Password: "",
		DB:       0,
	})

	err := client.Set(MODEL_WEIGHTS_KEY, address, 0).Err()
	if err != nil {
		logger.Fatal(err)
	}
	logger.Debugf("Set %s to %s", MODEL_WEIGHTS_KEY, address)

	return
}

func CleanLocalUpdateStore() {
	client := redis.NewClient(&redis.Options{
		Addr:     connection,
		Password: "",
		DB:       0,
	})

	err := client.Del(LOCAL_UPDATES_KEY).Err()
	if err != nil {
		logger.Fatal(err)
	}

	logger.Debugf("Reset local update store at %s", LOCAL_UPDATES_KEY)
}

func FlushRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     connection,
		Password: "",
		DB:       0,
	})

	err := client.FlushAll().Err()
	if err != nil {
		logger.Fatal(err)
	}
}

func RedisIsReachable() bool {

	client := redis.NewClient(&redis.Options{
		Addr:     connection,
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	if err != nil || pong != "PONG" {
		return false
	} else {
		return true
	}
}

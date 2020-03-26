package chain

import (
	"github.com/go-redis/redis"
)

const connection string = "localhost:6379"

const globalModelConfigKey string = "globalModelConfiguration"
const globalModelWeightsKey string = "globalModelWeights"
const localUpdatesKey string = "localUpdates"
const iterationKey string = "iteration"
const modelIdentifierKey string = "modelIdentifier"

func DeployInitialModel(configuration string, weights string) {

	client := redis.NewClient(&redis.Options{
		Addr:     connection,
		Password: "",
		DB:       0,
	})

	err := client.Set(globalModelConfigKey, configuration, 0).Err()
	if err != nil {
		logger.Fatal(err)
	}
	logger.Debugf("Wrote configuration to %s", globalModelConfigKey)

	err = client.Set(globalModelWeightsKey, weights, 0).Err()
	if err != nil {
		logger.Fatal(err)
	}
	logger.Debugf("Wrote weights to %s", globalModelWeightsKey)

	return
}

func ModelConfigurationAddress() (configAddress string) {

	client := redis.NewClient(&redis.Options{
		Addr:     connection,
		Password: "",
		DB:       0,
	})

	configAddress, err := client.Get(globalModelConfigKey).Result()
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

	globalWeightsAddress, err := client.Get(globalModelWeightsKey).Result()
	if err != nil {
		logger.Fatal(err)
	}

	return
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

	err := client.SAdd(localUpdatesKey, address).Err()
	if err != nil {
		logger.Fatal(err)
	}
	logger.Debugf("Appended %s to %s", address, localUpdatesKey)

	return
}

func LocalUpdateAddresses() (addresses []string) {
	client := redis.NewClient(&redis.Options{
		Addr:     connection,
		Password: "",
		DB:       0,
	})

	addresses, err := client.SMembers(localUpdatesKey).Result()
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

	err := client.Set(globalModelWeightsKey, address, 0).Err()
	if err != nil {
		logger.Fatal(err)
	}
	logger.Debugf("Set %s to %s", globalModelWeightsKey, address)

	return
}

func CleanLocalUpdateStore() {
	client := redis.NewClient(&redis.Options{
		Addr:     connection,
		Password: "",
		DB:       0,
	})

	err := client.Del(localUpdatesKey).Err()
	if err != nil {
		logger.Fatal(err)
	}

	logger.Debugf("Reset local update store at %s", localUpdatesKey)
}

package chain

import (
	"github.com/FMorsbach/dlog"
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
		dlog.Fatal(err)
	}

	err = client.Set(globalModelWeightsKey, weights, 0).Err()
	if err != nil {
		dlog.Fatal(err)
	}

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
		dlog.Fatal(err)
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
		dlog.Fatal(err)
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
		dlog.Fatal(err)
	}

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
		dlog.Fatal(err)
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
		dlog.Fatal(err)
	}
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
		dlog.Fatal(err)
	}
}

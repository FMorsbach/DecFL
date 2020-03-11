package chain

import (
	"log"

	"github.com/go-redis/redis"
)

const connection string = "localhost:6379"
const globalModelConfigKey string = "globalModelConfiguration"
const globalModelWeightsKey string = "globalModelWeights"

func DeployInitialModel(configuration string, weights string) {

	client := redis.NewClient(&redis.Options{
		Addr:     connection,
		Password: "",
		DB:       0,
	})

	err := client.Set(globalModelConfigKey, configuration, 0).Err()
	if err != nil {
		log.Fatal(err)
	}

	err = client.Set(globalModelWeightsKey, weights, 0).Err()
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
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
		log.Fatal(err)
	}

	return
}

func WriteLocalTraingResult(key string, weights string) {

	client := redis.NewClient(&redis.Options{
		Addr:     connection,
		Password: "",
		DB:       0,
	})

	err := client.Set(key, weights, 0).Err()
	if err != nil {
		log.Fatal(err)
	}

	return
}

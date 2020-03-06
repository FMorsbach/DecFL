package storage

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/go-redis/redis/v7"
)

const connection string = "localhost:6379"

func main() {
	key, err := saveLocalUpdate("Tes32t")
	weights, err := loadGlobalState(key)
	fmt.Println(weights)
	_ = err
}

func loadGlobalState(key string) (weights string, err error) {

	client := redis.NewClient(&redis.Options{
		Addr:     connection,
		Password: "",
		DB:       0,
	})

	weights, err = client.Get(key).Result()
	if err == redis.Nil {
		fmt.Printf("Key %s does not exist\n", key)
		weights = ""
		return
	} else if err != nil {
		return
	}

	return

}

func saveLocalUpdate(weights string) (key string, err error) {

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

package chain

import (
	"testing"

	"github.com/go-redis/redis"
)

func TestDeploy(t *testing.T) {

	const testConfigString string = "testConfiguration"
	const testWeightsString string = "testWeights"

	DeployInitialModel(testConfigString, testWeightsString)

	rtnValue := ModelConfigurationAddress()
	if testConfigString != rtnValue {
		t.Errorf("Weights do not match. Expected %s but got %s", testConfigString, rtnValue)
	}

	rtnValue = GlobalWeightsAddress()
	if testWeightsString != rtnValue {
		t.Errorf("Weights do not match. Expected %s but got %s", testWeightsString, rtnValue)
	}

	return
}

func TestUpdate(t *testing.T) {

	const key string = "testLocalTrainingResult"
	const data string = "udpatedWeights"
	WriteLocalTraingResult(key, data)

	client := redis.NewClient(&redis.Options{
		Addr:     connection,
		Password: "",
		DB:       0,
	})

	rtnValue, err := client.Get(key).Result()
	if err != nil {
		t.Error(err)
	}

	if rtnValue != data {
		t.Errorf("Update mismatch: Expected %s but got %s", data, rtnValue)
	}

	return
}

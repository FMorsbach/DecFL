package mocks

import (
	"math/rand"
	"strconv"
	"testing"
	"time"

	c "github.com/FMorsbach/DecFL/communication"
	"github.com/FMorsbach/DecFL/communication/chain"
	"github.com/FMorsbach/DecFL/communication/chain/interface_tests"
	"github.com/FMorsbach/dlog"
)

const connection string = "localhost:6379"

var testAddress0 c.StorageAddress = FillAddress()
var testAddress1 c.StorageAddress = FillAddress()
var testAddress2 c.StorageAddress = FillAddress()
var trainerID c.TrainerIdentifier

var redis1 *Redis

func FillAddress() c.StorageAddress {
	rand.Seed(time.Now().UnixNano())
	return c.StorageAddress(strconv.Itoa(rand.Int()))
}

func init() {

	redis1 = NewRedis(connection)

	if ok, err := NewRedis(connection).IsReachable(); !ok {
		dlog.Fatal("Cant reach redis: ", err)
	}
}

func TestNewRedis(t *testing.T) {

	if redis1.client.Options().Addr != connection {
		t.Errorf("Expected %s but got %s", connection, redis1.client.Options().Addr)
	}
}

func TestFlushRedis(t *testing.T) {

	id, err := redis1.DeployModel(testAddress0, testAddress1, chain.Hyperparameters{UpdatesTillAggregation: 3})
	if err != nil {
		t.Error(err)
	}

	err = redis1.FlushRedis()
	if err != nil {
		t.Error(err)
	}

	if ad, err := redis1.ModelConfigurationAddress(id); err == nil {
		t.Errorf("Config: Got %s but expected error", ad)
	}

	if ad, err := redis1.GlobalWeightsAddress(id); err == nil {
		t.Errorf("Weights: Got %s but expected error", ad)
	}
}

func TestClearLocalUpdateAddresses(t *testing.T) {

	modelID, err := redis1.DeployModel(testAddress0, testAddress1, chain.Hyperparameters{UpdatesTillAggregation: 3})
	if err != nil {
		t.Error(err)
	}

	update := c.Update{
		Trainer: trainerID,
		Address: testAddress2,
	}

	err = redis1.SubmitLocalUpdate(modelID, update)
	if err != nil {
		t.Error(err)
	}

	addresses, err := redis1.LocalUpdates(modelID)
	if err != nil {
		t.Error(err)
	}

	count := len(addresses)
	if count == 0 {
		t.Error("Expected more than zero elements in list")
	}

	err = redis1.ClearLocalUpdateAddresses(modelID)
	if err != nil {
		t.Error(err)
	}

	if addresses, err = redis1.LocalUpdates(modelID); err != nil {
		t.Error(err)
	} else if len(addresses) != 0 {
		t.Errorf("Expected 0 elements in list but got %d", len(addresses))
	}
}

func TestDeployModelAndReadModel(t *testing.T) {
	interface_tests.DeployModelAndReadModel(redis1, t)
}

func TestLocalUpdateSubmission(t *testing.T) {
	interface_tests.LocalUpdateSubmission(redis1, trainerID, t)
}

func TestSubmitAggregationAndAggregation(t *testing.T) {
	interface_tests.SubmitAggregationAndAggregation(redis1, t)
}

func TestModelEpoch(t *testing.T) {
	interface_tests.ModelEpoch(redis1, t)
}

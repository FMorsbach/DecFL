package chain

import (
	"math/rand"
	"reflect"
	"strconv"
	"testing"
	"time"

	c "github.com/FMorsbach/DecFL/communication"
	"github.com/FMorsbach/DecFL/communication/mocks"
)

var testAddress0 c.StorageAddress = FillAddress()
var testAddress1 c.StorageAddress = FillAddress()
var testAddress2 c.StorageAddress = FillAddress()

func FillAddress() c.StorageAddress {
	rand.Seed(time.Now().UnixNano())
	return c.StorageAddress(strconv.Itoa(rand.Int()))
}

var testObjects []Chain

func init() {

	//logger.SetDebug(true)

	redis := mocks.NewRedis()

	if ok, err := redis.IsReachable(); !ok {
		logger.Fatal("Cant reach redis for testing", err)
	}

	if err := redis.FlushRedis(); err != nil {
		logger.Fatal(err)
	}

	testObjects = []Chain{
		redis,
	}
}

// TODO: Maybe add tests for wrong modelIds ?

func TestDeployModelAndReadModel(t *testing.T) {

	for _, chain := range testObjects {

		id, err := chain.DeployModel(testAddress0, testAddress1)
		if err != nil {
			t.Error(determineImplementation(chain), err)
		}

		if id == "" {
			t.Error(determineImplementation(chain), "Returned id is empty")
		}

		if key, err := chain.ModelConfigurationAddress(id); err != nil {
			t.Error(determineImplementation(chain), err)
		} else if key != testAddress0 {
			t.Errorf("%s Expected %s as ModelConfigurationAddress but got %s", determineImplementation(chain), testAddress0, key)
		}

		if key, err := chain.GlobalWeightsAddress(id); err != nil {
			t.Error(err)
		} else if key != testAddress1 {
			t.Errorf("%s Expected %s as GlobalWeightsAddress but got %s", determineImplementation(chain), testAddress1, key)
		}
	}

}

func TestSetGlobalWeightsAddress(t *testing.T) {

	for _, chain := range testObjects {

		id, err := chain.DeployModel(testAddress0, testAddress1)
		if err != nil {
			t.Error(determineImplementation(chain), err)
		}

		err = chain.SetGlobalWeightsAddress(id, testAddress0)
		if err != nil {
			t.Error(determineImplementation(chain), err)
		}

		if key, err := chain.GlobalWeightsAddress(id); err != nil {
			t.Error(err)
		} else if key != testAddress0 {
			t.Errorf("%s Expected %s as GlobalWeightsAddress but got %s", determineImplementation(chain), testAddress0, key)
		}
	}
}

func TestLocalUpdateSubmission(t *testing.T) {

	for _, chain := range testObjects {

		modelID, err := chain.DeployModel(testAddress0, testAddress1)
		if err != nil {
			t.Error(determineImplementation(chain), err)
		}

		updates, err := chain.LocalUpdates(modelID)
		if err != nil {
			t.Error(determineImplementation(chain), err)
		}

		count := len(updates)

		trainID := c.TrainerIdentifier("someTrainer")
		update := c.Update{
			Trainer: trainID,
			Address: testAddress2,
		}

		err = chain.SubmitLocalUpdate(modelID, update)
		if err != nil {
			t.Error(determineImplementation(chain), err)
		}

		if updates, err = chain.LocalUpdates(modelID); err != nil {
			t.Error(err)
		} else if len(updates) != count+1 {
			t.Errorf("%s Expected %d as number of elements but got %d", determineImplementation(chain), count+1, len(updates))
		} else if updates[count].Address != testAddress2 {
			t.Errorf("%s Expected %s as appended address but got %s", determineImplementation(chain), testAddress2, updates[count].Address)
		} else if updates[count].Trainer != trainID {
			t.Errorf("%s Expected %s as trainer ID but got %s", determineImplementation(chain), trainID, updates[count].Trainer)
		}
	}
}

func TestClearLocalUpdateAddresses(t *testing.T) {

	for _, chain := range testObjects {

		modelID, err := chain.DeployModel(testAddress0, testAddress1)
		if err != nil {
			t.Error(determineImplementation(chain), err)
		}

		trainID := c.TrainerIdentifier("someTrainer")
		update := c.Update{
			Trainer: trainID,
			Address: testAddress2,
		}

		err = chain.SubmitLocalUpdate(modelID, update)
		if err != nil {
			t.Error(determineImplementation(chain), err)
		}

		addresses, err := chain.LocalUpdates(modelID)
		if err != nil {
			t.Error(determineImplementation(chain), err)
		}

		count := len(addresses)
		if count == 0 {
			t.Error(determineImplementation(chain), "Expected more than zero elements in list")
		}

		err = chain.ClearLocalUpdateAddresses(modelID)
		if err != nil {
			t.Error(determineImplementation(chain), err)
		}

		if addresses, err = chain.LocalUpdates(modelID); err != nil {
			t.Error(err)
		} else if len(addresses) != 0 {
			t.Errorf("%s Expected 0 elements in list but got %d", determineImplementation(chain), len(addresses))
		}
	}
}

func determineImplementation(object Chain) string {
	switch t := object.(type) {
	case *mocks.Redis:
		return reflect.TypeOf(t).String()
	}

	panic("Can't determine type of object")
}

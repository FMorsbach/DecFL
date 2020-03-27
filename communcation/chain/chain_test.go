package chain

import (
	"math/rand"
	"reflect"
	"strconv"
	"testing"
	"time"
)

var testAddress0 storageAddress = fillAddress()
var testAddress1 storageAddress = fillAddress()
var testAddress2 storageAddress = fillAddress()

func fillAddress() storageAddress {
	rand.Seed(time.Now().UnixNano())
	return storageAddress(strconv.Itoa(rand.Int()))
}

var testObjects []Chain

func init() {

	//logger.SetDebug(true)

	redis := NewRedis()

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

		id, err := chain.DeployModel(testAddress0, testAddress1)
		if err != nil {
			t.Error(determineImplementation(chain), err)
		}

		addresses, err := chain.LocalUpdateAddresses(id)
		if err != nil {
			t.Error(determineImplementation(chain), err)
		}

		count := len(addresses)

		err = chain.SubmitLocalUpdate(id, testAddress2)
		if err != nil {
			t.Error(determineImplementation(chain), err)
		}

		if addresses, err = chain.LocalUpdateAddresses(id); err != nil {
			t.Error(err)
		} else if len(addresses) != count+1 {
			t.Errorf("%s Expected %d as number of elements but got %d", determineImplementation(chain), count+1, len(addresses))
		} else if addresses[count] != testAddress2 {
			t.Errorf("%s Expected %s as appended address but got %s", determineImplementation(chain), testAddress2, addresses[count])
		}
	}
}

func TestClearLocalUpdateAddresses(t *testing.T) {

	for _, chain := range testObjects {

		modelID, err := chain.DeployModel(testAddress0, testAddress1)
		if err != nil {
			t.Error(determineImplementation(chain), err)
		}

		err = chain.SubmitLocalUpdate(modelID, testAddress2)
		if err != nil {
			t.Error(determineImplementation(chain), err)
		}

		addresses, err := chain.LocalUpdateAddresses(modelID)
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

		if addresses, err = chain.LocalUpdateAddresses(modelID); err != nil {
			t.Error(err)
		} else if len(addresses) != 0 {
			t.Errorf("%s Expected 0 elements in list but got %d", determineImplementation(chain), len(addresses))
		}
	}
}

func determineImplementation(object Chain) string {
	switch t := object.(type) {
	case *Redis:
		return reflect.TypeOf(t).String()
	}

	panic("Can't determine type of object")
}

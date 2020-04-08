package interface_tests

import (
	"math/rand"
	"strconv"
	"testing"
	"time"

	c "github.com/FMorsbach/DecFL/communication"
	chainPkg "github.com/FMorsbach/DecFL/communication/chain"
)

var testConfigAddress c.StorageAddress = FillAddress()
var testWeightsAddress c.StorageAddress = FillAddress()

func FillAddress() c.StorageAddress {
	rand.Seed(time.Now().UnixNano())
	return c.StorageAddress(strconv.Itoa(rand.Int()))
}

// TODO: Maybe add tests for wrong modelIds ?

func DeployModelAndReadModel(chain chainPkg.Chain, t *testing.T) {

	id, err := chain.DeployModel(
		testConfigAddress,
		testWeightsAddress,
		chainPkg.Hyperparameters{UpdatesTillAggregation: 3})
	if err != nil {
		t.Error(err)
	}

	if id == "" {
		t.Error("Returned id is empty")
	}

	if key, err := chain.ModelConfigurationAddress(id); err != nil {
		t.Error(err)
	} else if key != testConfigAddress {
		t.Errorf("Expected %s as ModelConfigurationAddress but got %s", testConfigAddress, key)
	}

	if key, err := chain.GlobalWeightsAddress(id); err != nil {
		t.Error(err)
	} else if key != testWeightsAddress {
		t.Errorf("Expected %s as GlobalWeightsAddress but got %s", testWeightsAddress, key)
	}
}

func LocalUpdateSubmission(chain chainPkg.Chain, trainerID c.TrainerIdentifier, t *testing.T) {

	modelID, err := chain.DeployModel(
		testConfigAddress,
		testWeightsAddress,
		chainPkg.Hyperparameters{UpdatesTillAggregation: 3})
	if err != nil {
		t.Error(err)
	}

	updates, err := chain.LocalUpdates(modelID)
	if err != nil {
		t.Error(err)
	}

	count := len(updates)

	randomTestAddress1 := FillAddress()
	update := c.Update{
		Trainer: trainerID,
		Address: randomTestAddress1,
	}

	err = chain.SubmitLocalUpdate(modelID, update)
	if err != nil {
		t.Error(err)
	}

	if updates, err = chain.LocalUpdates(modelID); err != nil {
		t.Error(err)
	} else if len(updates) != count+1 {
		t.Errorf("Expected %d as number of elements but got %d", count+1, len(updates))
	} else if updates[count].Address != randomTestAddress1 {
		t.Errorf("Expected %s as appended address but got %s", randomTestAddress1, updates[count].Address)
	} else if updates[count].Trainer != trainerID {
		t.Errorf("Expected %s as trainer ID but got %s", trainerID, updates[count].Trainer)
	}
}

func SubmitAggregationAndAggregation(chain chainPkg.Chain, t *testing.T) {

	randomTestAddress1 := FillAddress()
	randomTestAddress2 := FillAddress()

	testCases := []struct {
		updates  []c.StorageAddress
		expected int
	}{
		{[]c.StorageAddress{
			randomTestAddress1,
			randomTestAddress1,
			randomTestAddress1,
		}, 0},
		{[]c.StorageAddress{
			randomTestAddress1,
			randomTestAddress1,
			randomTestAddress2,
		}, 0},
		{[]c.StorageAddress{
			randomTestAddress1,
			randomTestAddress2,
			randomTestAddress2,
		}, 1},
	}

	for i, testCase := range testCases {

		id, err := chain.DeployModel(
			testConfigAddress,
			testWeightsAddress,
			chainPkg.Hyperparameters{UpdatesTillAggregation: len(testCase.updates)},
		)
		if err != nil {
			t.Error(err)
		}

		for _, update := range testCase.updates {

			err = chain.SubmitAggregation(id, update)
			if err != nil {
				t.Error(err)
			}
		}

		if key, err := chain.GlobalWeightsAddress(id); err != nil {
			t.Error(err)
		} else if key != testCase.updates[testCase.expected] {
			t.Errorf("Case: %d Expected %s as GlobalWeightsAddress but got %s", i, testCase.updates[testCase.expected], key)
		}
	}
}

func ModelEpoch(chain chainPkg.Chain, t *testing.T) {

	id, err := chain.DeployModel(testConfigAddress, testWeightsAddress, chainPkg.Hyperparameters{UpdatesTillAggregation: 3})
	if err != nil {
		t.Error(err)
	}

	epoch, err := chain.ModelEpoch(id)
	if err != nil {
		t.Error(err)
	} else if epoch != 0 {
		t.Errorf("Expected epoch to be zero but got %d", epoch)
	}

	randomTestAddress1 := FillAddress()
	for i := 0; i < 3; i++ {

		err = chain.SubmitAggregation(id, randomTestAddress1)
		if err != nil {
			t.Error(err)
		}
	}

	epoch, err = chain.ModelEpoch(id)
	if err != nil {
		t.Error(err)
	} else if epoch != 1 {
		t.Errorf("Expected epoch to be 1 but got %d", epoch)
	}
}

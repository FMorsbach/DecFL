package interface_tests

import (
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/FMorsbach/DecFL/model/chain"
	"github.com/FMorsbach/DecFL/model/common"
)

var testConfigAddress common.StorageAddress = FillAddress()
var testWeightsAddress common.StorageAddress = FillAddress()

func FillAddress() common.StorageAddress {
	rand.Seed(time.Now().UnixNano())
	return common.StorageAddress(strconv.Itoa(rand.Int()))
}

// TODO: Maybe add tests for wrong modelIds ?

func DeployModelAndReadModel(chain chain.Chain, t *testing.T) {

	id, err := chain.DeployModel(
		testConfigAddress,
		testWeightsAddress,
		common.Hyperparameters{UpdatesTillAggregation: 3})
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

func LocalUpdateSubmission(chain chain.Chain, trainerID common.TrainerIdentifier, t *testing.T) {

	modelID, err := chain.DeployModel(
		testConfigAddress,
		testWeightsAddress,
		common.Hyperparameters{UpdatesTillAggregation: 3})
	if err != nil {
		t.Error(err)
	}

	updates, err := chain.LocalUpdates(modelID)
	if err != nil {
		t.Error(err)
	}

	count := len(updates)

	randomTestAddress1 := FillAddress()
	update := common.Update{
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

func SubmitAggregationAndAggregation(chain chain.Chain, t *testing.T) {

	randomTestAddress1 := FillAddress()
	randomTestAddress2 := FillAddress()

	testCases := []struct {
		updates  []common.StorageAddress
		expected int
	}{
		{[]common.StorageAddress{
			randomTestAddress1,
			randomTestAddress1,
			randomTestAddress1,
		}, 0},
		{[]common.StorageAddress{
			randomTestAddress1,
			randomTestAddress1,
			randomTestAddress2,
		}, 0},
		{[]common.StorageAddress{
			randomTestAddress1,
			randomTestAddress2,
			randomTestAddress2,
		}, 1},
		{[]common.StorageAddress{
			randomTestAddress1,
			randomTestAddress2,
			randomTestAddress1,
			randomTestAddress2,
			randomTestAddress2,
		}, 1},
	}

	for i, testCase := range testCases {

		id, err := chain.DeployModel(
			testConfigAddress,
			testWeightsAddress,
			common.Hyperparameters{UpdatesTillAggregation: len(testCase.updates)},
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

func ModelEpochAndMultipleSuccedingAggregations(chain chain.Chain, t *testing.T) {

	updatesTillAggregation := 3
	id, err := chain.DeployModel(
		testConfigAddress,
		testWeightsAddress,
		common.Hyperparameters{UpdatesTillAggregation: updatesTillAggregation},
	)
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
	for i := 0; i < updatesTillAggregation; i++ {

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

	randomTestAddress2 := FillAddress()
	for i := 0; i < updatesTillAggregation; i++ {

		err = chain.SubmitAggregation(id, randomTestAddress2)
		if err != nil {
			t.Error(err)
		}
	}

	epoch, err = chain.ModelEpoch(id)
	if err != nil {
		t.Error(err)
	} else if epoch != 2 {
		t.Errorf("Expected epoch to be 2 but got %d", epoch)
	}

	if key, err := chain.GlobalWeightsAddress(id); err != nil {
		t.Error(err)
	} else if key != randomTestAddress2 {
		t.Errorf("Expected %s as GlobalWeightsAddress but got %s", randomTestAddress2, key)
	}
}

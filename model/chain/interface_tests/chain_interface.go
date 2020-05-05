package interface_tests

import (
	"math/rand"
	"strconv"
	"testing"
	"time"

	ch "github.com/FMorsbach/DecFL/model/chain"
	"github.com/FMorsbach/DecFL/model/common"
)

var testConfigAddress common.StorageAddress = FillAddress()
var testWeightsAddress common.StorageAddress = FillAddress()
var testScriptsAddress common.StorageAddress = FillAddress()

func FillAddress() common.StorageAddress {
	rand.Seed(time.Now().UnixNano())
	return common.StorageAddress(strconv.Itoa(rand.Int()))
}

// TODO: Maybe add tests for wrong modelIds ?

func DeployModelAndReadModel(chain ch.Chain, t *testing.T) {

	id, err := chain.DeployModel(
		testConfigAddress,
		testWeightsAddress,
		testScriptsAddress,
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

	if key, err := chain.ScriptsAddress(id); err != nil {
		t.Error(err)
	} else if key != testScriptsAddress {
		t.Errorf("Expected %s as GlobalWeightsAddress but got %s", testScriptsAddress, key)
	}
}

func LocalUpdateSubmission(chain ch.Chain, trainerID common.TrainerIdentifier, t *testing.T) {

	modelID, err := chain.DeployModel(
		testConfigAddress,
		testWeightsAddress,
		testScriptsAddress,
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

	err = chain.SubmitLocalUpdate(modelID, randomTestAddress1)
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

func SubmitAggregationAndAggregation(chain ch.Chain, t *testing.T) {

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
			testScriptsAddress,
			common.Hyperparameters{
				UpdatesTillAggregation: len(testCase.updates),
				Epochs:                 3,
			},
		)
		if err != nil {
			t.Error(err)
		}

		for _, update := range testCase.updates {

			err = chain.SubmitLocalUpdate(id, update)
			if err != nil {
				t.Error(err)
			}
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

func ModelEpochAndMultipleSuccedingAggregations(chain ch.Chain, t *testing.T) {

	updatesTillAggregation := 3
	id, err := chain.DeployModel(
		testConfigAddress,
		testWeightsAddress,
		testScriptsAddress,
		common.Hyperparameters{
			UpdatesTillAggregation: updatesTillAggregation,
			Epochs:                 2,
		},
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

		err = chain.SubmitLocalUpdate(id, randomTestAddress1)
		if err != nil {
			t.Error(err)
		}
	}
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

		err = chain.SubmitLocalUpdate(id, randomTestAddress2)
		if err != nil {
			t.Error(err)
		}
	}
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

func StateTransitions(chain ch.Chain, t *testing.T) {

	randomTestAddress1 := FillAddress()

	testCases := []struct {
		updates []common.StorageAddress
	}{
		{[]common.StorageAddress{
			randomTestAddress1,
		}},
		{[]common.StorageAddress{
			randomTestAddress1,
			randomTestAddress1,
			randomTestAddress1,
		}},
		{[]common.StorageAddress{
			randomTestAddress1,
			randomTestAddress1,
			randomTestAddress1,
			randomTestAddress1,
			randomTestAddress1,
		}},
	}

	for i, testCase := range testCases {

		id, err := chain.DeployModel(
			testConfigAddress,
			testWeightsAddress,
			testScriptsAddress,
			common.Hyperparameters{
				UpdatesTillAggregation: len(testCase.updates),
				Epochs:                 2,
			},
		)
		if err != nil {
			t.Error(err)
		}

		for j, update := range testCase.updates {

			state, err := chain.State(id)
			if err != nil {
				t.Error(err)
			} else if state != common.Training {
				t.Errorf("Case: %d State was %d after %d updates instead of after %d", i, state, j, len(testCase.updates))
			}

			err = chain.SubmitLocalUpdate(id, update)
			if err != nil {
				t.Error(err)
			}
		}

		for j, update := range testCase.updates {

			state, err := chain.State(id)
			if err != nil {
				t.Error(err)
			} else if state != common.Aggregation {
				t.Errorf("Case: %d State was %d after %d updates and %d aggregations", i, state, len(testCase.updates), j)
			}

			err = chain.SubmitAggregation(id, update)
			if err != nil {
				t.Error(err)
			}
		}

		state, err := chain.State(id)
		if err != nil {
			t.Error(err)
		} else if state != common.Training {
			t.Errorf("Case: %d Did not reset after %d aggregations", i, len(testCase.updates))
		}

		for j, update := range testCase.updates {

			state, err := chain.State(id)
			if err != nil {
				t.Error(err)
			} else if state != common.Training {
				t.Errorf("Case: %d State was %d after %d updates instead of after %d", i, state, j, len(testCase.updates))
			}

			err = chain.SubmitLocalUpdate(id, update)
			if err != nil {
				t.Error(err)
			}
		}

		for j, update := range testCase.updates {

			state, err := chain.State(id)
			if err != nil {
				t.Error(err)
			} else if state != common.Aggregation {
				t.Errorf("Case: %d State was %d after %d updates and %d aggregations", i, state, len(testCase.updates), j)
			}

			err = chain.SubmitAggregation(id, update)
			if err != nil {
				t.Error(err)
			}
		}

		state, err = chain.State(id)
		if err != nil {
			t.Error(err)
		} else if state != common.Finished {
			t.Errorf("Case: %d Expected state %d but got %d", i, common.Finished, state)
		}

	}
}

func ResetLocalUpdatesAfterAggregation(chain ch.Chain, t *testing.T) {

	updatesTillAggregation := 2
	modelID, err := chain.DeployModel(
		testConfigAddress,
		testWeightsAddress,
		testScriptsAddress,
		common.Hyperparameters{UpdatesTillAggregation: updatesTillAggregation})
	if err != nil {
		t.Error(err)
	}

	randomTestAddress1 := FillAddress()
	for i := 0; i < updatesTillAggregation; i++ {

		err = chain.SubmitLocalUpdate(modelID, randomTestAddress1)
		if err != nil {
			t.Error(err)
		}
	}

	updates, err := chain.LocalUpdates(modelID)
	if err != nil {
		t.Error(err)
	}

	if len(updates) != updatesTillAggregation {
		t.Errorf("Expected %d stored updates but got %d", updatesTillAggregation, len(updates))
	}

	for i := 0; i < updatesTillAggregation; i++ {

		err = chain.SubmitAggregation(modelID, randomTestAddress1)
		if err != nil {
			t.Error(err)
		}
	}

	updates, err = chain.LocalUpdates(modelID)
	if err != nil {
		t.Error(err)
	}

	if len(updates) != 0 {
		t.Errorf("Expected 0 stored updates but got %d", len(updates))
	}
}

func Authorization(chain1 ch.Chain, chain2 ch.Chain, trainerID2 common.TrainerIdentifier, t *testing.T) {

	testCases := []struct {
		client    ch.Chain
		tID       common.TrainerIdentifier
		accepted  bool
		addClient bool
	}{
		{chain1, common.TrainerIdentifier{}, true, false},
		{chain2, trainerID2, false, false},
		{chain2, trainerID2, true, true},
	}

	for i, test := range testCases {

		id, err := chain1.DeployModel(
			testConfigAddress,
			testWeightsAddress,
			testScriptsAddress,
			common.Hyperparameters{
				UpdatesTillAggregation: 2,
				Epochs:                 2,
			},
		)
		if err != nil {
			t.Error(err)
		}

		before, err := chain1.LocalUpdates(id)
		if err != nil {
			t.Error(err)
		}
		if len(before) != 0 {
			t.Errorf("Case %d: Expected 0 updates but got %d", i, len(before))
		}

		if test.addClient {
			err := chain1.AddTrainer(id, test.tID)
			if err != nil {
				t.Error(err)
			}
		}

		err = test.client.SubmitLocalUpdate(id, testWeightsAddress)
		if !test.accepted && err.Error() == "VM Exception while processing transaction: revert Not an authorized trainer" {

		} else if err != nil {
			t.Errorf("Case %d: %s", i, err)
		}

		after, err := chain1.LocalUpdates(id)
		if err != nil {
			t.Error(err)
		}
		if test.accepted {
			if len(after) != len(before)+1 {
				t.Errorf("Case %d: Expected %d updates but got %d", i, len(before)+1, len(after))
			}
		} else {
			if len(after) != len(before) {
				t.Errorf("Case %d: Expected %d updates but got %d", i, len(before), len(after))
			}
		}
	}
}

func StateRejection(chain ch.Chain, t *testing.T) {

	id, err := chain.DeployModel(
		testConfigAddress,
		testWeightsAddress,
		testScriptsAddress,
		common.Hyperparameters{
			UpdatesTillAggregation: 1,
			Epochs:                 2,
		},
	)
	if err != nil {
		t.Error(err)
	}

	errorMsg := "VM Exception while processing transaction: revert Not valid at this state"

	err = chain.SubmitAggregation(id, testWeightsAddress)
	if err.Error() != errorMsg {
		t.Errorf("Instead got '%s'", err.Error())
	}

	err = chain.SubmitLocalUpdate(id, testWeightsAddress)
	if err != nil {
		t.Error(err)
	}

	err = chain.SubmitLocalUpdate(id, testWeightsAddress)
	if err.Error() != errorMsg {
		t.Errorf("Instead got '%s'", err.Error())
	}

}

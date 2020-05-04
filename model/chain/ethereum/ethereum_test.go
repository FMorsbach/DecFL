package ethereum

import (
	"crypto/ecdsa"
	"fmt"
	"testing"

	"github.com/FMorsbach/DecFL/model/chain"
	"github.com/FMorsbach/DecFL/model/chain/interface_tests"
	"github.com/FMorsbach/DecFL/model/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var ethChain chain.Chain
var ethChain2 chain.Chain
var trainerID common.TrainerIdentifier

func init() {

	key := "3b3a098805d048bab52b82b8767da2117af104cc97ec820acbe1b63e768ebba7"
	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		return
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		err = fmt.Errorf("%s%s", "Cannot assert type", err)
		return
	}
	trainerID = common.TrainerIdentifier(crypto.PubkeyToAddress(*publicKeyECDSA))

	ethChain, err = NewEthereum(
		"http://localhost:8545",
		key,
	)

	ethChain2, err = NewEthereum(
		"http://localhost:8545",
		"ea7e514344e7495c942cf559cbff1954abbac43f45df04608ded63ab5af82f5c",
	)

	if err != nil {
		logger.Fatal(err)
	}

}

func TestDeployModelAndReadModel(t *testing.T) {
	interface_tests.DeployModelAndReadModel(ethChain, t)
}

func TestLocalUpdateSubmission(t *testing.T) {
	interface_tests.LocalUpdateSubmission(ethChain, trainerID, t)
}

func TestSubmitAggregationAndAggregation(t *testing.T) {
	interface_tests.SubmitAggregationAndAggregation(ethChain, t)
}

func TestModelEpoch(t *testing.T) {
	interface_tests.ModelEpochAndMultipleSuccedingAggregations(ethChain, t)
}

func TestStateTransitions(t *testing.T) {
	interface_tests.StateTransitions(ethChain, t)
}

func TestResetLocalUpdatesAfterAggregation(t *testing.T) {
	interface_tests.ResetLocalUpdatesAfterAggregation(ethChain, t)
}

func TestAuthorization(t *testing.T) {
	interface_tests.Authorization(ethChain2, ethChain, trainerID, t)
}

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

	key := "47890574e4835892e0de116ccc144d134027ceba239789c84e58b47fde50e6ca"
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
		"9887d8fa8f3d50d2d044beaf72e1fa15d23c9b79fa3cd81426a35cd424106100",
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

func TestStateRejection(t *testing.T) {
	interface_tests.StateRejection(ethChain, t)
}

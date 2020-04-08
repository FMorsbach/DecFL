package ethereum

import (
	"crypto/ecdsa"
	"fmt"
	"testing"

	c "github.com/FMorsbach/DecFL/communication"
	com "github.com/FMorsbach/DecFL/communication"
	"github.com/FMorsbach/DecFL/communication/chain"
	"github.com/FMorsbach/DecFL/communication/chain/interface_tests"
	"github.com/ethereum/go-ethereum/crypto"
)

var ethChain chain.Chain
var trainerID c.TrainerIdentifier

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
	trainerID = com.TrainerIdentifier(crypto.PubkeyToAddress(*publicKeyECDSA))

	ethChain, err = NewEthereum(
		"http://localhost:8545",
		key,
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
	interface_tests.ModelEpoch(ethChain, t)
}

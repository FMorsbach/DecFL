package ethereum

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"

	com "github.com/FMorsbach/DecFL/communication"
	"github.com/FMorsbach/DecFL/communication/chain"
	"github.com/FMorsbach/DecFL/communication/chain/ethereum/contract"
	"github.com/FMorsbach/dlog"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var logger = dlog.New(os.Stderr, "Chain: ", log.LstdFlags, false)

func EnableDebug(b bool) {
	logger.SetDebug(b)
}

type ethereumChain struct {
	client        ethclient.Client
	privateKey    ecdsa.PrivateKey
	publicKey     ecdsa.PublicKey
	publicAddress common.Address
}

func NewEthereum(chainAddress string, key string) (instance chain.Chain, err error) {

	client, err := ethclient.Dial(chainAddress)
	if err != nil {
		return
	}

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
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	return &ethereumChain{
		client:        *client,
		privateKey:    *privateKey,
		publicKey:     *publicKeyECDSA,
		publicAddress: fromAddress,
	}, nil
}

func (c *ethereumChain) DeployModel(configAddress com.StorageAddress, weightsAddress com.StorageAddress, params chain.Hyperparameters) (id com.ModelIdentifier, err error) {

	nonce, err := c.client.PendingNonceAt(context.Background(), c.publicAddress)
	if err != nil {
		return

	}

	gasPrice, err := c.client.SuggestGasPrice(context.Background())
	if err != nil {
		return
	}

	auth := bind.NewKeyedTransactor(&(c.privateKey))
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	address, tx, _, err := contract.DeployContract(
		auth,
		&(c.client),
		string(configAddress),
		string(weightsAddress),
		big.NewInt(int64(params.UpdatesTillAggregation)),
	)
	if err != nil {
		return
	}

	id = com.ModelIdentifier(address.Hex())
	logger.Debugf("Deployed contract in transaction %s", tx.Hash().Hex())

	return
}

func (c *ethereumChain) ModelConfigurationAddress(id com.ModelIdentifier) (address com.StorageAddress, err error) {

	instance, err := contract.NewContract(common.HexToAddress(string(id)), &(c.client))
	if err != nil {
		return
	}

	value, err := instance.ConfigurationAddress(nil)
	if err != nil {
		return
	}

	address = com.StorageAddress(value)
	return
}

func (c *ethereumChain) GlobalWeightsAddress(id com.ModelIdentifier) (address com.StorageAddress, err error) {

	instance, err := contract.NewContract(common.HexToAddress(string(id)), &(c.client))
	if err != nil {
		return
	}

	value, err := instance.WeightsAddress(nil)
	if err != nil {
		return
	}

	address = com.StorageAddress(value)
	return
}

func (c *ethereumChain) SubmitAggregation(id com.ModelIdentifier, address com.StorageAddress) (err error) {

	instance, err := contract.NewContract(common.HexToAddress(string(id)), &(c.client))
	if err != nil {
		return
	}

	nonce, err := c.client.PendingNonceAt(context.Background(), c.publicAddress)
	if err != nil {
		return

	}

	gasPrice, err := c.client.SuggestGasPrice(context.Background())
	if err != nil {
		return
	}

	auth := bind.NewKeyedTransactor(&(c.privateKey))
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	tx, err := instance.SubmitLocalAggregation(auth, string(address))
	if err != nil {
		return
	}

	logger.Debugf("Wrote local update to chain as tx: %s", tx.Hash().Hex())
	return
}

func (c *ethereumChain) SubmitLocalUpdate(id com.ModelIdentifier, update com.Update) (err error) {

	instance, err := contract.NewContract(common.HexToAddress(string(id)), &(c.client))
	if err != nil {
		return
	}

	nonce, err := c.client.PendingNonceAt(context.Background(), c.publicAddress)
	if err != nil {
		return

	}

	gasPrice, err := c.client.SuggestGasPrice(context.Background())
	if err != nil {
		return
	}

	auth := bind.NewKeyedTransactor(&(c.privateKey))
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	tx, err := instance.SubmitLocalUpdate(auth, string(update.Address))
	if err != nil {
		return
	}

	logger.Debugf("Wrote local update to chain as tx: %s", tx.Hash().Hex())
	return
}

func (c *ethereumChain) LocalUpdates(id com.ModelIdentifier) (updates []com.Update, err error) {

	instance, err := contract.NewContract(common.HexToAddress(string(id)), &(c.client))
	if err != nil {
		return
	}

	count, err := instance.LocalUpdatesCount(nil)
	if err != nil {
		return
	}

	for i := big.NewInt(0); i.Cmp(count) == -1; i.Add(i, big.NewInt(1)) {

		update, err := instance.LocalUpdates(nil, i)
		if err != nil {
			return nil, err
		}
		updates = append(updates, com.Update{
			Trainer: com.TrainerIdentifier(update.Trainer),
			Address: com.StorageAddress(update.StorageAddress),
		})
	}

	return
}

func (c *ethereumChain) ModelEpoch(id com.ModelIdentifier) (epoch int, err error) {

	instance, err := contract.NewContract(common.HexToAddress(string(id)), &(c.client))
	if err != nil {
		return
	}

	value, err := instance.Epoch(nil)
	if err != nil {
		return
	}

	epoch = int(value.Int64())
	return
}

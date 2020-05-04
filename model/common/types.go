package common

import "github.com/ethereum/go-ethereum/common"

type StorageAddress string
type ModelIdentifier string
type TrainerIdentifier common.Address

type Update struct {
	Trainer TrainerIdentifier
	Address StorageAddress
}

type Hyperparameters struct {
	UpdatesTillAggregation int
	Epochs                 int
}

const Training uint8 = 0
const Aggregation uint8 = 1
const Finished uint8 = 2

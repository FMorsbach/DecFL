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
}

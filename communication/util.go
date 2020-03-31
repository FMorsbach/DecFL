package communication

type StorageAddress string
type ModelIdentifier string
type TrainerIdentifier string

type Update struct {
	Trainer TrainerIdentifier
	Address StorageAddress
}

package model

import (
	"flag"
	"fmt"

	ch "github.com/FMorsbach/DecFL/model/chain"
	"github.com/FMorsbach/DecFL/model/chain/ethereum"
	"github.com/FMorsbach/DecFL/model/common"
	st "github.com/FMorsbach/DecFL/model/storage"
)

func ParseCLIConfig() (chain ch.Chain, store st.Storage, modelID common.ModelIdentifier, err error) {

	var storageType string
	var chainConnection string
	var storageConnection string
	var privateKey string
	var modelAddress string
	var redisPassword string

	flag.StringVar(&storageType, "storageType", "", "Wheter to choose IPFS or redis as storage")
	flag.StringVar(&chainConnection, "c", "", "The connection identifier to the chain")
	flag.StringVar(&storageConnection, "s", "", "The connection identifier to the storage service")
	flag.StringVar(&privateKey, "k", "", "the private key for the chain")
	flag.StringVar(&redisPassword, "redisPW", "", "(Optional) The password for redis if used")
	flag.StringVar(&modelAddress, "m", "", "The address of the model")

	required := []string{"c", "s", "k"}
	flag.Parse()

	seen := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) { seen[f.Name] = true })
	for _, req := range required {
		if !seen[req] {
			err = fmt.Errorf("missing required -%s argument/flag\n", req)
			return
		}
	}

	chain, err = ethereum.NewEthereum(chainConnection, privateKey)
	if err != nil {
		return
	}

	// setup storage
	switch storageType {
	case "redis":
		redis := st.NewRedis(storageConnection, redisPassword)
		var ok bool
		if ok, err = redis.IsReachable(); !ok {
			err = fmt.Errorf("Cant reach redis: %s", err)
			return
		}
		store = redis

	case "ipfs":
		ipfs := st.NewIPFS(storageConnection)
		if !ipfs.IsReachable() {
			err = fmt.Errorf("Cant reach ipfs node")
			return
		}
		store = ipfs

	default:
		err = fmt.Errorf("Invalid storage type")
		return
	}

	// setup model
	if modelAddress != "" {
		modelID = common.ModelIdentifier(modelAddress)
	}

	return
}

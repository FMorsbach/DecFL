#!/bin/bash
echo "Deploying on remote setup"

CHAIN=http://DecFL-S:8545
STORAGE=localhost:5001
STORAGE_TYPE=ipfs
KEY=3b3a098805d048bab52b82b8767da2117af104cc97ec820acbe1b63e768ebba7

go run controller.go -c $CHAIN -s $STORAGE -k $KEY -storageType $STORAGE_TYPE

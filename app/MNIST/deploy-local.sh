#!/bin/bash
echo "Deploying on local setup"

CHAIN=http://localhost:8545
STORAGE=localhost:6379
STORAGE_TYPE=redis
KEY=3b3a098805d048bab52b82b8767da2117af104cc97ec820acbe1b63e768ebba7

go run controller.go -c $CHAIN -s $STORAGE -k $KEY -storageType $STORAGE_TYPE

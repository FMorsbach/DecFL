#!/bin/bash
CHAIN=http://localhost:8545
STORAGE=localhost:6379
STORAGE_TYPE=redis
KEY=$(cat secrets/local/key)

echo "Packing scripts"
back=$PWD
cd $APP
tar -czvf scripts.tar.gz scripts > /dev/null
mv scripts.tar.gz $back/scripts.tar.gz
cd $back

echo "Deploying to local network"
go run controller.go \
-chain $CHAIN \
-storage $STORAGE \
-key $KEY \
-storageType $STORAGE_TYPE \
-scripts scripts.tar.gz \
-config $APP/configuration.txt \
-weights $APP/weights.txt \
-trainer secrets/local/trainers

rm scripts.tar.gz

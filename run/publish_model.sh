#!/bin/bash

while getopts a:n: option
do
case "${option}"
in
a) APP=${OPTARG};;
n) NETWORK=${OPTARG};;
esac
done


if [ -z "$APP" ];
then
    echo "You need to specify an application to deploy with -a"
    exit 1
fi

if [ -z "$NETWORK" ];
then
    echo "You need to specify a valid network configuration with -n"
    exit 1
fi

CHAIN=$(cat ../res/secrets/$NETWORK/chain)
STORAGE=$(cat ../res/secrets/$NETWORK/storage)
STORAGE_TYPE=$(cat ../res/secrets/$NETWORK/storage_type)
KEY=$(cat ../res/secrets/$NETWORK/deploy_key)
TRAINERS=../res/secrets/$NETWORK/trainers


echo "Packing scripts"
back=$PWD
cd $APP
tar -czvf scripts.tar.gz scripts > /dev/null
mv scripts.tar.gz $back/scripts.tar.gz
cd $back

echo "Deploying to network"
go run ../app/deploy/deploy.go \
-chain $CHAIN \
-storage $STORAGE \
-storageType $STORAGE_TYPE \
-key $KEY \
-scripts scripts.tar.gz \
-config $APP/configuration.txt \
-weights $APP/weights.txt \
-trainer $TRAINERS

rm scripts.tar.gz

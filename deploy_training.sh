#!/bin/bash

while getopts a:s: option
do
case "${option}"
in
a) APP=${OPTARG};;
s) SCENARIO=${OPTARG};;
esac
done


if [ -z "$APP" ];
then
    echo "You need to specify an application to deploy with -a"
    exit 1
fi

if [ -z "$SCENARIO" ];
then
    echo "You need to specify a valid scenario with -s"
    exit 1
fi

CHAIN=$(cat scenarios/$SCENARIO/chain)
STORAGE=$(cat scenarios/$SCENARIO/storage)
STORAGE_TYPE=$(cat scenarios/$SCENARIO/storage_type)
KEY=$(cat scenarios/$SCENARIO/deploy_key)
TRAINERS=scenarios/$SCENARIO/trainers


echo "Packing scripts"
back=$PWD
cd $APP
tar -czvf scripts.tar.gz scripts > /dev/null
mv scripts.tar.gz $back/scripts.tar.gz
cd $back

echo "Deploying to $SCENARIO"
go run app/deploy/deploy.go \
-chain $CHAIN \
-storage $STORAGE \
-storageType $STORAGE_TYPE \
-key $KEY \
-scripts scripts.tar.gz \
-config $APP/configuration.txt \
-weights $APP/weights.txt \
-trainer $TRAINERS

rm scripts.tar.gz

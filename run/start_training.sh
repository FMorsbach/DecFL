#!/bin/bash

while getopts m:h:n: option
do
case "${option}"
in
m) MODEL=${OPTARG};;
h) HOST=${OPTARG};;
n) NETWORK=${OPTARG};;
esac
done

if [ -z "$MODEL" ];
then
    echo "you need to specify a model with -m"
    exit 1
fi

if [ -z "$NETWORK" ];
then
    echo "you need to specify a target network with -n"
    exit 1
fi

if [ -z "$HOST" ];
then
    echo "You need to specify the target host with -h"
    exit 1
fi

if [ "$HOST" == "local" ];
then

    if [ "$NETWORK" == "local" ]
    then 
        cp ../res/secrets/local/network_local.env local/network.env
        cp ../res/secrets/local/accounts_local.env local/.env
    else
        cp ../res/secrets/local/network_$NETWORK.env local/network.env
        cp ../res/secrets/local/accounts_$NETWORK.env local/.env
    fi
   
    cd local
    MODEL=$MODEL docker-compose up
    rm .env
    rm network.env
    cd ../
else
    CHAIN=$(cat ../res/secrets/$NETWORK/chain)
    ansible-playbook remote/run.yml -e CHAIN=$CHAIN -e MODEL=$MODEL -e HOST=$HOST
fi

    
#!/bin/bash

while getopts c:s: option
do
case "${option}"
in
c) CONTRACT=${OPTARG};;
s) SCENARIO=${OPTARG};;
esac
done

if [ -z "$CONTRACT" ];
then
    echo "you need to specify a contract id with -c"
    exit 1
fi

if [ -z "$SCENARIO" ];
then
    echo "You need to specify the scenario with -s"
    exit 1
fi

if [ "$SCENARIO" == "local" ];
then
    cd scenarios/local

    sudo MODEL=$CONTRACT docker-compose up
else
    CHAIN=$(cat scenarios/$SCENARIO/chain)
    ansible-playbook remote/run.yml -e CHAIN=$CHAIN -e MODEL=$CONTRACT -e HOST=$HOST
fi

    
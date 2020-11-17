#!/bin/bash

while getopts c:s:n:k: option
do
case "${option}"
in
c) CONTRACT=${OPTARG};;
s) SCENARIO=${OPTARG};;
n) NAME=${OPTARG};;
k) SIZE=${OPTARG};;
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

if [ -z "$NAME" ];
then
    echo "You need to specify the experiment name with -n"
    exit 1
fi

if [ -z "$SIZE" ];
then
    echo "You need to specify the number of clients with -k"
    exit 1
fi


if [ "$SCENARIO" == "local" ];
then
    echo "Not implemented yet"
    exit 1
else
    ansible-playbook \
        -i scenarios/$SCENARIO/hosts \
        -e MODEL=$CONTRACT \
        -e EXPERIMENT=$NAME \
        -e SIZE=$SIZE \
        playbooks/logs.yml 
fi

    

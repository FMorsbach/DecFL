#!/bin/bash

while getopts s: option
do
case "${option}"
in
s) SCENARIO=${OPTARG};;
esac
done

if [ -z "$SCENARIO" ];
then
    echo "You need to specify the scenario with -s"
    exit 1
fi

ansible-playbook remote/deploy.yml -e HOST=$SCENARIO
#!/bin/bash

while getopts h: option
do
case "${option}"
in
h) HOST=${OPTARG};;
esac
done

if [ -z "$HOST" ];
then
    echo "You need to specify the target host with -h"
    exit 1
fi

ansible-playbook remote/deploy.yml -e HOST=$HOST
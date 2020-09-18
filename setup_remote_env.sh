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


ansible-playbook -i scenarios/$SCENARIO/hosts playbooks/setup.yml
ansible-playbook -i scenarios/$SCENARIO/hosts playbooks/deploy.yml


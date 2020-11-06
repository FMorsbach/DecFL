#!/bin/bash


while getopts m: option
do
case "${option}"
in
m) MODE=${OPTARG};;
esac
done


cd app/worker

echo ""
echo ""
echo "##########################"
echo "## Building executables ##"
echo "##########################"
echo ""
echo ""

if [ "$MODE" == "native" ];
then
    echo "Building natively"
    echo ""
    go build
    cd ../deploy
    go build
    cd ../worker
else
    sudo docker run --rm -v "$PWD/../../":/usr/src -w /usr/src/app/worker golang:1.15 go build
    sudo docker run --rm -v "$PWD/../../":/usr/src -w /usr/src/app/deploy golang:1.15 go build
fi

echo ""
echo ""
echo "##########################"
echo "##### Building image #####"
echo "##########################"
echo ""
echo ""
sudo docker build -t decfl-worker:latest .

echo ""
echo ""

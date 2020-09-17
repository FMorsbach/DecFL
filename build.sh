#!/bin/bash

cd app/worker

echo ""
echo ""
echo "##########################"
echo "## Building executables ##"
echo "##########################"
echo ""
echo ""
sudo docker run --rm -v "$PWD/../../":/usr/src -w /usr/src/app/worker golang:1.15 go build
sudo docker run --rm -v "$PWD/../../":/usr/src -w /usr/src/app/deploy golang:1.15 go build

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

#!/bin/bash
echo "Building executable"
go build

echo "Building image"
docker build -t worker:latest . >/dev/null 



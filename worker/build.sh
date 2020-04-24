#!/bin/bash
echo "Building executable"
go build
echo ""

echo "Building image"
docker build -t worker:latest .
echo ""



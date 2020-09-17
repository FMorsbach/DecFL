#!/bin/bash

sudo docker network create decfl-net
sudo docker run --rm -d --name chain -p 127.0.0.1:8545:8545 --network decfl-net trufflesuite/ganache-cli -m "$(cat scenarios/local/mnemonic)"
sudo docker run --rm -d --name redis -p 127.0.0.1:6379:6379 --network decfl-net redis:latest

#!/bin/bash

sudo docker network inspect decfl_net >/dev/null 2>&1 ||  sudo docker network create decfl_net
sudo docker run --rm -d --name chain -p 127.0.0.1:8545:8545 --network decfl_net trufflesuite/ganache-cli -m "$(cat scenarios/local/mnemonic)"
sudo docker run --rm -d --name redis -p 127.0.0.1:6379:6379 --network decfl_net redis:latest

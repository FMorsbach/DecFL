# DecFL

## Requirements
- docker
- docker-compose
- golang (optional, required if you want to build the application for your host machine rather than for docker)
- git (optional, for cloning this repository)

## Compile the application
Easiest to be done in a throw away golang docker container (https://hub.docker.com/_/golang/).
The build script is already setup in this way and will produce a decfl-worker:latest image by default.
Build the image via "./build.sh".

## Scenario 1: Development setup, all local
This scenario is best to use for development as no remote machines or advanced Ethereum account management is required.
In this scenario all components are run on the host machine as docker containers with a simulated Ethereum chain. 
In the default configuration there will be 3 worker nodes, 1 redis container for storage (https://hub.docker.com/_/redis/) and 1 ganache-cli container for simulating the ethereum network (https://github.com/trufflesuite/ganache-cli).

## Prepare the local environment
The "setup_local_env.sh" script will setup a local environment to be used in this scenario. It will create a ganache-cli container with a specific mnemonic to match the default accounts stored in the scenarios/local directory. It will also generate a dedicated docker network for the inter-container communcation called decfl-net. 

Hint: You might adjust this as you please but remember to make sure that the accounts and private keys match the configuration of your ganache-cli setup. 

Setup the environment with "./setup_local_env.sh". 

## Deploy a model to be trained
By default there is one model predefined and can be trained immediately. 
To deploy the model (create a smart contract and store the model configuration in the storage system) use the "deploy_training.sh" script. 
The script requires two arguments,
- "-a" a path to the model you wish to deploy. Use "res/MNIST" for the included model.
- "-s" the scenario which you would like to use. Use "local" for local deployment.

Run "./deploy_trainging.sh -a res/MNIST -s local". The script will return the address of the deployed smart contract. You will need this address in the next step.

## Training the model 
For training a model locally with 3 clients there is the "start_training.sh" prebuild script. Used in this scenario it will spin up 3 docker decfl-worker containers and start training. 
The script requires two arguments,
- "-c" the contract id of the contract you want the clients to train. Use the address from the previous step here.
- "-s" the scenario which you would like to use. Use "local" for local deployment.

Run "./start_training.sh -c <CONTRACT_ID> -s local". 

## Scenario 2: Multiple machines, a shared ETH node (i.e. infura.io)


### HOSTS
- can be local to use multiple docker container running on the host
- can be a set of remote servers managed by ansible

### MODEL
the contract ID

### NETWORK
what ethereum node to use
- can be a local one such as the ganache-cli docker container for testing purpose
- can be a remote testing one such as the infura endpoint for the rinkeby test network
- can be a local node for the rinkeby test network
- any other ethereum node

### Secrets
For easier configuration and automation connection details and other informations can be placed into files in the /res/secrets directory. The scripts in /run expect secret files in a certain structure which is explained in the following. Relevant secrets are:
- $NETWORK/chain holds a single line with the URL of the ethereum node to connect to (i.e. for infura.io rinkeby node: "https://rinkeby.infura.io/v3/<API_KEY>")
- $NETWORK/storage holds a single line with the URL of the storage system to connect to (i.e. for a local IPFS instance : "localhost:5001")
- $NETWORK/storage_type a single line with the type of the storage system, currently supported are "redis" or "ipfs" as valid values.
- $NETWORK/deploy_key contains the private key of the ethereum account which should be used to create the training contract
- $NETWORK/trainers contains the public keys of all ethereum accounts (one per line) that should be registered as trainers in the smart contract. Only these accounts are allowed to submit training or aggregation results. Hint: No trailing empty line!
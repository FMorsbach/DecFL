# DecFL

## Requirements
- docker
- docker-compose
- ansible (optional, required if you want to control remote hosts for training using the provided playbooks and scripts)
- golang (optional, required if you want to build the application for your host machine rather than for docker)
- git (optional, for cloning this repository)

## Compile the application
Easiest to be done in a throw away [golang docker container](https://hub.docker.com/_/golang/).
The build script is already setup in this way and will produce a `decfl-worker:latest` image and the deployment application `deploy` by default.
Build the image via `./build.sh`.

## Scenario 1: Development setup, all local
This scenario is best to use for development as no remote machines or advanced Ethereum account management is required.
In this scenario all components are run on the host machine as docker containers with a simulated Ethereum chain. 
In the default configuration there will be 3 worker nodes, 1 [redis container](https://hub.docker.com/_/redis/) for storage and 1 [ganache-cli container](https://github.com/trufflesuite/ganache-cli) for simulating the Ethereum network.

### Prepare the local environment
The `setup_local_env.sh` script will setup a local environment to be used in this scenario. It will create a ganache-cli container with a specific mnemonic to match the default accounts stored in the scenarios/local directory. It will also generate a dedicated docker network for the inter-container communcation called decfl-net. 

Hint: You might adjust this as you please but remember to make sure that the accounts and private keys match the configuration of your ganache-cli setup. 

Setup the environment with `./setup_local_env.sh`. 

### Deploy a model to be trained
By default there is one model predefined and can be trained immediately. 
To deploy the model (create a smart contract and store the model configuration in the storage system) use the `deploy_training.sh` script. 
The script requires two arguments,
- `-a` a path to the model you wish to deploy. Use `res/MNIST"` for the included model.
- `-s` the scenario which you would like to use. Use `local` for local deployment.

Run `./deploy_trainging.sh -a res/MNIST -s local`. The script will return the address of the deployed smart contract. You will need this address in the next step.

### Training the model 
For training a model locally with 3 clients there is the `start_training.sh` script. Used in this scenario it will spin up 3 docker `decfl-worker` containers and start training. 
The script requires two arguments,
- `-c` the contract id of the contract you want the clients to train. Use the address from the previous step here.
- `-s` the scenario which you would like to use. Use `local` for local deployment.

Run `./start_training.sh -c <CONTRACT_ID> -s local`. 

## Scenario 2: Multiple machines, a shared ETH node (i.e. infura.io)

Create a directory in the `scenarios` directory which will hold your confiugration. The configuration will be done through simple files inside this directory. 

### Hosts
Create a `hosts` file with all the trainer hosts you want to control. Ansible will use this file as [inventory](https://docs.ansible.com/ansible/latest/user_guide/intro_inventory.html). 

Make sure to validate your setup by using the ping module of ansible:
`ansible all -i scenarios/<SCENARIO_NAME>/hosts -m ping`.

### Prepare the scenario configuration
- `chain` holds a single line with the URL of the ethereum node to connect to (i.e. for the infura.io rinkeby node: "https://rinkeby.infura.io/v3/<API_KEY>". You will need to generate your own infura API key if you plan to use its node.)
- `deploy_key` contains the private key of the ethereum account which should be used to create the training contract
- `deploy_storage` holds a single line with the URL of the storage system to connect to when deploying the model (i.e. for the infura.io IPFS gateway: "https://ipfs.infura.io:5001")
- `trainers` contains the public keys of all ethereum accounts (one per line) that should be registered as trainers in the smart contract. Only these accounts are allowed to submit training or aggregation results. Hint: No trailing empty line!

### Prepare the nodes
On the nodes you need to install docker. 
After that you can run the playbooks/setup.yml and playbooks/deploy.yml playbooks to prepare the nodes.
Additionally you need to manually provision each node with an identity file. This file has to exisist as ~/app/identity and provide 3 key-value pairs. 
- `KEY=<ehtereum-private-key>` the private key to be used
- `DECFL_PARTITIONS=<value>`, the number of participating trainer nodes
- `DECFL_ID=<value>`, the id of this node, must be in the range `[0,DECFL_PARTITIONS-1]`
There is currently no better way implemented, see Issue #8 for further information. 

### Deploy a model to be trained
The process of deploying a model is the same as on a local setup. Use the `deploy_training.sh` script and pass your scenario name as argument for `-s`. This might take a while in a non local setup. 

### Start the training
The process of starting the training is also the same as on a local setup. Use the `start_training.sh` script and pass your scenario name as argument for `-s`.

You can monitor the remote training with ansible as well:
`ansible all -i scenario/<SCENARIO_NAME>/hosts --become -a "docker logs worker_<FIRST_6_SYMBOLS_OF_CONTRACT_ID>"`


## Adding new models
*to be done* 
# DecFL

## Requirements
- docker
- docker-compose
- ansible (optional, required if you want to control remote hosts for training using the provided playbooks and scripts)
- golang (optional, required if you want to build the application for your host machine rather than for the docker image)
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
The `setup_local_env.sh` script will setup a local environment to be used in this scenario. It will create a ganache-cli container with a specific mnemonic to match the default accounts stored in the scenarios/local directory. It will also generate a dedicated docker network for the inter-container communication called decfl-net. 

Hint: You might adjust this as you please but remember to make sure that the accounts and private keys match the configuration of your ganache-cli setup. 

Setup the environment with `./setup_local_env.sh`. 

### Deploy a model to be trained
By default there is one model predefined and can be trained immediately. 
To deploy the model (create a smart contract and store the model configuration in the storage system) use the `deploy_training.sh` script. 
The script will return the address of the deployed smart contract. You will need this address in the next step.

The script requires two arguments,
- `-a` the path to the model you wish to deploy. Use `res/MNIST` for the included model.
- `-s` the scenario which you would like to use. Use `local` for local deployment.

Run `./deploy_training.sh -a res/MNIST -s local`. 

### Training the model 
For training a model locally with 3 clients there is the `start_training.sh` script. Used in this scenario it will spin up 3 docker `decfl-worker` containers and start training. 
The script requires two arguments,
- `-c` the contract id of the contract you want the clients to train. Use the address from the previous step here.
- `-s` the scenario which you would like to use. Use `local` for local deployment.

Run `./start_training.sh -c <CONTRACT_ID> -s local`. 
You will see the output of all three containers printed to the screen.

## Scenario 2: Multiple machines, a shared ETH node (i.e. infura.io)

In this scenario the training will happen on multiple dedicated machines to simulate real world communication implications. Each node will host it's own IPFS node and participate in the official IPFS network. 
As synching a full Ethereum node can take a lot of time and the impact on the experiment are negligible, a shared public Ethereum node will be used (i.e. from the infura.io service).

For configuring the scenario create a directory in the corresponding scenarios` directory which will hold your configuration. The configuration will be done through multiple files inside this directory. 

### Prepare the scenario configuration
Create the following files for configuration:

- The `hosts` file contains the trainer hosts you want to control. 
Ansible will use this file as [inventory](https://docs.ansible.com/ansible/latest/user_guide/intro_inventory.html). 
Make sure to validate your setup by using the ping module of ansible:
`ansible all -i scenarios/<SCENARIO_NAME>/hosts -m ping`.

- The `chain` file contains a single line with the URL of the Ethereum node to connect to (i.e. for the infura.io rinkeby node: "https://rinkeby.infura.io/v3/<API_KEY>". You will need to generate your own infura API key if you plan to use infura).

- The `deploy_key` file contains the private key of the Ethereum account which should be used to create the training contract.

- The `deploy_storage` file contains a single line with the URL of the storage system to connect to when deploying the model (i.e. for the infura.io IPFS gateway: "https://ipfs.infura.io:5001").

- The `trainers` file contains the public keys of all Ethereum accounts (one per line) that should be registered as trainers in the smart contract. Only these accounts are allowed to submit training or aggregation results. Hint: No trailing empty line!

HINT: Make sure that all Ethereum accounts have sufficient funds to pay for the transactions. The rinkeby test network is sufficient for testing and you can request funds for free.

### Prepare the nodes
Each node requires a working docker installation, check your OS/distribution for install instructions.
Once you have a working docker installation, you can run the `playbooks/setup.yml` and `playbooks/deploy.yml` playbooks to prepare the nodes. 
You can also use the `setup_remote_env.sh` script for this:

`./setup_remote_env.sh -s <SCENARIO_NAME>`

Additionally you need to manually provision each node with an identity file. This file has to exists as ~/app/identity and provide 3 key-value pairs. 
- `KEY=<ethereum-private-key>` the private key to be used
- `DECFL_PARTITIONS=<value>`, the number of participating trainer nodes
- `DECFL_ID=<value>`, the id of this node, must be in the range `[0,DECFL_PARTITIONS-1]`

Unfortunately, there is no better way to automate this process implemented yet, see Issue #8 for further information. 

### Deploy a model to be trained
The process of deploying a model is the same as on a local setup. Use the `deploy_training.sh` script and pass your scenario name as argument for `-s`. This might take a while in a non local setup. 

### Start the training
The process of starting the training is also the same as on a local setup. Use the `start_training.sh` script and pass your scenario name as argument for `-s`.

You can monitor the remote training with ansible as well:

`ansible all -i scenario/<SCENARIO_NAME>/hosts --become -a "docker logs worker_<FIRST_6_SYMBOLS_OF_CONTRACT_ID>"`

## Adding new models
The core `DecFL` application is model agnostic, so one can implement custom models.
With the current setup, only python-based Tensorflow models are supported (To use a different machine learning framework, see section [Adding new vendor drivers](https://github.com/FMorsbach/DecFL/#adding-new-vendor-drivers). 

To guarantee interoperability with the implemented tensorflow driver a model implementation must adhere to the following specification: 

- All relevant files should be placed in a single directory (i.e. in `res/<MODEL_NAME>`).
- The model's configuration (i.e. the serialized definition of the neural net) must be placed in `<MODEL_ROOT>/configuration.txt`.
- The model's initial weights must be placed in `<MODEL_ROOT>/weights.txt`.
- All source code required for training, aggregation and evaluation must be placed in `<MODEL_ROOT>/scripts/`. Note that the contents of this directory will be uploaded unencrypted to the storage system and distributed to each worker node.
- The scripts directory needs to at least contain the following three executable python scripts:

    - `train.py`: This script is responsible for executing one global training round. It must accept the following three command line arguments and behave as follows:
        - `-c` The path to the configuration file. This file is to be read by the script and parsed as configuration for the model.
        - `-w` The path to the weights file. This file is to be read by the script and parsed as starting weights for the current training round.
        - `-o` The path to the output file. The script must write it's training result (i.e. the model's weights after the training) to this file.
    
    - `aggregate.py`: This script is responsible for aggregating multiple weights files into a single weights file. It must accept the following two command line arguments and behave as follows:
        - `-w`: The path to a directory that contains all input weights. Each input weights file will have the suffix `_trainingWeights.in`. 
        - `-o`: The path to the output file. The script must write it's aggregation result to this file.

    - `evaluate.py`: This script is responsible for evaluating the performance of a given model. It must accept the following three command line arguments and behave as follows:
        - `-c` The path to the configuration file. This file is to be read by the script and parsed as configuration for the model.
        - `-w` The path to the weights file. This file is to be read by the script and parsed as weights for the model.
        - `-o` The path to the output file. The script must write it's evaluation result to this file. The output must match the following format: `[<LOSS>, <ACCURACY>]`.

Any unrecoverable error should be printed to `sys.stderr`. 
You might refer to the provided MNIST implementation for easier bootstrapping. 

## Adding new vendor drivers
`DecFL`is built to be vendor agnostic. This means that one can use different  decentralized storage, decentralized computing or machine learning frameworks. 
For this one has to implement such called vendor drivers and adhere to the interface specification of said driver. One example of this is that there are two storage drivers implemented at the moment, one for local development with redis and one for IPFS in a distributed setup. Unfortunately there is no documentation on how to do this yet, but one should be able to bootstrap new drivers by looking at the existing ones fairly easy.
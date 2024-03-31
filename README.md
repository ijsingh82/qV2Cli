# qV2Cli

`qV2Cli` is a CLI application used to interact with the V2 permission contracts for the GoQuorum permission network. It utilizes the GoQuorum fork of `geth`, the communication is unsecure so please use in development only.

## Features

- Interact with V2 permission contracts on GoQuorum permissioned networks.
- Boot up a V2 permission network.
- Perform transactions using an address and secret key.
- Perform transactions using Ethereum keystore with or without a password.
- Vendor folder contains all dependencies.

## Prerequisites

- `go` installed on your machine.
- Replace the `permission-config.json` file present in the root folder for communicating with the network with the generated file from your network (see example permission-config.json).
- Add a new line for 'rpc' with the rpc for the network

## Installation

To install `qV2Cli`, you can clone the repository and build the binary yourself:

```bash
git clone https://github.com/ijsingh82/qV2Cli.git
cd qV2Cli
go build -o qV2Cli
```

OR

```bash
git clone https://github.com/ijsingh82/qV2Cli.git
cd qV2Cli
make build
```

## Usage

1. Ensure that permission-config.json is present in the root folder and properly configured to communicate with your V2 permission network. Add a new line for 'rpc' with the rpc for the network see example file.

2. Get info from the V2 contracts with commands and arugments for non trasaction calls all require arguments.

```bash
qV2Cli acct balKey [keystorePath] [password]
qV2Cli acct count
```

The above commands allow you to get information from the chain -h for all commands and subcommands.

3. Transactional functions use flags with a secrect key and address or a keystore. To get all needed flags use -h.

   Keystore:

```bash
qV2Cli addAdminKey -newAdmin="TEST" -keyfile="Path/to/kestore/file" -passowrd=""
or
qV2Cli addAdminKey -a="TEST" -k="Path/to/kestore/file" -p=""
or
qV2Cli addAdmin -a="TEST" -k="secretKey" -g="guradianAddress"
or
qV2Cli addAdmin -newAdmin="TEST" -guardKey="secretKey" -guardAdd="guradianAddress"
```

The commands above use an Ethereum keystore for performing transactions. Optionally, you can provide a password if the keystore is encrypted.

# License

This project is licensed under the MIT License.

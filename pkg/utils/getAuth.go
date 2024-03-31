package utils

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetNextAuth(add, key string, client *ethclient.Client) (*bind.TransactOpts, error) {

	var privateKey *ecdsa.PrivateKey
	// Get gasprice
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	if key[:2] != "0x" {
		// Get Private Key
		privateKey, err = crypto.HexToECDSA(key)
		if err != nil {
			return nil, err
		}
	} else {
		// Get Private Key
		privateKey, err = crypto.HexToECDSA(key[2:])
		if err != nil {
			return nil, err
		}
	}

	// Get nonce
	nonce, err := client.PendingNonceAt(context.Background(), common.HexToAddress(add))
	if err != nil {
		fmt.Println("Nonce Error")
		return nil, err
	}

	// Get ChainId
	id, err := client.ChainID(context.Background())
	if err != nil {
		fmt.Println("Chain Id Error")
		return nil, err
	}

	// /crate an authorization to deploy the contract
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, id)
	if err != nil {
		fmt.Println("new transactor Error")
		return nil, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)       // in wei
	auth.GasLimit = uint64(30000000) // in units
	auth.GasPrice = gasPrice

	return auth, nil
}

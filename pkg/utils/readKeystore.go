package utils

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// returns address then private key
func ReadKeystore(pass, path string) (string, string, error) {

	b, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("readfile error")
		return "", "", err
	}

	var key *keystore.Key

	if pass == "" {
		key, err = keystore.DecryptKey(b, "")
		if err != nil {
			fmt.Println("Private Key decrypt error")
			return "", "", err
		}
	} else {
		key, err = keystore.DecryptKey(b, pass)
		if err != nil {
			fmt.Println("Private Key decrypt error")
			return "", "", err
		}
	}

	data := crypto.FromECDSA(key.PrivateKey)

	privateKey := hexutil.Encode(data)

	address := crypto.PubkeyToAddress(key.PrivateKey.PublicKey).Hex()
	return address, privateKey, nil

}

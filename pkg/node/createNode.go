package node

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ijsingh82/qV2Cli/pkg/model"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// returns path of data folder
func Create(instName, password string) (model.CreatedNode, error) {

	var node model.CreatedNode
	path := filepath.Dir("./" + instName + "Node/data/")
	acctPath := filepath.Dir("./" + instName + "Node/data/keystore/")

	ks := keystore.NewKeyStore(path, keystore.StandardScryptN, keystore.StandardScryptP)
	ks0 := keystore.NewKeyStore(acctPath, keystore.StandardScryptN, keystore.StandardScryptP)
	// am := accounts.NewManager(&accounts.Config{InsecureUnlockAllowed: false}, ks)
	acc, err := ks.NewAccount(password)
	if err != nil {
		return node, err

	}

	valAccnt, err := ks0.NewAccount(password)
	if err != nil {
		return node, err

	}

	b, err := os.ReadFile(acc.URL.Path)
	if err != nil {
		return node, err
	}

	key, err := keystore.DecryptKey(b, password)
	if err != nil {
		fmt.Println("Private Key decrypt error")
		return node, err
	}

	data := crypto.FromECDSA(key.PrivateKey)
	//this is the private key for node
	privateKey := hexutil.Encode(data)
	err = func() error {
		var data []byte = []byte(privateKey[2:])
		var name string = fmt.Sprintf("%v/nodekey", path)
		f, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.FileMode(0644))
		if err != nil {
			return err
		}
		_, err = f.Write(data)
		if err1 := f.Close(); err1 != nil && err == nil {
			err = err1
		}
		return err
	}()
	if err != nil {
		return node, err
	}

	pKey := crypto.FromECDSAPub(&key.PrivateKey.PublicKey)
	// This is the public key
	publicKey := hexutil.Encode(pKey)
	err = func() error {
		var data []byte = []byte(publicKey[4:])
		var name string = fmt.Sprintf("%v/nodekey.pub", path)
		f, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.FileMode(0644))
		if err != nil {
			return err
		}
		_, err = f.Write(data)
		if err1 := f.Close(); err1 != nil && err == nil {
			err = err1
		}
		return err
	}()
	if err != nil {
		return node, err
	}

	address := crypto.PubkeyToAddress(key.PrivateKey.PublicKey).Hex()
	err = func() error {
		var data []byte = []byte(address[2:])
		var name string = fmt.Sprintf("%v/address", path)
		f, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.FileMode(0644))
		if err != nil {
			return err
		}
		_, err = f.Write(data)
		if err1 := f.Close(); err1 != nil && err == nil {
			err = err1
		}
		return err
	}()
	if err != nil {
		return node, err
	}

	node.ValidatorAccount = valAccnt.Address.String()
	node.NodePath = path

	os.Remove(acc.URL.Path)

	return node, nil

}

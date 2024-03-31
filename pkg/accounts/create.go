package accounts

import (
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func CreateAccount(password, name string) (accounts.Account, error) {

	acctPath := filepath.Dir("./keystore/")

	ks0 := keystore.NewKeyStore(acctPath, keystore.StandardScryptN, keystore.StandardScryptP)

	acc0, err := ks0.NewAccount(password)
	if err != nil {
		return accounts.Account{}, err

	}

	err = os.Rename(acc0.URL.Path, "./keystore/"+name)
	if err != nil {
		return accounts.Account{}, err

	}

	return acc0, nil

}

package accounts

import (
	"context"
	"math/big"

	"github.com/ijsingh82/qV2Cli/cmd"
	"github.com/ijsingh82/qV2Cli/pkg/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetBalKeyfile(keyfile, password string) (*big.Int, error) {

	//get client
	client, err := ethclient.Dial(cmd.Conf.Rpc)
	if err != nil {
		return &big.Int{}, err
	}
	defer client.Close()

	add, _, err := utils.ReadKeystore(password, keyfile)
	if err != nil {
		return &big.Int{}, err
	}

	bal, err := client.BalanceAt(context.Background(), common.HexToAddress(add), nil)
	if err != nil {
		return &big.Int{}, err
	}

	return bal, nil
}

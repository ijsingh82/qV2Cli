package accounts

import (
	"context"
	"math/big"

	"github.com/ijsingh82/qV2Cli/cmd"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetBal(add string) (*big.Int, error) {

	//get client
	client, err := ethclient.Dial(cmd.Conf.Rpc)
	if err != nil {
		return &big.Int{}, err
	}
	defer client.Close()

	bal, err := client.BalanceAt(context.Background(), common.HexToAddress(add), nil)
	if err != nil {
		return &big.Int{}, err
	}

	return bal, nil
}

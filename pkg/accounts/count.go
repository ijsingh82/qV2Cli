package accounts

import (
	"context"
	"math/big"

	b "github.com/ijsingh82/qV2Cli/bind"
	"github.com/ijsingh82/qV2Cli/cmd"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetAccountCount() (*big.Int, error) {

	client, err := ethclient.DialContext(context.Background(), cmd.Conf.Rpc)
	if err != nil {
		return big.NewInt(0), err
	}
	defer client.Close()

	acctMngr, err := b.NewAcctManager(common.HexToAddress(cmd.Conf.AccountMgrAddress), client)
	if err != nil {
		return big.NewInt(0), err
	}

	d, err := acctMngr.GetNumberOfAccounts(nil)
	if err != nil {
		return big.NewInt(0), err
	}

	return d, nil
}

package accounts

import (
	"context"

	b "github.com/ijsingh82/qV2Cli/bind"
	"github.com/ijsingh82/qV2Cli/cmd"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func CheckOrgAdmin(add, org, ult string) (bool, error) {

	client, err := ethclient.DialContext(context.Background(), cmd.Conf.Rpc)
	if err != nil {
		return false, err
	}
	defer client.Close()

	acctMngr, err := b.NewAcctManager(common.HexToAddress(cmd.Conf.AccountMgrAddress), client)
	if err != nil {
		return false, err
	}

	d, err := acctMngr.CheckOrgAdmin(nil, common.HexToAddress(add), org, ult)
	if err != nil {
		return false, err
	}

	return d, nil
}

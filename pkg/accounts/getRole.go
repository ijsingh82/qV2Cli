package accounts

import (
	"context"

	b "github.com/ijsingh82/qV2Cli/bind"
	"github.com/ijsingh82/qV2Cli/cmd"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetRoleForAccount(acct string) (string, error) {

	client, err := ethclient.DialContext(context.Background(), cmd.Conf.Rpc)
	if err != nil {
		return "", err
	}
	defer client.Close()

	acctMngr, err := b.NewAcctManager(common.HexToAddress(cmd.Conf.AccountMgrAddress), client)
	if err != nil {
		return "", err
	}

	role, err := acctMngr.GetAccountRole(nil, common.HexToAddress(acct))
	if err != nil {
		return "", err
	}

	return role, nil
}

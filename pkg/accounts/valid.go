package accounts

import (
	"context"

	b "github.com/ijsingh82/qV2Cli/bind"
	"github.com/ijsingh82/qV2Cli/cmd"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// returns bool true if the account does not exists or exists and belongs passed org
func ValidatedAccount(acct, org string) (bool, error) {

	client, err := ethclient.DialContext(context.Background(), cmd.Conf.Rpc)
	if err != nil {
		return false, err
	}
	defer client.Close()

	acctMngr, err := b.NewAcctManager(common.HexToAddress(cmd.Conf.AccountMgrAddress), client)
	if err != nil {
		return false, err
	}

	v, err := acctMngr.ValidateAccount(nil, common.HexToAddress(acct), org)
	if err != nil {
		return false, err
	}

	return v, nil
}

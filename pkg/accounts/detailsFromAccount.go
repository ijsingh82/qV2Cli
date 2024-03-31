package accounts

import (
	"context"

	b "github.com/ijsingh82/qV2Cli/bind"
	"github.com/ijsingh82/qV2Cli/cmd"
	"github.com/ijsingh82/qV2Cli/pkg/model"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetAccountDetails(acct string) (model.AccountDetails, error) {
	var d model.AccountDetails

	client, err := ethclient.DialContext(context.Background(), cmd.Conf.Rpc)
	if err != nil {
		return d, err
	}
	defer client.Close()

	acctMngr, err := b.NewAcctManager(common.HexToAddress(cmd.Conf.AccountMgrAddress), client)
	if err != nil {
		return d, err
	}

	d.Id, d.OrgId, d.Role, d.Status, d.IsAdmin, d.Err = acctMngr.GetAccountDetails(nil, common.HexToAddress(acct))
	if d.Err != nil {
		return d, err
	}

	return d, nil
}

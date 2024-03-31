package accounts

import (
	"context"

	b "github.com/ijsingh82/qV2Cli/bind"
	"github.com/ijsingh82/qV2Cli/cmd"
	"github.com/ijsingh82/qV2Cli/pkg/model"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetAccountOrgDetails(acct string) (model.AccountInfo, error) {
	var d model.AccountInfo

	client, err := ethclient.DialContext(context.Background(), cmd.Conf.Rpc)
	if err != nil {
		return d, err
	}
	defer client.Close()

	acctMngr, err := b.NewAcctManager(common.HexToAddress(cmd.Conf.AccountMgrAddress), client)
	if err != nil {
		return d, err
	}

	d.OrgId, d.Role, d.Err = acctMngr.GetAccountOrgRole(nil, common.HexToAddress(acct))
	if d.Err != nil {
		return d, err
	}

	return d, nil
}

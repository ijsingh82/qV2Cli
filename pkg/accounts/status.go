package accounts

import (
	"context"

	b "github.com/ijsingh82/qV2Cli/bind"
	"github.com/ijsingh82/qV2Cli/cmd"
	"github.com/ijsingh82/qV2Cli/pkg/model"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Account Status types
// 0 - Not in list
// 1 - Account pending approval
// 2 - Active
// 3 - Inactive
// 4 - Suspended
// 5 - Blacklisted
// 6 - Revoked
// 7 - Recovery Initiated for blacklisted accounts and pending approval
//
//	from network admins
func GetStatusOfAccount(acct string) (model.Account, error) {

	client, err := ethclient.DialContext(context.Background(), cmd.Conf.Rpc)
	if err != nil {
		return model.Account{}, err
	}
	defer client.Close()

	acctMngr, err := b.NewAcctManager(common.HexToAddress(cmd.Conf.AccountMgrAddress), client)
	if err != nil {
		return model.Account{}, err
	}

	status, err := acctMngr.GetAccountStatus(nil, common.HexToAddress(acct))
	if err != nil {
		return model.Account{}, err
	}

	modStatus := model.Account{}

	s, err := modStatus.AcctStatusString(status)
	if err != nil {
		return model.Account{}, err
	}

	return s, nil
}

package accounts

import (
	"context"
	"fmt"
	"math/big"

	b "github.com/ijsingh82/qV2Cli/bind"
	"github.com/ijsingh82/qV2Cli/cmd"
	"github.com/ijsingh82/qV2Cli/pkg/model"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetAccountDetailsFromIndex(index int64) (model.AccountDetails, error) {
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

	cnt, err := GetAccountCount()
	if err != nil {
		return d, err
	}

	if index > cnt.Int64()-1 {
		return d, fmt.Errorf("there are not that many accounts the current count is %v.  use an index less than count", cnt)
	}

	d.Id, d.OrgId, d.Role, d.Status, d.IsAdmin, d.Err = acctMngr.GetAccountDetailsFromIndex(nil, big.NewInt(index))
	if d.Err != nil {
		return d, err
	}

	return d, nil
}

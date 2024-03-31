package impl

import (
	"context"

	b "github.com/ijsingh82/qV2Cli/bind"
	"github.com/ijsingh82/qV2Cli/cmd"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ValidateAccount(acct, org string) (bool, error) {
	var v bool
	client, err := ethclient.DialContext(context.Background(), cmd.Conf.Rpc)
	if err != nil {
		return v, err
	}
	defer client.Close()

	permImpl, err := b.NewPermImpl(common.HexToAddress(cmd.Conf.ImplAddress), client)
	if err != nil {
		return v, err
	}

	// returns bool true if the account does not exists or exists and belongs passed org
	v, err = permImpl.ValidateAccount(nil, common.HexToAddress(acct), org)
	if err != nil {
		return v, err
	}

	return v, nil
}

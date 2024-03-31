package role

import (
	"math/big"

	b "github.com/ijsingh82/qV2Cli/bind"
	"github.com/ijsingh82/qV2Cli/cmd"

	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Count() (*big.Int, error) {

	var cnt *big.Int
	client, err := ethclient.DialContext(context.Background(), cmd.Conf.Rpc)
	if err != nil {
		return cnt, err
	}
	defer client.Close()

	roleMngr, err := b.NewRoleManager(common.HexToAddress(cmd.Conf.RoleMgrAddress), client)
	if err != nil {
		return cnt, err
	}

	cnt, err = roleMngr.GetNumberOfRoles(nil)
	if err != nil {
		return cnt, err
	}
	return cnt, nil
}

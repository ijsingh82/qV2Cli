package role

import (
	"context"

	b "github.com/ijsingh82/qV2Cli/bind"
	"github.com/ijsingh82/qV2Cli/cmd"
	"github.com/ijsingh82/qV2Cli/pkg/model"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Access(roleId, orgId, ultPrnt string) (string, error) {

	var roleAcc string
	client, err := ethclient.DialContext(context.Background(), cmd.Conf.Rpc)
	if err != nil {
		return roleAcc, err
	}
	defer client.Close()

	roleMngr, err := b.NewRoleManager(common.HexToAddress(cmd.Conf.RoleMgrAddress), client)
	if err != nil {
		return roleAcc, err
	}

	d, err := roleMngr.RoleAccess(nil, roleId, orgId, ultPrnt)
	if err != nil {
		return roleAcc, err
	}

	roleAcc, err = model.RoleAccesString(d)
	if err != nil {
		return roleAcc, err
	}
	return roleAcc, nil
}

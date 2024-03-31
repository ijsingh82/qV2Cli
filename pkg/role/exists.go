package role

import (
	"context"

	b "github.com/ijsingh82/qV2Cli/bind"
	"github.com/ijsingh82/qV2Cli/cmd"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Exists(roleId, orgId, ultPrnt string) (bool, error) {

	var exists bool
	client, err := ethclient.DialContext(context.Background(), cmd.Conf.Rpc)
	if err != nil {
		return exists, err
	}
	defer client.Close()

	roleMngr, err := b.NewRoleManager(common.HexToAddress(cmd.Conf.RoleMgrAddress), client)
	if err != nil {
		return exists, err
	}

	exists, err = roleMngr.RoleExists(nil, roleId, orgId, ultPrnt)
	if err != nil {
		return exists, err
	}

	return exists, nil
}

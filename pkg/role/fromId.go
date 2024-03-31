package role

import (
	b "github.com/ijsingh82/qV2Cli/bind"
	"github.com/ijsingh82/qV2Cli/cmd"
	"github.com/ijsingh82/qV2Cli/pkg/model"

	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func FromId(roleId, orgId string) (model.RoleDetail, string, error) {

	var d model.RoleDetail
	var stat string
	client, err := ethclient.DialContext(context.Background(), cmd.Conf.Rpc)
	if err != nil {
		fmt.Println("Client Dial Error")
		return d, stat, err
	}
	defer client.Close()

	roleMngr, err := b.NewRoleManager(common.HexToAddress(cmd.Conf.RoleMgrAddress), client)
	if err != nil {
		return d, stat, err
	}

	d, err = roleMngr.GetRoleDetails(nil, roleId, orgId)
	if err != nil {
		return d, stat, err
	}

	stat, err = model.RoleAccesString(d.AccessType)
	if err != nil {
		return d, stat, err
	}

	return d, stat, nil
}

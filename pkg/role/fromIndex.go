package role

import (
	b "github.com/ijsingh82/qV2Cli/bind"
	"github.com/ijsingh82/qV2Cli/cmd"
	"github.com/ijsingh82/qV2Cli/pkg/model"

	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func FromIndex(index int64) (model.RoleDetail, string, error) {

	var d model.RoleDetail
	var stat string

	client, err := ethclient.DialContext(context.Background(), cmd.Conf.Rpc)
	if err != nil {
		return d, stat, err
	}
	defer client.Close()

	roleMngr, err := b.NewRoleManager(common.HexToAddress(cmd.Conf.RoleMgrAddress), client)
	if err != nil {
		return d, stat, err
	}

	cnt, err := Count()
	if err != nil {
		return d, stat, err
	}

	if index > cnt.Int64()-1 {
		return d, stat, fmt.Errorf("there are only %v Roles, use an index that is lower", cnt)
	}
	d, err = roleMngr.GetRoleDetailsFromIndex(nil, big.NewInt(int64(index)))
	if err != nil {
		return d, stat, err
	}

	stat, err = model.RoleAccesString(d.AccessType)
	if err != nil {
		return d, stat, err
	}

	return d, stat, nil
}

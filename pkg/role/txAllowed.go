package role

import (
	b "github.com/ijsingh82/qV2Cli/bind"
	"github.com/ijsingh82/qV2Cli/cmd"

	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TxAllowed(roleId, orgId, ultPrnt string, txType int64) (bool, error) {

	client, err := ethclient.DialContext(context.Background(), cmd.Conf.Rpc)
	if err != nil {
		fmt.Println("Client Dial Error")
		return false, err
	}
	defer client.Close()

	roleMngr, err := b.NewRoleManager(common.HexToAddress(cmd.Conf.RoleMgrAddress), client)
	if err != nil {
		fmt.Println("Client Dial Error")
		return false, err
	}

	d, err := roleMngr.TransactionAllowed(nil, roleId, orgId, ultPrnt, big.NewInt(txType))
	if err != nil {
		fmt.Println("RoleAccess Error")
		return false, err
	}

	return d, nil
}

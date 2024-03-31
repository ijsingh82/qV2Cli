package inter

import (
	"context"

	b "github.com/ijsingh82/qV2Cli/bind"
	"github.com/ijsingh82/qV2Cli/cmd"
	"github.com/ijsingh82/qV2Cli/pkg/utils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ApproveAdminRoleKey(guardKeyFile, password, account, orgId string) (*types.Receipt, error) {

	var receipt *types.Receipt
	client, err := ethclient.Dial(cmd.Conf.Rpc)
	if err != nil {
		return receipt, err
	}
	defer client.Close()

	add, key, err := utils.ReadKeystore(password, guardKeyFile)
	if err != nil {
		return receipt, err
	}

	auth, err := utils.GetNextAuth(add, key, client)
	if err != nil {
		return receipt, err
	}
	transactor, err := b.NewPermInterfaceTransactor(common.HexToAddress(cmd.Conf.InterfaceAddress), client)
	if err != nil {
		return receipt, err
	}

	tx, err := transactor.ApproveAdminRole(auth, orgId, common.HexToAddress(account))
	if err != nil {
		return receipt, err
	}

	receipt, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return receipt, err
	}

	if receipt.Status == 0 {
		_, err := utils.GetFailingMessage(client, tx.Hash())
		return receipt, err
	}

	return receipt, nil
}

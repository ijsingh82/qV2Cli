package inter

import (
	"context"
	"math/big"

	b "github.com/ijsingh82/qV2Cli/bind"
	"github.com/ijsingh82/qV2Cli/cmd"
	"github.com/ijsingh82/qV2Cli/pkg/utils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// access levels
// 0 - Read only
// 1 - value transfer
// 2 - contract deploy
// 3 - full access
// 4 - contract call
// 5 - value transfer and contract call
// 6 - value transfer and contract deploy
// 7 - contract call and deploy
func AddNewRole(guardKey, guardAdd, roleId, orgId string, voter, admin bool, accsLvl int64) (*types.Receipt, error) {

	var receipt *types.Receipt
	client, err := ethclient.Dial(cmd.Conf.Rpc)
	if err != nil {
		return receipt, err
	}
	defer client.Close()

	auth, err := utils.GetNextAuth(guardAdd, guardKey, client)
	if err != nil {
		return receipt, err
	}

	transactor, err := b.NewPermInterface(common.HexToAddress(cmd.Conf.InterfaceAddress), client)
	if err != nil {
		return receipt, err
	}

	tx, err := transactor.AddNewRole(auth, roleId, orgId, big.NewInt(accsLvl), voter, admin)
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
